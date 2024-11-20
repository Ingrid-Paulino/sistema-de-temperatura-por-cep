package client

type ClientInterface interface {
	GetCEP(c string) (Cep, error)
	GetLocality(key, location string) (Locality, error)
}
