// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo client HTTP transport
//
// Command:
// $ goa gen WantGoApi/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the WantGo service endpoint HTTP clients.
type Client struct {
	// GetSimpleCardList Doer is the HTTP client used to make requests to the
	// getSimpleCardList endpoint.
	GetSimpleCardListDoer goahttp.Doer

	// GetCardInfo Doer is the HTTP client used to make requests to the getCardInfo
	// endpoint.
	GetCardInfoDoer goahttp.Doer

	// PostCardInfo Doer is the HTTP client used to make requests to the
	// postCardInfo endpoint.
	PostCardInfoDoer goahttp.Doer

	// PutCardInfo Doer is the HTTP client used to make requests to the putCardInfo
	// endpoint.
	PutCardInfoDoer goahttp.Doer

	// DeleteCardInfo Doer is the HTTP client used to make requests to the
	// deleteCardInfo endpoint.
	DeleteCardInfoDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the WantGo service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetSimpleCardListDoer: doer,
		GetCardInfoDoer:       doer,
		PostCardInfoDoer:      doer,
		PutCardInfoDoer:       doer,
		DeleteCardInfoDoer:    doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// GetSimpleCardList returns an endpoint that makes HTTP requests to the WantGo
// service getSimpleCardList server.
func (c *Client) GetSimpleCardList() goa.Endpoint {
	var (
		decodeResponse = DecodeGetSimpleCardListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetSimpleCardListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetSimpleCardListDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("WantGo", "getSimpleCardList", err)
		}
		return decodeResponse(resp)
	}
}

// GetCardInfo returns an endpoint that makes HTTP requests to the WantGo
// service getCardInfo server.
func (c *Client) GetCardInfo() goa.Endpoint {
	var (
		decodeResponse = DecodeGetCardInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCardInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCardInfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("WantGo", "getCardInfo", err)
		}
		return decodeResponse(resp)
	}
}

// PostCardInfo returns an endpoint that makes HTTP requests to the WantGo
// service postCardInfo server.
func (c *Client) PostCardInfo() goa.Endpoint {
	var (
		encodeRequest  = EncodePostCardInfoRequest(c.encoder)
		decodeResponse = DecodePostCardInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPostCardInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PostCardInfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("WantGo", "postCardInfo", err)
		}
		return decodeResponse(resp)
	}
}

// PutCardInfo returns an endpoint that makes HTTP requests to the WantGo
// service putCardInfo server.
func (c *Client) PutCardInfo() goa.Endpoint {
	var (
		encodeRequest  = EncodePutCardInfoRequest(c.encoder)
		decodeResponse = DecodePutCardInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPutCardInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PutCardInfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("WantGo", "putCardInfo", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCardInfo returns an endpoint that makes HTTP requests to the WantGo
// service deleteCardInfo server.
func (c *Client) DeleteCardInfo() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteCardInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCardInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCardInfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("WantGo", "deleteCardInfo", err)
		}
		return decodeResponse(resp)
	}
}
