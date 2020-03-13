// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo endpoints
//
// Command:
// $ goa gen WantGoApi/design

package wantgo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "WantGo" service endpoints.
type Endpoints struct {
	GetSimpleCardList goa.Endpoint
	GetCardInfo       goa.Endpoint
	PostCardInfo      goa.Endpoint
	PutCardInfo       goa.Endpoint
	DeleteCardInfo    goa.Endpoint
}

// NewEndpoints wraps the methods of the "WantGo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetSimpleCardList: NewGetSimpleCardListEndpoint(s),
		GetCardInfo:       NewGetCardInfoEndpoint(s),
		PostCardInfo:      NewPostCardInfoEndpoint(s),
		PutCardInfo:       NewPutCardInfoEndpoint(s),
		DeleteCardInfo:    NewDeleteCardInfoEndpoint(s),
	}
}

// Use applies the given middleware to all the "WantGo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetSimpleCardList = m(e.GetSimpleCardList)
	e.GetCardInfo = m(e.GetCardInfo)
	e.PostCardInfo = m(e.PostCardInfo)
	e.PutCardInfo = m(e.PutCardInfo)
	e.DeleteCardInfo = m(e.DeleteCardInfo)
}

// NewGetSimpleCardListEndpoint returns an endpoint function that calls the
// method "getSimpleCardList" of service "WantGo".
func NewGetSimpleCardListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSimpleCardList(ctx)
	}
}

// NewGetCardInfoEndpoint returns an endpoint function that calls the method
// "getCardInfo" of service "WantGo".
func NewGetCardInfoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCardInfoPayload)
		return s.GetCardInfo(ctx, p)
	}
}

// NewPostCardInfoEndpoint returns an endpoint function that calls the method
// "postCardInfo" of service "WantGo".
func NewPostCardInfoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PostCardInfoPayload)
		return nil, s.PostCardInfo(ctx, p)
	}
}

// NewPutCardInfoEndpoint returns an endpoint function that calls the method
// "putCardInfo" of service "WantGo".
func NewPutCardInfoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PutCardInfoPayload)
		return nil, s.PutCardInfo(ctx, p)
	}
}

// NewDeleteCardInfoEndpoint returns an endpoint function that calls the method
// "deleteCardInfo" of service "WantGo".
func NewDeleteCardInfoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteCardInfoPayload)
		return nil, s.DeleteCardInfo(ctx, p)
	}
}