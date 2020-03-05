// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGoApi HTTP client CLI support package
//
// Command:
// $ goa gen WantGoApi/design

package cli

import (
	wantgoc "WantGoApi/gen/http/want_go/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `want-go (get-simple-card-list|get-card-info|post-card-info|put-card-info|delete-card-info)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` want-go get-simple-card-list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		wantGoFlags = flag.NewFlagSet("want-go", flag.ContinueOnError)

		wantGoGetSimpleCardListFlags = flag.NewFlagSet("get-simple-card-list", flag.ExitOnError)

		wantGoGetCardInfoFlags      = flag.NewFlagSet("get-card-info", flag.ExitOnError)
		wantGoGetCardInfoCardIDFlag = wantGoGetCardInfoFlags.String("card-id", "REQUIRED", "")

		wantGoPostCardInfoFlags      = flag.NewFlagSet("post-card-info", flag.ExitOnError)
		wantGoPostCardInfoBodyFlag   = wantGoPostCardInfoFlags.String("body", "REQUIRED", "")
		wantGoPostCardInfoCardIDFlag = wantGoPostCardInfoFlags.String("card-id", "REQUIRED", "card id")

		wantGoPutCardInfoFlags      = flag.NewFlagSet("put-card-info", flag.ExitOnError)
		wantGoPutCardInfoBodyFlag   = wantGoPutCardInfoFlags.String("body", "REQUIRED", "")
		wantGoPutCardInfoCardIDFlag = wantGoPutCardInfoFlags.String("card-id", "REQUIRED", "card id")

		wantGoDeleteCardInfoFlags      = flag.NewFlagSet("delete-card-info", flag.ExitOnError)
		wantGoDeleteCardInfoCardIDFlag = wantGoDeleteCardInfoFlags.String("card-id", "REQUIRED", "card id")
	)
	wantGoFlags.Usage = wantGoUsage
	wantGoGetSimpleCardListFlags.Usage = wantGoGetSimpleCardListUsage
	wantGoGetCardInfoFlags.Usage = wantGoGetCardInfoUsage
	wantGoPostCardInfoFlags.Usage = wantGoPostCardInfoUsage
	wantGoPutCardInfoFlags.Usage = wantGoPutCardInfoUsage
	wantGoDeleteCardInfoFlags.Usage = wantGoDeleteCardInfoUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "want-go":
			svcf = wantGoFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "want-go":
			switch epn {
			case "get-simple-card-list":
				epf = wantGoGetSimpleCardListFlags

			case "get-card-info":
				epf = wantGoGetCardInfoFlags

			case "post-card-info":
				epf = wantGoPostCardInfoFlags

			case "put-card-info":
				epf = wantGoPutCardInfoFlags

			case "delete-card-info":
				epf = wantGoDeleteCardInfoFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "want-go":
			c := wantgoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-simple-card-list":
				endpoint = c.GetSimpleCardList()
				data = nil
			case "get-card-info":
				endpoint = c.GetCardInfo()
				data, err = wantgoc.BuildGetCardInfoPayload(*wantGoGetCardInfoCardIDFlag)
			case "post-card-info":
				endpoint = c.PostCardInfo()
				data, err = wantgoc.BuildPostCardInfoPayload(*wantGoPostCardInfoBodyFlag, *wantGoPostCardInfoCardIDFlag)
			case "put-card-info":
				endpoint = c.PutCardInfo()
				data, err = wantgoc.BuildPutCardInfoPayload(*wantGoPutCardInfoBodyFlag, *wantGoPutCardInfoCardIDFlag)
			case "delete-card-info":
				endpoint = c.DeleteCardInfo()
				data, err = wantgoc.BuildDeleteCardInfoPayload(*wantGoDeleteCardInfoCardIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// want-goUsage displays the usage of the want-go command and its subcommands.
func wantGoUsage() {
	fmt.Fprintf(os.Stderr, `The WantGo service
Usage:
    %s [globalflags] want-go COMMAND [flags]

COMMAND:
    get-simple-card-list: GetSimpleCardList implements getSimpleCardList.
    get-card-info: GetCardInfo implements getCardInfo.
    post-card-info: PostCardInfo implements postCardInfo.
    put-card-info: PutCardInfo implements putCardInfo.
    delete-card-info: DeleteCardInfo implements deleteCardInfo.

Additional help:
    %s want-go COMMAND --help
`, os.Args[0], os.Args[0])
}
func wantGoGetSimpleCardListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] want-go get-simple-card-list

GetSimpleCardList implements getSimpleCardList.

Example:
    `+os.Args[0]+` want-go get-simple-card-list
`, os.Args[0])
}

func wantGoGetCardInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] want-go get-card-info -card-id STRING

GetCardInfo implements getCardInfo.
    -card-id STRING: 

Example:
    `+os.Args[0]+` want-go get-card-info --card-id "In et dolor aut culpa amet consequatur."
`, os.Args[0])
}

func wantGoPostCardInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] want-go post-card-info -body STRING -card-id INT

PostCardInfo implements postCardInfo.
    -body STRING: 
    -card-id INT: card id

Example:
    `+os.Args[0]+` want-go post-card-info --body "Similique et cupiditate labore repudiandae et." --card-id 7961823537002089878
`, os.Args[0])
}

func wantGoPutCardInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] want-go put-card-info -body STRING -card-id INT

PutCardInfo implements putCardInfo.
    -body STRING: 
    -card-id INT: card id

Example:
    `+os.Args[0]+` want-go put-card-info --body "Quidem eos minima ab sit eum." --card-id 8314951880641919166
`, os.Args[0])
}

func wantGoDeleteCardInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] want-go delete-card-info -card-id INT

DeleteCardInfo implements deleteCardInfo.
    -card-id INT: card id

Example:
    `+os.Args[0]+` want-go delete-card-info --card-id 1407836864382774373
`, os.Args[0])
}
