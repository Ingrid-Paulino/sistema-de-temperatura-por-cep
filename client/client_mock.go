package client

type MockClient struct{}

func (m *MockClient) GetCEP(c string) (Cep, error) {
	return Cep{Localidade: "Belo Horizonte"}, nil
}

func (m *MockClient) GetLocality(key, location string) (Locality, error) {
	localityResult := Locality{}
	localityResult.Current.TempC = 25.0 // Simulando uma temperatura de 25 graus Celsius
	return localityResult, nil
}
