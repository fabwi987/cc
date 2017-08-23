package main

import "testing"

func TestCurrencies(t *testing.T) {

	err := CheckValidCurrency("AUD")
	if err != nil {
		t.Error("Correct currency was not accepted")
	}

	err = CheckValidCurrency("ABC")
	if err == nil {
		t.Error("False currency was accepted")
	}

}

func TestAPI(t *testing.T) {

	_, err := MakeAPICall("USD")
	if err != nil {
		t.Error("Error trying to reach the API")
	}

}

func TestConverter(t *testing.T) {

	var currency string
	var amount float64
	var testRates FixerRates

	testRates.Rates.AUD = 10
	testRates.Rates.JPY = 20
	testRates.Rates.USD = 30
	currency = "SEK"
	amount = 10

	checker, err := ConvertToCurrencies(testRates, currency, amount)
	if err != nil {
		t.Error("Error converting")
	}

	if checker.Converted.AUD != 100 {
		t.Errorf("\n...expected = %v\n...obtained = %v", 100, checker.Converted.AUD)
	}
	if checker.Converted.JPY != 200 {
		t.Errorf("\n...expected = %v\n...obtained = %v", 200, checker.Converted.JPY)
	}
	if checker.Converted.USD != 300 {
		t.Errorf("\n...expected = %v\n...obtained = %v", 300, checker.Converted.USD)
	}

}
