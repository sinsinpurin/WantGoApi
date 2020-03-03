package wantgoapi

import (
	wantgo "WantGoApi/gen/want_go"
	"context"
	"log"
)

// WantGo service example implementation.
// The example methods log the requests and return zero values.
type wantGosrvc struct {
	logger *log.Logger
}

// NewWantGo returns the WantGo service implementation.
func NewWantGo(logger *log.Logger) wantgo.Service {
	return &wantGosrvc{logger}
}

// GetSimpleCardList implements getSimpleCardList.
func (s *wantGosrvc) GetSimpleCardList(ctx context.Context) (res []*wantgo.SimpleCard, err error) {
	s.logger.Print("wantGo.getSimpleCardList")
	return
}

// GetCardInfo implements getCardInfo.
func (s *wantGosrvc) GetCardInfo(ctx context.Context, p *wantgo.GetCardInfoPayload) (res *wantgo.CardInfo, err error) {
	res = &wantgo.CardInfo{}
	s.logger.Print("wantGo.getCardInfo")
	return
}

// PostCardInfo implements postCardInfo.
func (s *wantGosrvc) PostCardInfo(ctx context.Context, p *wantgo.PostCardInfoPayload) (err error) {
	s.logger.Print("wantGo.postCardInfo")
	return
}

// PutCardInfo implements putCardInfo.
func (s *wantGosrvc) PutCardInfo(ctx context.Context, p *wantgo.PutCardInfoPayload) (err error) {
	s.logger.Print("wantGo.putCardInfo")
	return
}

// DeleteCardInfo implements deleteCardInfo.
func (s *wantGosrvc) DeleteCardInfo(ctx context.Context, p *wantgo.DeleteCardInfoPayload) (err error) {
	s.logger.Print("wantGo.deleteCardInfo")
	return
}
