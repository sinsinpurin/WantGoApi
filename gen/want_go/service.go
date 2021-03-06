// Code generated by goa v3.0.10, DO NOT EDIT.
//
// WantGo service
//
// Command:
// $ goa gen WantGoApi/design

package wantgo

import (
	"context"
)

// The WantGo service
type Service interface {
	// GetSimpleCardList implements getSimpleCardList.
	GetSimpleCardList(context.Context) (res []*SimpleCard, err error)
	// GetCardInfo implements getCardInfo.
	GetCardInfo(context.Context, *GetCardInfoPayload) (res *CardInfo, err error)
	// PostCardInfo implements postCardInfo.
	PostCardInfo(context.Context, *PostCardInfoPayload) (err error)
	// PutCardInfo implements putCardInfo.
	PutCardInfo(context.Context, *PutCardInfoPayload) (err error)
	// DeleteCardInfo implements deleteCardInfo.
	DeleteCardInfo(context.Context, *DeleteCardInfoPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "WantGo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"getSimpleCardList", "getCardInfo", "postCardInfo", "putCardInfo", "deleteCardInfo"}

// GetCardInfoPayload is the payload type of the WantGo service getCardInfo
// method.
type GetCardInfoPayload struct {
	CardID string
}

// CardInfo is the result type of the WantGo service getCardInfo method.
type CardInfo struct {
	CardAuthor      string
	CardTitle       string
	CardDescription string
	Tags            []string
	ImageURL        string
	LocationAddress string
	LocationURL     string
}

// PostCardInfoPayload is the payload type of the WantGo service postCardInfo
// method.
type PostCardInfoPayload struct {
	CardAuthor      string
	CardTitle       string
	CardDescription string
	Tags            []string
	ImageURL        string
	LocationAddress string
	LocationURL     string
}

// PutCardInfoPayload is the payload type of the WantGo service putCardInfo
// method.
type PutCardInfoPayload struct {
	CardID          string
	CardAuthor      string
	CardTitle       string
	CardDescription string
	Tags            []string
	ImageURL        string
	LocationAddress string
	LocationURL     string
}

// DeleteCardInfoPayload is the payload type of the WantGo service
// deleteCardInfo method.
type DeleteCardInfoPayload struct {
	// card id
	CardID int
}

type SimpleCard struct {
	CardID     int
	CardAuthor string
	CardTitle  string
	ImageURL   string
}
