package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Locality struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type Client struct{}

func (r *Client) GetCEP(c string) (Cep, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", c)
	response, err := http.Get(url)
	if err != nil {
		return Cep{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Cep{}, errors.New("erro ao buscar cep")
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return Cep{}, fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	var cep Cep
	err = json.Unmarshal(res, &cep)
	if err != nil {
		return Cep{}, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	return cep, nil
}

func (r *Client) GetLocality(key, location string) (Locality, error) {
	newStr := strings.ReplaceAll(location, " ", "+")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", key, newStr)
	response, err := http.Get(url)

	if err != nil {
		return Locality{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Locality{}, err
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return Locality{}, fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	var locality Locality
	err = json.Unmarshal(res, &locality)
	if err != nil {
		return Locality{}, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	return locality, nil
}
