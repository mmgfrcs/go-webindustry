package main

import (
	"log"

	"github.com/aymerick/raymond"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: handlebars.New("./views", ".hbs").AddFunc("ifeq", func(arg1 string, arg2 string, options *raymond.Options) string {
			if arg1 == arg2 {
				return options.Fn()
			} else {
				return options.Inverse()
			}
		}),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", RenderData{
			SidebarItems: []MenuItem{
				{Type: "item", Text: "Dashboard", Link: "/"},
				{Type: "dropdown", Text: "Resources", Items: []DropdownItem{{Link: "#", Text: "Overview"}, {Link: "#", Text: "Supply/Demand"}}},
				{Type: "item-label", Text: "Base", Label: "99 new", Link: "/"},
				{Type: "item", Text: "Production", Link: "/"},
				{Type: "item", Text: "Market", Link: "/"},
				{Type: "item", Text: "Bank", Link: "/"},
				{Type: "item", Text: "Customers", Link: "/"},
				{Type: "item-label", Text: "Inbox", Label: "1*", Link: "/", LabelColor: "bg-red-500 dark:bg-red-400"},
			},
			ProfileDropItems: []MenuItem{
				{Type: "item", Text: "Profile", Link: "#"},
			},
		})
	})

	app.Static("/static", "./public", fiber.Static{
		Compress: true,
		Download: false,
	})

	log.Fatal(app.Listen(":3001"))
}
