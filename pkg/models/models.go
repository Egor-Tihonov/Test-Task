package models

type Weather struct {
	Country                               string
	FeelsLike, MaxTemp, MinTemp, FactTemp int
	WindSpeed                             float64
	Status                                string
}
