package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("WantGoApi", func() {
	Title("WantGo API")
	Description("This is WantGoAPI")
	Server("WantGoApi", func() {
		Host("development", func() {
			Description("Development host")
			URI("http://localhost:8081")
		})
		// production host
		// Host("production", func() {
		//     Description("Production hosts.")
		//     // URIs can be parameterized using {param} notation.
		//     URI("https://{version}.goa.design/calc")

		//     // Variable describes a URI variable.
		//     Variable("version", String, "API version", func() {
		//         // URL parameters must have a default value and/or an
		//         // enum validation.
		//         Default("v1")
		//     })
		// })
	})
})

var simpleCard = Type("simpleCard", func() {
	Attribute("cardId", Int)
	Attribute("cardAuthor", String)
	Attribute("cardTitle", String)
	Attribute("imageUrl", String)
})

var cardInfo = Type("cardInfo", func() {
	Attribute("cardAuthor", String)
	Attribute("cardTitle", String)
	Attribute("cardDescription", String)
	Attribute("tags", ArrayOf(String))
	Attribute("imageUrl", String)
	Attribute("locationAddress", String)
	Attribute("locationUrl", String)
})

var _ = Service("WantGo", func() {
	Description("The WantGo service")

	Method("getSimpleCardList", func() {
		Payload(func() {

		})

		Result(ArrayOf(simpleCard))

		HTTP(func() {
			GET("/card-list")
		})
	})

	Method("getCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", Int, "card id")
			Required("cardId")
		})

		Result(cardInfo)

		HTTP(func() {
			GET("/card/{cardId}")
		})
	})

	Method("postCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", Int, "card id")

			Attribute("cardAuthor", String)
			Attribute("cardTitle", String)
			Attribute("cardDescription", String)
			Attribute("tags", ArrayOf(String))
			Attribute("imageUrl", String)
			Attribute("locationAddress", String)
			Attribute("locationUrl", String)

			Required("cardId", "cardAuthor", "cardTitle", "cardDescription")
		})

		Result(Empty)

		HTTP(func() {
			POST("/card/{cardId}")
			Body("cardAuthor", "cardTitle", "cardDescription", "tags", "imageUrl", "locationAddress", "locationUrl")
		})
	})

	Method("putCardInfo", func() {
		Payload(func() {
			Field(1, "cardId", Int, "card id")

			Attribute("cardAuthor", String)
			Attribute("cardTitle", String)
			Attribute("cardDescription", String)
			Attribute("tags", ArrayOf(String))
			Attribute("imageUrl", String)
			Attribute("locationAddress", String)
			Attribute("locationUrl", String)

			Required("cardId", "cardAuthor", "cardTitle", "cardDescription")
		})

		Result(Empty)

		HTTP(func() {
			PUT("/card/{cardId}")
			Body("cardAuthor", "cardTitle", "cardDescription", "tags", "imageUrl", "locationAddress", "locationUrl")
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
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
