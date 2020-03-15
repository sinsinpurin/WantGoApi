package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("WantGoApi", func() {
	Title("WantGo API")
	Description("This is WantGoAPI")
	// API License
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Server("WantGoApi", func() {
		Host("development", func() {
			Description("Development host")
			URI("http://localhost:8081")
		})
		// production host
		Host("production", func() {
			Description("Production hosts.")
			// URIs can be parameterized using {param} notation.
			URI("https://want-go-api.herokuapp.com")
		})
	})

	cors.Origin("https://wantgo-facf0.firebaseapp.com", func() {
		cors.Headers("X-Requested-With", "Content-Type", "X-Token-Auth", "Authorization", "application/json", "Origin", "Accept")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})

	cors.Origin("https://wantgo-facf0.web.app", func() {
		cors.Headers("X-Requested-With", "Content-Type", "X-Token-Auth", "Authorization", "application/json", "Origin", "Accept")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})
})

var simpleCard = Type("simpleCard", func() {
	Attribute("cardId", Int, func() { Default(0) })
	Attribute("cardAuthor", String, func() { Default("default") })
	Attribute("cardTitle", String, func() { Default("default") })
	Attribute("imageUrl", String, func() { Default("default") })

})

var cardInfo = Type("cardInfo", func() {
	Attribute("cardAuthor", String, func() { Default("default") })
	Attribute("cardTitle", String, func() { Default("default") })
	Attribute("cardDescription", String, func() { Default("default") })
	Attribute("tags", ArrayOf(String))
	Attribute("imageUrl", String, func() { Default("default") })
	Attribute("locationAddress", String, func() { Default("default") })
	Attribute("locationUrl", String, func() { Default("default") })
})

var _ = Service("WantGo", func() {
	Description("The WantGo service")

	Method("getSimpleCardList", func() {
		Payload(func() {
		})

		Result(ArrayOf(simpleCard))

		HTTP(func() {
			GET("/card-list")
			Response(StatusOK)
		})
	})

	Method("getCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", String)
			Required("cardId")
		})

		Result(cardInfo)

		HTTP(func() {
			GET("/card/{cardId}")
			Response(StatusOK)
		})
	})

	Method("postCardInfo", func() {
		Payload(func() {
			Attribute("cardAuthor", String)
			Attribute("cardTitle", String)
			Attribute("cardDescription", String)
			Attribute("tags", ArrayOf(String))
			Attribute("imageUrl", String, func() { Default("default") })
			Attribute("locationAddress", String, func() { Default("default") })
			Attribute("locationUrl", String, func() { Default("default") })

			Required("cardAuthor", "cardTitle", "cardDescription")
		})

		Result(Empty)

		HTTP(func() {
			POST("/card")

			Response(StatusOK)
		})
	})

	Method("putCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", String)

			Attribute("cardAuthor", String)
			Attribute("cardTitle", String)
			Attribute("cardDescription", String)
			Attribute("tags", ArrayOf(String))
			Attribute("imageUrl", String, func() { Default("default") })
			Attribute("locationAddress", String, func() { Default("default") })
			Attribute("locationUrl", String, func() { Default("default") })

			Required("cardId", "cardAuthor", "cardTitle", "cardDescription")
		})

		Result(Empty)

		HTTP(func() {
			PUT("/card/{cardId}")

			Response(StatusOK)
		})
	})

	Method("deleteCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", Int, "card id")
			Required("cardId")
		})

		Result(Empty)

		HTTP(func() {
			DELETE("/card/{cardId}")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
