package service

import (
	"github.com/Egor-Tihonov/Test-Task/pkg/models"
	owm "github.com/briandowns/openweathermap"
)

// GetWeather: get weather using api open_weather_map
func (s *Service) GetWeather(city string) (*models.Weather, error) {
	w, err := owm.NewCurrent("C", "ru", s.Key)
	if err != nil {
		return nil, err
	}

	w.CurrentByName(city)

	data := models.Weather{
		WindSpeed: w.Wind.Speed,
		Country:   w.Name,
		FactTemp:  int(w.Main.Temp),
		FeelsLike: int(w.Main.FeelsLike),
		MaxTemp:   int(w.Main.TempMax),
		MinTemp:   int(w.Main.TempMin),
	}

	for _, a := range w.Weather {
		data.Status = a.Description
	}

	return &data, nil
}
