package main

import (
	"encoding/json"
	"net/http"

	"github.com/Ingrid-Paulino/temperature-system/client"
	"github.com/Ingrid-Paulino/temperature-system/pkg"
)

var Client = client.ClientInterface(&client.Client{})

func main() {
	http.HandleFunc("/", TemperatureConversion)
	http.ListenAndServe(":8080", nil)
}

func TemperatureConversion(w http.ResponseWriter, r *http.Request) {
	CEP := "31330500"
	if len(CEP) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return

	}

	cep, err := Client.GetCEP(CEP)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	locality, err := Client.GetLocality("ec3b6524a81f4014934123201242610", cep.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	result := pkg.NewCelsiusConversion(int(locality.Current.TempC))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
