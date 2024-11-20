package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ingrid-Paulino/temperature-system/client"
	"github.com/Ingrid-Paulino/temperature-system/pkg"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureConversionBatch(t *testing.T) {
	type fields struct {
		client client.ClientInterface
	}

	tests := []struct {
		name   string
		fields fields
		//want    pkg.Temperature
		expectedStatus int
	}{
		{
			name: "Should return the temperature",
			fields: fields{
				client: &client.MockClient2{
					GetCEPFn: func(cep string) (client.Cep, error) {
						return client.Cep{
							Localidade: "Belo Horizonte",
							Cep:        "31330-500"}, nil
					},
					GetLocalityFn: func(token, locality string) (client.Locality, error) {
						localityResult := client.Locality{}
						localityResult.Current.TempC = 25.0
						return localityResult, nil
					},
				},
			},
			//want: pkg.Temperature{
			//	TempC: 25.0,
			//	TempF: 77.0,
			//	TempK: 298,
			//},
			expectedStatus: http.StatusOK,
		},
		//{
		//	name: "Should return error when cep is invalid",
		//	expectedStatus: http.StatusNotFound,
		//},
		{
			name: "Should return error when fail GetCEP",
			fields: fields{
				client: &client.MockClient2{
					GetCEPFn: func(cep string) (client.Cep, error) {
						return client.Cep{}, errors.New("can not find zipcode")
					},
				},
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Should return error when fail GetLocality",
			fields: fields{
				client: &client.MockClient2{
					GetCEPFn: func(cep string) (client.Cep, error) {
						return client.Cep{
							Localidade: "Belo Horizonte",
							Cep:        "31330-500"}, nil
					},
					GetLocalityFn: func(token, locality string) (client.Locality, error) {
						return client.Locality{}, errors.New("can not find zipcode")
					},
				},
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client = tt.fields.client
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(TemperatureConversion)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler retornou o status %v, esperado %v", status, http.StatusOK)
			}

			//var temp pkg.Temperature
			//err = json.NewDecoder(rr.Body).Decode(&temp)
			//if err != nil {
			//	t.Fatalf("não foi possível decodificar a resposta: %v", err)
			//}
			//
			//if temp != tt.want {
			//	t.Errorf("handler retornou a temperatura %v, esperado %v", temp, tt.want)
			//}
		})
	}

}

// forma 1 sem mock
func TestTemperatureConversion(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// forma 2 com mock
func TestTemperatureConversion2(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// forma 3 com mock
func TestTemperatureConversion3(t *testing.T) {
	Client = &client.MockClient{}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou o status %v, esperado %v", status, http.StatusOK)
	}

	var temp pkg.Temperature
	err = json.NewDecoder(rr.Body).Decode(&temp)
	if err != nil {
		t.Fatalf("não foi possível decodificar a resposta: %v", err)
	}

	expectedTemp := 25
	if temp.TempC != expectedTemp {
		t.Errorf("handler retornou a temperatura %v, esperado %v", temp.TempC, expectedTemp)
	}
}

// forma 4 com mock
func TestTemperatureConversion4(t *testing.T) {

	mockClient := &client.MockClient2{
		GetCEPFn: func(cep string) (client.Cep, error) {
			return client.Cep{
				Localidade: "Belo Horizonte",
				Cep:        "31330-500"}, nil
		},
		GetLocalityFn: func(token, locality string) (client.Locality, error) {
			localityResult := client.Locality{}
			localityResult.Current.TempC = 25.0
			return localityResult, nil
		},
	}

	Client = mockClient

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou o status %v, esperado %v", status, http.StatusOK)
	}

	var temp pkg.Temperature
	err = json.NewDecoder(rr.Body).Decode(&temp)
	if err != nil {
		t.Fatalf("não foi possível decodificar a resposta: %v", err)
	}

	expectedTemp := pkg.Temperature{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298,
	}

	if temp != expectedTemp {
		t.Errorf("handler retornou a temperatura %v, esperado %v", temp, expectedTemp)
	}
}

// forma 5 com mock
func TestTemperatureConversion5(t *testing.T) {
	mock := &client.MockClient3{}
	localityResult := client.Locality{}
	localityResult.Current.TempC = 25.0
	mock.On("GetCEP", "31330500").Return(client.Cep{Localidade: "Belo Horizonte"}, nil)
	mock.On("GetLocality", "ec3b6524a81f4014934123201242610", "Belo Horizonte").Return(localityResult, nil)

	Client = mock

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou o status %v, esperado %v", status, http.StatusOK)
	}

	var temp pkg.Temperature
	err = json.NewDecoder(rr.Body).Decode(&temp)
	if err != nil {
		t.Fatalf("não foi possível decodificar a resposta: %v", err)
	}

	expectedTemp := 25
	if temp.TempC != expectedTemp {
		t.Errorf("handler retornou a temperatura %v, esperado %v", temp.TempC, expectedTemp)
	}
}

// forma 6 com mock
func TestTemperatureConversion6(t *testing.T) {
	mock := &client.MockClient3{}
	localityResult := client.Locality{}
	localityResult.Current.TempC = 25.0
	mock.On("GetCEP", "31330500").Return(client.Cep{Localidade: "Belo Horizonte"}, nil)
	mock.On("GetLocality", "ec3b6524a81f4014934123201242610", "Belo Horizonte").Return(localityResult, nil)

	Client = mock

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemperatureConversion)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code) //verifica se o status code é 200
	//assert.Equal(t, `{"temp_C":25,"temp_F":77,"temp_K":298}`, rr.Body.String())
	mock.AssertExpectations(t)                 //verifica se todas as expectativas foram atendidas
	mock.AssertCalled(t, "GetCEP", "31330500") //verifica se o método foi chamado com os argumentos corretos
}
