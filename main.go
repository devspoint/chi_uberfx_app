package main

import (
	"devspoint/chi_uberfx_app/hello"
	"devspoint/chi_uberfx_app/server"
	"devspoint/chi_uberfx_app/weather"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Options(hello.HelloModule, weather.WeatherModule),
		fx.Invoke(server.Serve),
	)
	app.Run()
}
