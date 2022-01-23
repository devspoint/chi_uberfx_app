package weather

import (
	server2 "devspoint/chi_uberfx_app/server"
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"io/ioutil"
	"net/http"
)

type (
	Handler  struct{}

	Response struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
)

var WeatherModule = fx.Provide(NewWeatherHandler)

func NewWeatherHandler() server2.HandlerOutput {
	return server2.HandlerOutput{Handler: Handler{}}
}

func (h Handler) Method() string {
	return http.MethodGet
}

func (h Handler) Pattern() string {
	return "/weather"
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r, _ := http.Get("https://api.openweathermap.org/data/2.5/weather?q=São Paulo&appid=107de59bf08f7d22f746f20772f2be3e&units=metric")
	data, _ := ioutil.ReadAll(r.Body)
	var weather Response
	json.Unmarshal(data, &weather)
	writer.Write([]byte(fmt.Sprintf("A temperatura atual é de: %v°C\n", weather.Main.Temp)))
}
