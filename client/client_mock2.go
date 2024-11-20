package client

type MockClient2 struct {
	GetCEPFn      func(cep string) (Cep, error)
	GetLocalityFn func(token, locality string) (Locality, error)
}

func (m *MockClient2) GetCEP(cep string) (Cep, error) {
	return m.GetCEPFn(cep)
}

func (m *MockClient2) GetLocality(token, locality string) (Locality, error) {
	return m.GetLocalityFn(token, locality)
}
