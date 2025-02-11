package testapp

import "github.com/ardanlabs/service/foundation/web"

func Routes(app *web.App) {
	app.HandleFunc("GET /test", test)
}
