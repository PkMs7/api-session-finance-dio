package utils

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedTime = StringToDate("2024-02-12T14:15:18")

	if convertedTime.Year() != 2024 {
		t.Errorf("Espera que o ano seja 2024")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja 02")
	}
}
