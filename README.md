# WantGoApi

this is WantGo api
Goa v3 framework


## generate gen

goa gen WantGoApi/design

## build

go build ./cmd/want_go_api && go build ./cmd/want_go_api-cli

## serve

./want_go_api 


(develop) localhost:8081

- GET /card-list
- GET /card/{cardId}

- POST /card/{cardId}

- PUT /card/{cardId}

- DELETE /card/{cardId}