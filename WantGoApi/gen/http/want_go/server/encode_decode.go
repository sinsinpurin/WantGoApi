// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo HTTP server encoders and decoders
//
// Command:
// $ goa gen WantGoApi/design

package server

import (
	wantgo "WantGoApi/gen/want_go"
	"context"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetSimpleCardListResponse returns an encoder for responses returned by
// the WantGo getSimpleCardList endpoint.
func EncodeGetSimpleCardListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*wantgo.SimpleCard)
		enc := encoder(ctx, w)
		body := NewGetSimpleCardListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetCardInfoResponse returns an encoder for responses returned by the
// WantGo getCardInfo endpoint.
func EncodeGetCardInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*wantgo.CardInfo)
		enc := encoder(ctx, w)
		body := NewGetCardInfoResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetCardInfoRequest returns a decoder for requests sent to the WantGo
// getCardInfo endpoint.
func DecodeGetCardInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			cardID string

			params = mux.Vars(r)
		)
		cardID = params["cardId"]
		payload := NewGetCardInfoPayload(cardID)

		return payload, nil
	}
}

// EncodePostCardInfoResponse returns an encoder for responses returned by the
// WantGo postCardInfo endpoint.
func EncodePostCardInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodePostCardInfoRequest returns a decoder for requests sent to the WantGo
// postCardInfo endpoint.
func DecodePostCardInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body PostCardInfoRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidatePostCardInfoRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewPostCardInfoPayload(&body)

		return payload, nil
	}
}

// EncodePutCardInfoResponse returns an encoder for responses returned by the
// WantGo putCardInfo endpoint.
func EncodePutCardInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodePutCardInfoRequest returns a decoder for requests sent to the WantGo
// putCardInfo endpoint.
func DecodePutCardInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			cardID int

			params = mux.Vars(r)
		)
		{
			cardIDRaw := params["cardId"]
			v, err2 := strconv.ParseInt(cardIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("cardID", cardIDRaw, "integer"))
			}
			cardID = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewPutCardInfoPayload(body, cardID)

		return payload, nil
	}
}

// EncodeDeleteCardInfoResponse returns an encoder for responses returned by
// the WantGo deleteCardInfo endpoint.
func EncodeDeleteCardInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteCardInfoRequest returns a decoder for requests sent to the
// WantGo deleteCardInfo endpoint.
func DecodeDeleteCardInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			cardID int
			err    error

			params = mux.Vars(r)
		)
		{
			cardIDRaw := params["cardId"]
			v, err2 := strconv.ParseInt(cardIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("cardID", cardIDRaw, "integer"))
			}
			cardID = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteCardInfoPayload(cardID)

		return payload, nil
	}
}
