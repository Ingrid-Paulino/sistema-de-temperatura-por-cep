package pkg

import "testing"

func TestNewCelsiusConversion(t *testing.T) {
	celsius := 25
	expectedFahrenheit := 77.0
	expectedKelvin := 298.0

	result := NewCelsiusConversion(celsius)

	if result.TempF != expectedFahrenheit {
		t.Errorf("Expected %f, but got %f", expectedFahrenheit, result.TempF)
	}

	if result.TempK != expectedKelvin {
		t.Errorf("Expected %f, but got %f", expectedKelvin, result.TempK)
	}

	if result.TempC != celsius {
		t.Errorf("Expected %d, but got %d", celsius, result.TempC)
	}
}
