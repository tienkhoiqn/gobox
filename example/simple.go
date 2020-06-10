package main

import (
	"github.com/tpphu/gobox"
)

func main() {
	app := gobox.NewApp(
		gobox.Name("test"),
		gobox.WithHTTPService(":3000"))
	app.Description = "test"
	app.Run()
}
