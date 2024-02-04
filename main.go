package main

import (
	"github.com/collinstommy/go-chat/components"
	"github.com/collinstommy/go-chat/template"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	template.NewTemplateRenderer(e)

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {

		component := components.Index("Jon")

		return template.Html(c, component)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
