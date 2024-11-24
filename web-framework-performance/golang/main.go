package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/", getIndex)

	e.Logger.Fatal(e.Start(":5001"))
}

func getIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

// import (
// 	"log"
//
// 	"github.com/gofiber/fiber/v3"
// )
//
// func main() {
// 	app := fiber.New()
//
// 	app.Get("/api/hello", func(c fiber.Ctx) error {
// 		return c.SendString("Hello World")
// 	})
//
// 	log.Fatal(app.Listen(":5001"))
// }

// func main() {
// 	m := func(ctx *fasthttp.RequestCtx) {
// 		switch string(ctx.Path()) {
// 		case "/api/hello":
// 			root(ctx)
// 		default:
// 			ctx.Error("not found", fasthttp.StatusNotFound)
// 		}
// 	}
//
// 	fasthttp.ListenAndServe(":5001", m)
// }
//
// func root(ctx *fasthttp.RequestCtx) {
// 	fmt.Fprintf(ctx, "Hello, world")
// }

// func main() {
// 	mux := http.NewServeMux()
//
// 	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, World!"))
// 	})
//
// 	http.ListenAndServe(":5001", mux)
// }
