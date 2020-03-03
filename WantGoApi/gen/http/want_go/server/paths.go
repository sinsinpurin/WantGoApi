// Code generated by goa v3.0.10, DO NOT EDIT.
//
// HTTP request path constructors for the WantGo service.
//
// Command:
// $ goa gen WantGoApi/design

package server

import (
	"fmt"
)

// GetSimpleCardListWantGoPath returns the URL path to the WantGo service getSimpleCardList HTTP endpoint.
func GetSimpleCardListWantGoPath() string {
	return "/card-list"
}

// GetCardInfoWantGoPath returns the URL path to the WantGo service getCardInfo HTTP endpoint.
func GetCardInfoWantGoPath(cardID int) string {
	return fmt.Sprintf("/card/%v", cardID)
}

// PostCardInfoWantGoPath returns the URL path to the WantGo service postCardInfo HTTP endpoint.
func PostCardInfoWantGoPath(cardID int) string {
	return fmt.Sprintf("/card/%v", cardID)
}

// PutCardInfoWantGoPath returns the URL path to the WantGo service putCardInfo HTTP endpoint.
func PutCardInfoWantGoPath(cardID int) string {
	return fmt.Sprintf("/card/%v", cardID)
}

// DeleteCardInfoWantGoPath returns the URL path to the WantGo service deleteCardInfo HTTP endpoint.
func DeleteCardInfoWantGoPath(cardID int) string {
	return fmt.Sprintf("/card/%v", cardID)
}
