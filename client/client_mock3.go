package client

import "github.com/stretchr/testify/mock"

type MockClient3 struct {
	mock.Mock
}

func (m *MockClient3) GetCEP(c string) (Cep, error) {
	args := m.Called(c)
	return args.Get(0).(Cep), args.Error(1)
}

func (m *MockClient3) GetLocality(key, location string) (Locality, error) {
	args := m.Called(key, location)
	return args.Get(0).(Locality), args.Error(1)
}
