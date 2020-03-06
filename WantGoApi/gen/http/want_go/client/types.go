// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo HTTP client types
//
// Command:
// $ goa gen WantGoApi/design

package client

import (
	wantgo "WantGoApi/gen/want_go"
)

// PostCardInfoRequestBody is the type of the "WantGo" service "postCardInfo"
// endpoint HTTP request body.
type PostCardInfoRequestBody struct {
	CardAuthor      string   `form:"cardAuthor" json:"cardAuthor" xml:"cardAuthor"`
	CardTitle       string   `form:"cardTitle" json:"cardTitle" xml:"cardTitle"`
	CardDescription string   `form:"cardDescription" json:"cardDescription" xml:"cardDescription"`
	Tags            []string `form:"tags" json:"tags" xml:"tags"`
	ImageURL        string   `form:"imageUrl" json:"imageUrl" xml:"imageUrl"`
	LocationAddress string   `form:"locationAddress" json:"locationAddress" xml:"locationAddress"`
	LocationURL     string   `form:"locationUrl" json:"locationUrl" xml:"locationUrl"`
}

// PutCardInfoRequestBody is the type of the "WantGo" service "putCardInfo"
// endpoint HTTP request body.
type PutCardInfoRequestBody struct {
	CardAuthor      string   `form:"cardAuthor" json:"cardAuthor" xml:"cardAuthor"`
	CardTitle       string   `form:"cardTitle" json:"cardTitle" xml:"cardTitle"`
	CardDescription string   `form:"cardDescription" json:"cardDescription" xml:"cardDescription"`
	Tags            []string `form:"tags" json:"tags" xml:"tags"`
	ImageURL        string   `form:"imageUrl" json:"imageUrl" xml:"imageUrl"`
	LocationAddress string   `form:"locationAddress" json:"locationAddress" xml:"locationAddress"`
	LocationURL     string   `form:"locationUrl" json:"locationUrl" xml:"locationUrl"`
}

// GetSimpleCardListResponseBody is the type of the "WantGo" service
// "getSimpleCardList" endpoint HTTP response body.
type GetSimpleCardListResponseBody []*SimpleCardResponse

// GetCardInfoResponseBody is the type of the "WantGo" service "getCardInfo"
// endpoint HTTP response body.
type GetCardInfoResponseBody struct {
	CardAuthor      *string  `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle       *string  `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	CardDescription *string  `form:"cardDescription,omitempty" json:"cardDescription,omitempty" xml:"cardDescription,omitempty"`
	Tags            []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	ImageURL        *string  `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	LocationAddress *string  `form:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationURL     *string  `form:"locationUrl,omitempty" json:"locationUrl,omitempty" xml:"locationUrl,omitempty"`
}

// SimpleCardResponse is used to define fields on response body types.
type SimpleCardResponse struct {
	CardID     *int    `form:"cardId,omitempty" json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardAuthor *string `form:"cardAuthor,omitempty" json:"cardAuthor,omitempty" xml:"cardAuthor,omitempty"`
	CardTitle  *string `form:"cardTitle,omitempty" json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	ImageURL   *string `form:"imageUrl,omitempty" json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
}

// NewPostCardInfoRequestBody builds the HTTP request body from the payload of
// the "postCardInfo" endpoint of the "WantGo" service.
func NewPostCardInfoRequestBody(p *wantgo.PostCardInfoPayload) *PostCardInfoRequestBody {
	body := &PostCardInfoRequestBody{
		CardAuthor:      p.CardAuthor,
		CardTitle:       p.CardTitle,
		CardDescription: p.CardDescription,
		ImageURL:        p.ImageURL,
		LocationAddress: p.LocationAddress,
		LocationURL:     p.LocationURL,
	}
	if p.Tags != nil {
		body.Tags = make([]string, len(p.Tags))
		for i, val := range p.Tags {
			body.Tags[i] = val
		}
	}
	return body
}

// NewPutCardInfoRequestBody builds the HTTP request body from the payload of
// the "putCardInfo" endpoint of the "WantGo" service.
func NewPutCardInfoRequestBody(p *wantgo.PutCardInfoPayload) *PutCardInfoRequestBody {
	body := &PutCardInfoRequestBody{
		CardAuthor:      p.CardAuthor,
		CardTitle:       p.CardTitle,
		CardDescription: p.CardDescription,
		ImageURL:        p.ImageURL,
		LocationAddress: p.LocationAddress,
		LocationURL:     p.LocationURL,
	}
	if p.Tags != nil {
		body.Tags = make([]string, len(p.Tags))
		for i, val := range p.Tags {
			body.Tags[i] = val
		}
	}
	return body
}

// NewGetSimpleCardListSimpleCardOK builds a "WantGo" service
// "getSimpleCardList" endpoint result from a HTTP "OK" response.
func NewGetSimpleCardListSimpleCardOK(body []*SimpleCardResponse) []*wantgo.SimpleCard {
	v := make([]*wantgo.SimpleCard, len(body))
	for i, val := range body {
		v[i] = &wantgo.SimpleCard{}
		if val.CardID != nil {
			v[i].CardID = *val.CardID
		}
		if val.CardAuthor != nil {
			v[i].CardAuthor = *val.CardAuthor
		}
		if val.CardTitle != nil {
			v[i].CardTitle = *val.CardTitle
		}
		if val.ImageURL != nil {
			v[i].ImageURL = *val.ImageURL
		}
		if val.CardID == nil {
			v[i].CardID = 0
		}
		if val.CardAuthor == nil {
			v[i].CardAuthor = "default"
		}
		if val.CardTitle == nil {
			v[i].CardTitle = "default"
		}
		if val.ImageURL == nil {
			v[i].ImageURL = "default"
		}
	}
	return v
}

// NewGetCardInfoCardInfoOK builds a "WantGo" service "getCardInfo" endpoint
// result from a HTTP "OK" response.
func NewGetCardInfoCardInfoOK(body *GetCardInfoResponseBody) *wantgo.CardInfo {
	v := &wantgo.CardInfo{}
	if body.CardAuthor != nil {
		v.CardAuthor = *body.CardAuthor
	}
	if body.CardTitle != nil {
		v.CardTitle = *body.CardTitle
	}
	if body.CardDescription != nil {
		v.CardDescription = *body.CardDescription
	}
	if body.ImageURL != nil {
		v.ImageURL = *body.ImageURL
	}
	if body.LocationAddress != nil {
		v.LocationAddress = *body.LocationAddress
	}
	if body.LocationURL != nil {
		v.LocationURL = *body.LocationURL
	}
	if body.CardAuthor == nil {
		v.CardAuthor = "default"
	}
	if body.CardTitle == nil {
		v.CardTitle = "default"
	}
	if body.CardDescription == nil {
		v.CardDescription = "default"
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.ImageURL == nil {
		v.ImageURL = "default"
	}
	if body.LocationAddress == nil {
		v.LocationAddress = "default"
	}
	if body.LocationURL == nil {
		v.LocationURL = "default"
	}
	return v
}
