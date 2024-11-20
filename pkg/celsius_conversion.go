package pkg

type Temperature struct {
	TempC int     `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewCelsiusConversion(celsius int) Temperature {
	fahrenheit := (float64(celsius) * 1.8) + 32
	kelvin := float64(celsius) + 273

	return Temperature{
		TempC: celsius,
		TempF: fahrenheit,
		TempK: kelvin,
	}
}
