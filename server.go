package main

import (
	"github.com/flada-auxv/symmetrical-octo-barnacle/application/handler"
)

func main() {
	e := handler.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
