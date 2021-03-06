package main

import (
	wantgoapi "WantGoApi"
	wantgo "WantGoApi/gen/want_go"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func defaultAddr() string {
	if s := os.Getenv("PORT"); s != "" {
		return s
	}
	return "8081"
}

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "production", "Server host (valid values: production)")
		domainF   = flag.String("domain", "0.0.0.0", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", os.Getenv("PORT"), "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[wantgoapi] ", log.Ltime)
	}

	var (
		db *sql.DB
	)
	{
		var err error
		//db, err = sql.Open("postgres", "host=ec2-54-152-175-141.compute-1.amazonaws.com user=rxhxfpnzcqbssl password=6ef26778ae7ae04a7ca5392272505eafba6df47f8bd63a2f8d98e6fad1f9683b dbname=dacs3r0n7hnrmp sslmode=require")

		//*使わない
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

		// local db
		//db, err = sql.Open("postgres", "host=localhost user=postgres password=masaki19980929 dbname=want-go-db-test sslmode=disable")

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	}

	// Initialize the services.
	var (
		wantGoSvc wantgo.Service
	)
	{
		wantGoSvc = wantgoapi.NewWantGo(logger, db)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		wantGoEndpoints *wantgo.Endpoints
	)
	{
		wantGoEndpoints = wantgo.NewEndpoints(wantGoSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://0.0.0.0:8081"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":" + defaultAddr()
			}
			handleHTTPServer(ctx, u, wantGoEndpoints, &wg, errc, logger, *dbgF)
		}
	case "production":
		{
			addr := "http://0.0.0.0:8081"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":" + defaultAddr()
			}
			handleHTTPServer(ctx, u, wantGoEndpoints, &wg, errc, logger, *dbgF)
		}
	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
