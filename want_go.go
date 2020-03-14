package wantgoapi

import (
	wantgo "WantGoApi/gen/want_go"
	"context"
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// WantGo service example implementation.
// The example methods log the requests and return zero values.
type wantGosrvc struct {
	logger *log.Logger
	db     *sql.DB
}

// NewWantGo returns the WantGo service implementation.
func NewWantGo(logger *log.Logger, db *sql.DB) wantgo.Service {
	return &wantGosrvc{logger, db}
}

// GetSimpleCardList implements getSimpleCardList.
func (s *wantGosrvc) GetSimpleCardList(ctx context.Context) (res []*wantgo.SimpleCard, err error) {
	s.logger.Print("wantGo.getSimpleCardList")
	var responses []*wantgo.SimpleCard
	rows, err := s.db.Query(`SELECT "cardId", "cardAuthor", "cardTitle", "imageUrl" FROM "WantCard"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		r := wantgo.SimpleCard{}
		err := rows.Scan(&r.CardID, &r.CardAuthor, &r.CardTitle, &r.ImageURL)
		if err != nil {
			log.Fatal(err)
		}
		responses = append(responses, &r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return responses, nil
}

// GetCardInfo implements getCardInfo.
func (s *wantGosrvc) GetCardInfo(ctx context.Context, p *wantgo.GetCardInfoPayload) (res *wantgo.CardInfo, err error) {
	response := wantgo.CardInfo{}
	s.logger.Print("wantGo.getCardInfo")
	jsonTags := ``
	err = s.db.QueryRow(`SELECT "cardAuthor", "cardTitle", "cardDescription", "tags", "imageUrl" , "locationAddress", "locationUrl" FROM "WantCard" WHERE "cardId" = `+p.CardID).Scan(&response.CardAuthor, &response.CardTitle, &response.CardDescription, &jsonTags, &response.ImageURL, &response.LocationAddress, &response.LocationURL)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(jsonTags), &response.Tags)
	return &response, nil
}

// PostCardInfo implements postCardInfo.
func (s *wantGosrvc) PostCardInfo(ctx context.Context, p *wantgo.PostCardInfoPayload) (err error) {
	s.logger.Println("wantGo.postCardInfo")

	_, erro := s.db.Exec(`INSERT INTO "public"."WantCard" ("cardAuthor", "cardTitle", "cardDescription", "tags", "imageUrl", "locationAddress", "locationUrl") VALUES ( ? , ? , ? , ? , ? , ? , ? );`,
		p.CardAuthor,
		p.CardTitle,
		p.CardDescription,
		stringsToString(p.Tags),
		p.ImageURL,
		p.LocationAddress,
		p.LocationURL,
	)

	if erro != nil {
		log.Fatal(erro)
	}
	return nil
}

// PutCardInfo implements putCardInfo.
func (s *wantGosrvc) PutCardInfo(ctx context.Context, p *wantgo.PutCardInfoPayload) (err error) {
	s.logger.Print("wantGo.putCardInfo")

	_, erro := s.db.Exec(`UPDATE "public"."WantCard" SET "cardAuthor" = ? , "cardTitle" = ? , "cardDescription" = ? , "tags" = ? , "imageUrl" = ? , "locationAddress" = ? , "locationUrl" = ? WHERE "cardId" = ? ;`,
		p.CardAuthor,
		p.CardTitle,
		p.CardDescription,
		stringsToString(p.Tags),
		p.ImageURL,
		p.LocationAddress,
		p.LocationURL,
		p.CardID,
	)
	s.logger.Println(erro)
	if erro != nil {
		log.Fatal(erro)
	}
	return nil
}

// DeleteCardInfo implements deleteCardInfo.
func (s *wantGosrvc) DeleteCardInfo(ctx context.Context, p *wantgo.DeleteCardInfoPayload) (err error) {
	s.logger.Print("wantGo.deleteCardInfo")
	s.logger.Print(p.CardID)
	_, err = s.db.Exec(`DELETE FROM "WantCard" WHERE "cardId" = ?`, p.CardID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//[]string to string
func stringsToString(strs []string) string {
	str := `[`
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			str = str + `"` + strs[i] + `",`
		} else {
			str = str + `"` + strs[i] + `"`
		}
	}
	str = str + `]`
	return str
}
