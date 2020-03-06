// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo HTTP server types
//
// Command:
// $ goa gen WantGoApi/design

package server

import (
	wantgo "WantGoApi/gen/want_go"

	goa "goa.design/goa/v3/pkg"
)

// PostCardInfoRequestBody is the type of the "WantGo" service "postCardInfo"
// endpoint HTTP request body.
type PostCardInfoRequestBody struct {
	CardAuthor      *string  `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle       *string  `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	CardDescription *string  `form:"cardDescription,omitempty" json:"cardDescription,omitempty" xml:"cardDescription,omitempty"`
	Tags            []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	ImageURL        *string  `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	LocationAddress *string  `form:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationURL     *string  `form:"locationUrl,omitempty" json:"locationUrl,omitempty" xml:"locationUrl,omitempty"`
}

// PutCardInfoRequestBody is the type of the "WantGo" service "putCardInfo"
// endpoint HTTP request body.
type PutCardInfoRequestBody struct {
	CardAuthor      *string  `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle       *string  `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	CardDescription *string  `form:"cardDescription,omitempty" json:"cardDescription,omitempty" xml:"cardDescription,omitempty"`
	Tags            []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	ImageURL        *string  `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	LocationAddress *string  `form:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationURL     *string  `form:"locationUrl,omitempty" json:"locationUrl,omitempty" xml:"locationUrl,omitempty"`
}

// GetSimpleCardListResponseBody is the type of the "WantGo" service
// "getSimpleCardList" endpoint HTTP response body.
type GetSimpleCardListResponseBody []*SimpleCardResponse

// GetCardInfoResponseBody is the type of the "WantGo" service "getCardInfo"
// endpoint HTTP response body.
type GetCardInfoResponseBody struct {
	CardAuthor      string   `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle       string   `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	CardDescription string   `form:"cardDescription,omitempty" json:"cardDescription,omitempty" xml:"cardDescription,omitempty"`
	Tags            []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	ImageURL        string   `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	LocationAddress string   `form:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationURL     string   `form:"locationUrl,omitempty" json:"locationUrl,omitempty" xml:"locationUrl,omitempty"`
}

// SimpleCardResponse is used to define fields on response body types.
type SimpleCardResponse struct {
	CardID     int    `form:"cardId,omitempty" json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardAuthor string `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle  string `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	ImageURL   string `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
}

// NewGetSimpleCardListResponseBody builds the HTTP response body from the
// result of the "getSimpleCardList" endpoint of the "WantGo" service.
func NewGetSimpleCardListResponseBody(res []*wantgo.SimpleCard) GetSimpleCardListResponseBody {
	body := make([]*SimpleCardResponse, len(res))
	for i, val := range res {
		body[i] = &SimpleCardResponse{
			CardID:     val.CardID,
			CardAuthor: val.CardAuthor,
			CardTitle:  val.CardTitle,
			ImageURL:   val.ImageURL,
		}
	}
	return body
}

// NewGetCardInfoResponseBody builds the HTTP response body from the result of
// the "getCardInfo" endpoint of the "WantGo" service.
func NewGetCardInfoResponseBody(res *wantgo.CardInfo) *GetCardInfoResponseBody {
	body := &GetCardInfoResponseBody{
		CardAuthor:      res.CardAuthor,
		CardTitle:       res.CardTitle,
		CardDescription: res.CardDescription,
		ImageURL:        res.ImageURL,
		LocationAddress: res.LocationAddress,
		LocationURL:     res.LocationURL,
	}
	if res.Tags != nil {
		body.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			body.Tags[i] = val
		}
	}
	return body
}

// NewGetCardInfoPayload builds a WantGo service getCardInfo endpoint payload.
func NewGetCardInfoPayload(cardID string) *wantgo.GetCardInfoPayload {
	return &wantgo.GetCardInfoPayload{
		CardID: cardID,
	}
}

// NewPostCardInfoPayload builds a WantGo service postCardInfo endpoint payload.
func NewPostCardInfoPayload(body *PostCardInfoRequestBody) *wantgo.PostCardInfoPayload {
	v := &wantgo.PostCardInfoPayload{
		CardAuthor:      *body.CardAuthor,
		CardTitle:       *body.CardTitle,
		CardDescription: *body.CardDescription,
		ImageURL:        *body.ImageURL,
		LocationAddress: *body.LocationAddress,
		LocationURL:     *body.LocationURL,
	}
	v.Tags = make([]string, len(body.Tags))
	for i, val := range body.Tags {
		v.Tags[i] = val
	}
	return v
}

// NewPutCardInfoPayload builds a WantGo service putCardInfo endpoint payload.
func NewPutCardInfoPayload(body *PutCardInfoRequestBody, cardID string) *wantgo.PutCardInfoPayload {
	v := &wantgo.PutCardInfoPayload{
		CardAuthor:      *body.CardAuthor,
		CardTitle:       *body.CardTitle,
		CardDescription: *body.CardDescription,
		ImageURL:        *body.ImageURL,
		LocationAddress: *body.LocationAddress,
		LocationURL:     *body.LocationURL,
	}
	v.Tags = make([]string, len(body.Tags))
	for i, val := range body.Tags {
		v.Tags[i] = val
	}
	v.CardID = cardID
	return v
}

// NewDeleteCardInfoPayload builds a WantGo service deleteCardInfo endpoint
// payload.
func NewDeleteCardInfoPayload(cardID int) *wantgo.DeleteCardInfoPayload {
	return &wantgo.DeleteCardInfoPayload{
		CardID: cardID,
	}
}

// ValidatePostCardInfoRequestBody runs the validations defined on
// PostCardInfoRequestBody
func ValidatePostCardInfoRequestBody(body *PostCardInfoRequestBody) (err error) {
	if body.CardAuthor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardAuthor", "body"))
	}
	if body.CardTitle == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardTitle", "body"))
	}
	if body.CardDescription == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardDescription", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	if body.ImageURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("imageUrl", "body"))
	}
	if body.LocationAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("locationAddress", "body"))
	}
	if body.LocationURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("locationUrl", "body"))
	}
	return
}

// ValidatePutCardInfoRequestBody runs the validations defined on
// PutCardInfoRequestBody
func ValidatePutCardInfoRequestBody(body *PutCardInfoRequestBody) (err error) {
	if body.CardAuthor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardAuthor", "body"))
	}
	if body.CardTitle == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardTitle", "body"))
	}
	if body.CardDescription == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cardDescription", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	if body.ImageURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("imageUrl", "body"))
	}
	if body.LocationAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("locationAddress", "body"))
	}
	if body.LocationURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("locationUrl", "body"))
	}
	return
}
