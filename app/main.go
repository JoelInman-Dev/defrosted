package main

import (
	"github.com/joelinman-nxp/defrosted/app/data"
	"github.com/joelinman-nxp/defrosted/app/routes"
)

func main() {
		println("Hello World")

		routes.Setup()
		data.Connect()
}