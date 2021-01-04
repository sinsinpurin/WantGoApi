// Code generated by goa v3.2.6, DO NOT EDIT.
//
// WantGo HTTP client encoders and decoders
//
// Command:
// $ goa gen WantGoApi/design

package client

import (
	wantgo "WantGoApi/gen/want_go"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetSimpleCardListRequest instantiates a HTTP request object with method
// and path set to call the "WantGo" service "getSimpleCardList" endpoint
func (c *Client) BuildGetSimpleCardListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetSimpleCardListWantGoPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("WantGo", "getSimpleCardList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetSimpleCardListRequest returns an encoder for requests sent to the
// WantGo getSimpleCardList server.
func EncodeGetSimpleCardListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*wantgo.GetSimpleCardListPayload)
		if !ok {
			return goahttp.ErrInvalidType("WantGo", "getSimpleCardList", "*wantgo.GetSimpleCardListPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetSimpleCardListResponse returns a decoder for responses returned by
// the WantGo getSimpleCardList endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetSimpleCardListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetSimpleCardListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("WantGo", "getSimpleCardList", err)
			}
			res := NewGetSimpleCardListSimpleCardOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("WantGo", "getSimpleCardList", resp.StatusCode, string(body))
		}
	}
}

// BuildGetCardInfoRequest instantiates a HTTP request object with method and
// path set to call the "WantGo" service "getCardInfo" endpoint
func (c *Client) BuildGetCardInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		cardID string
	)
	{
		p, ok := v.(*wantgo.GetCardInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("WantGo", "getCardInfo", "*wantgo.GetCardInfoPayload", v)
		}
		cardID = p.CardID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCardInfoWantGoPath(cardID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("WantGo", "getCardInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetCardInfoRequest returns an encoder for requests sent to the WantGo
// getCardInfo server.
func EncodeGetCardInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*wantgo.GetCardInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("WantGo", "getCardInfo", "*wantgo.GetCardInfoPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetCardInfoResponse returns a decoder for responses returned by the
// WantGo getCardInfo endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetCardInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetCardInfoResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("WantGo", "getCardInfo", err)
			}
			res := NewGetCardInfoCardInfoOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("WantGo", "getCardInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildPostCardInfoRequest instantiates a HTTP request object with method and
// path set to call the "WantGo" service "postCardInfo" endpoint
func (c *Client) BuildPostCardInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PostCardInfoWantGoPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("WantGo", "postCardInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePostCardInfoRequest returns an encoder for requests sent to the WantGo
// postCardInfo server.
func EncodePostCardInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*wantgo.PostCardInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("WantGo", "postCardInfo", "*wantgo.PostCardInfoPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set("Authorization", head)
		}
		body := NewPostCardInfoRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("WantGo", "postCardInfo", err)
		}
		return nil
	}
}

// DecodePostCardInfoResponse returns a decoder for responses returned by the
// WantGo postCardInfo endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodePostCardInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("WantGo", "postCardInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildPutCardInfoRequest instantiates a HTTP request object with method and
// path set to call the "WantGo" service "putCardInfo" endpoint
func (c *Client) BuildPutCardInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		cardID string
	)
	{
		p, ok := v.(*wantgo.PutCardInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("WantGo", "putCardInfo", "*wantgo.PutCardInfoPayload", v)
		}
		cardID = p.CardID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PutCardInfoWantGoPath(cardID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("WantGo", "putCardInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePutCardInfoRequest returns an encoder for requests sent to the WantGo
// putCardInfo server.
func EncodePutCardInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*wantgo.PutCardInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("WantGo", "putCardInfo", "*wantgo.PutCardInfoPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set("Authorization", head)
		}
		body := NewPutCardInfoRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("WantGo", "putCardInfo", err)
		}
		return nil
	}
}

// DecodePutCardInfoResponse returns a decoder for responses returned by the
// WantGo putCardInfo endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodePutCardInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("WantGo", "putCardInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteCardInfoRequest instantiates a HTTP request object with method
// and path set to call the "WantGo" service "deleteCardInfo" endpoint
func (c *Client) BuildDeleteCardInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		cardID int
	)
	{
		p, ok := v.(*wantgo.DeleteCardInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("WantGo", "deleteCardInfo", "*wantgo.DeleteCardInfoPayload", v)
		}
		cardID = p.CardID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteCardInfoWantGoPath(cardID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("WantGo", "deleteCardInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteCardInfoRequest returns an encoder for requests sent to the
// WantGo deleteCardInfo server.
func EncodeDeleteCardInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*wantgo.DeleteCardInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("WantGo", "deleteCardInfo", "*wantgo.DeleteCardInfoPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeDeleteCardInfoResponse returns a decoder for responses returned by the
// WantGo deleteCardInfo endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDeleteCardInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("WantGo", "deleteCardInfo", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSimpleCardResponseToWantgoSimpleCard builds a value of type
// *wantgo.SimpleCard from a value of type *SimpleCardResponse.
func unmarshalSimpleCardResponseToWantgoSimpleCard(v *SimpleCardResponse) *wantgo.SimpleCard {
	res := &wantgo.SimpleCard{}
	if v.CardID != nil {
		res.CardID = *v.CardID
	}
	if v.CardAuthor != nil {
		res.CardAuthor = *v.CardAuthor
	}
	if v.CardTitle != nil {
		res.CardTitle = *v.CardTitle
	}
	if v.ImageURL != nil {
		res.ImageURL = *v.ImageURL
	}
	if v.CardID == nil {
		res.CardID = 0
	}
	if v.CardAuthor == nil {
		res.CardAuthor = "default"
	}
	if v.CardTitle == nil {
		res.CardTitle = "default"
	}
	if v.ImageURL == nil {
		res.ImageURL = "default"
	}

	return res
}
