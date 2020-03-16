// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo HTTP server
//
// Command:
// $ goa gen WantGoApi/design

package server

import (
	wantgo "WantGoApi/gen/want_go"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the WantGo service endpoint HTTP handlers.
type Server struct {
	Mounts            []*MountPoint
	GetSimpleCardList http.Handler
	GetCardInfo       http.Handler
	PostCardInfo      http.Handler
	PutCardInfo       http.Handler
	DeleteCardInfo    http.Handler
	CORS              http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the WantGo service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *wantgo.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetSimpleCardList", "GET", "/card-list"},
			{"GetCardInfo", "GET", "/card/{cardId}"},
			{"PostCardInfo", "POST", "/card"},
			{"PutCardInfo", "PUT", "/card/{cardId}"},
			{"DeleteCardInfo", "DELETE", "/card/{cardId}"},
			{"CORS", "OPTIONS", "/card-list"},
			{"CORS", "OPTIONS", "/card/{cardId}"},
			{"CORS", "OPTIONS", "/card"},
			{"CORS", "OPTIONS", "/openapi.json"},
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		GetSimpleCardList: NewGetSimpleCardListHandler(e.GetSimpleCardList, mux, decoder, encoder, errhandler, formatter),
		GetCardInfo:       NewGetCardInfoHandler(e.GetCardInfo, mux, decoder, encoder, errhandler, formatter),
		PostCardInfo:      NewPostCardInfoHandler(e.PostCardInfo, mux, decoder, encoder, errhandler, formatter),
		PutCardInfo:       NewPutCardInfoHandler(e.PutCardInfo, mux, decoder, encoder, errhandler, formatter),
		DeleteCardInfo:    NewDeleteCardInfoHandler(e.DeleteCardInfo, mux, decoder, encoder, errhandler, formatter),
		CORS:              NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "WantGo" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetSimpleCardList = m(s.GetSimpleCardList)
	s.GetCardInfo = m(s.GetCardInfo)
	s.PostCardInfo = m(s.PostCardInfo)
	s.PutCardInfo = m(s.PutCardInfo)
	s.DeleteCardInfo = m(s.DeleteCardInfo)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the WantGo endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetSimpleCardListHandler(mux, h.GetSimpleCardList)
	MountGetCardInfoHandler(mux, h.GetCardInfo)
	MountPostCardInfoHandler(mux, h.PostCardInfo)
	MountPutCardInfoHandler(mux, h.PutCardInfo)
	MountDeleteCardInfoHandler(mux, h.DeleteCardInfo)
	MountCORSHandler(mux, h.CORS)
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/http/openapi.json")
	}))
}

// MountGetSimpleCardListHandler configures the mux to serve the "WantGo"
// service "getSimpleCardList" endpoint.
func MountGetSimpleCardListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleWantGoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/card-list", f)
}

// NewGetSimpleCardListHandler creates a HTTP handler which loads the HTTP
// request and calls the "WantGo" service "getSimpleCardList" endpoint.
func NewGetSimpleCardListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeGetSimpleCardListResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getSimpleCardList")
		ctx = context.WithValue(ctx, goa.ServiceKey, "WantGo")
		var err error

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetCardInfoHandler configures the mux to serve the "WantGo" service
// "getCardInfo" endpoint.
func MountGetCardInfoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleWantGoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/card/{cardId}", f)
}

// NewGetCardInfoHandler creates a HTTP handler which loads the HTTP request
// and calls the "WantGo" service "getCardInfo" endpoint.
func NewGetCardInfoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetCardInfoRequest(mux, decoder)
		encodeResponse = EncodeGetCardInfoResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getCardInfo")
		ctx = context.WithValue(ctx, goa.ServiceKey, "WantGo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountPostCardInfoHandler configures the mux to serve the "WantGo" service
// "postCardInfo" endpoint.
func MountPostCardInfoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleWantGoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/card", f)
}

// NewPostCardInfoHandler creates a HTTP handler which loads the HTTP request
// and calls the "WantGo" service "postCardInfo" endpoint.
func NewPostCardInfoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePostCardInfoRequest(mux, decoder)
		encodeResponse = EncodePostCardInfoResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "postCardInfo")
		ctx = context.WithValue(ctx, goa.ServiceKey, "WantGo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountPutCardInfoHandler configures the mux to serve the "WantGo" service
// "putCardInfo" endpoint.
func MountPutCardInfoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleWantGoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/card/{cardId}", f)
}

// NewPutCardInfoHandler creates a HTTP handler which loads the HTTP request
// and calls the "WantGo" service "putCardInfo" endpoint.
func NewPutCardInfoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePutCardInfoRequest(mux, decoder)
		encodeResponse = EncodePutCardInfoResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "putCardInfo")
		ctx = context.WithValue(ctx, goa.ServiceKey, "WantGo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteCardInfoHandler configures the mux to serve the "WantGo" service
// "deleteCardInfo" endpoint.
func MountDeleteCardInfoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleWantGoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/card/{cardId}", f)
}

// NewDeleteCardInfoHandler creates a HTTP handler which loads the HTTP request
// and calls the "WantGo" service "deleteCardInfo" endpoint.
func NewDeleteCardInfoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteCardInfoRequest(mux, decoder)
		encodeResponse = EncodeDeleteCardInfoResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteCardInfo")
		ctx = context.WithValue(ctx, goa.ServiceKey, "WantGo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", handleWantGoOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service WantGo.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleWantGoOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/card-list", f)
	mux.Handle("OPTIONS", "/card/{cardId}", f)
	mux.Handle("OPTIONS", "/card", f)
	mux.Handle("OPTIONS", "/openapi.json", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleWantGoOrigin applies the CORS response headers corresponding to the
// origin for the service WantGo.
func handleWantGoOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://wantgo-facf0.firebaseapp.com") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, X-Token-Auth, Authorization, application/json, Origin, Accept, text/plain")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://wantgo-facf0.web.app") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, X-Token-Auth, Authorization, application/json, Origin, Accept, text/plain")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "＊") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, X-Token-Auth, Authorization, application/json, Origin, Accept, text/plain")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
