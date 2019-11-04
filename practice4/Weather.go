package practice4

import (
	"encoding/json"
	"fmt"
	"time"
)

type Wind struct {
	Speed float64
	Deg   int
	Gust  int
}

type Sun struct {
	Sunrise int
	Sunset  int
}

type Meta struct {
	CityName string
	Timezone int
}

type Weather struct {
	Temp        float64
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Description string
	Humidity    int
	Wind        Wind
	Sun         Sun
	Meta        Meta
}

type DailyWeather struct {
	Meta    Meta
	weather []Weather
}

type Meteorologist struct{}

type Sections map[string]json.RawMessage

func validateResponseCode(code interface{}, body *[]byte) {
	switch v := code.(type) {
	default:
		panic("Invalid type command")
	case float64:
		if v != float64(200) {
			panic("Something went wrong: " + string(*body))
		}
	case int:
		if v != 200 {
			panic("Something went wrong: " + string(*body))
		}
	case string:
		if v == "404" {
			panic("Invalid city name")
		}
	}
}

func (s *Meteorologist) WeatherForecast(city string) *Weather {
	var req = buildGetRequest(API_URL+"weather", map[string]string{
		"q": city,
	})

	body := executeRequest(req)
	var sections Sections
	if err := json.Unmarshal(body, &sections); err != nil {
		panic(err)
	}
	var code interface{}
	if err := json.Unmarshal(sections["cod"], &code); err != nil {
		panic(err)
	}
	validateResponseCode(code, &body)

	weatherSlice := &[]Weather{}
	if err := json.Unmarshal(sections["weather"], weatherSlice); err != nil {
		panic(err)
	}
	weather := &(*weatherSlice)[0]
	weather.Meta.CityName = city
	if err := json.Unmarshal(sections["main"], weather); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(sections["wind"], &weather.Wind); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &weather.Meta); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(sections["sys"], &weather.Sun); err != nil {
		panic(err)
	}

	return weather
}

func (s *Meteorologist) DailyForecast(city string, cnt int) DailyWeather {
	var req = buildGetRequest(API_URL+"forecast/daily", map[string]string{
		"q":   city,
		"cnt": string(cnt),
	})

	body := executeRequest(req)
	var sections Sections
	if err := json.Unmarshal(body, &sections); err != nil {
		panic(err)
	}
	var code interface{}
	if err := json.Unmarshal(sections["cod"], &code); err != nil {
		panic(err)
	}
	validateResponseCode(code, &body)

	return DailyWeather{}
}

func (s *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return s.Temp, s.TempMin, s.TempMax
}

func (s *Weather) GetCloudiness() string {
	return s.Description
}

func (s *Weather) GetHumidity() int {
	return s.Humidity
}

func (s *Weather) GetWind() *Wind {
	return &s.Wind
}

func (s *Weather) String() string {
	temp, _, _ := s.GetTemperature()
	wind := s.GetWind()
	loc := time.FixedZone("", s.Meta.Timezone)
	sunrise := time.Unix(int64(s.Sun.Sunrise), 0).In(loc)
	sunset := time.Unix(int64(s.Sun.Sunset), 0).In(loc)

	return fmt.Sprintf(
		"Сегодня в городе %s %s, температура воздуха %.1f°С, ветер %.1fм/с с порывами до %dм/с, направление %s. Влажность воздуха %d%%. Восход солнца %s, заход солнца %s",
		s.Meta.CityName,
		s.GetCloudiness(),
		temp,
		wind.Speed,
		wind.Gust,
		s.windCompass(),
		s.GetHumidity(),
		sunrise.Format("15:04"),
		sunset.Format("15:04"),
	)
}

func (s *Weather) windCompass() string {
	val := (float64(s.Wind.Deg) / 22.5) + .5
	arr := [...]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	return arr[int(val)%16]
}
