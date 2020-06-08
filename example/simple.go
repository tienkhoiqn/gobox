package main

import (
	"github.com/tpphu/gobox"
)

func main() {
	app := gobox.NewApp(gobox.Name("test"))
	app.Description = "test"
	app.Run()
}
