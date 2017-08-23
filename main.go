package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type FixerRates struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates struct {
		AUD float64 `json:"AUD"`
		BGN float64 `json:"BGN"`
		BRL float64 `json:"BRL"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		CNY float64 `json:"CNY"`
		CZK float64 `json:"CZK"`
		DKK float64 `json:"DKK"`
		GBP float64 `json:"GBP"`
		HKD float64 `json:"HKD"`
		HRK float64 `json:"HRK"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		INR float64 `json:"INR"`
		JPY float64 `json:"JPY"`
		KRW float64 `json:"KRW"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		NOK float64 `json:"NOK"`
		NZD float64 `json:"NZD"`
		PHP float64 `json:"PHP"`
		PLN float64 `json:"PLN"`
		RON float64 `json:"RON"`
		RUB float64 `json:"RUB"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		THB float64 `json:"THB"`
		TRY float64 `json:"TRY"`
		USD float64 `json:"USD"`
		ZAR float64 `json:"ZAR"`
	} `json:"rates"`
}

type Converted struct {
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Converted struct {
		AUD float64 `json:"AUD"`
		BGN float64 `json:"BGN"`
		BRL float64 `json:"BRL"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		CNY float64 `json:"CNY"`
		CZK float64 `json:"CZK"`
		DKK float64 `json:"DKK"`
		GBP float64 `json:"GBP"`
		HKD float64 `json:"HKD"`
		HRK float64 `json:"HRK"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		INR float64 `json:"INR"`
		JPY float64 `json:"JPY"`
		KRW float64 `json:"KRW"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		NOK float64 `json:"NOK"`
		NZD float64 `json:"NZD"`
		PHP float64 `json:"PHP"`
		PLN float64 `json:"PLN"`
		RON float64 `json:"RON"`
		RUB float64 `json:"RUB"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		THB float64 `json:"THB"`
		TRY float64 `json:"TRY"`
		USD float64 `json:"USD"`
		ZAR float64 `json:"ZAR"`
	} `json:"converted"`
}

type Currency struct {
	curr string
}

var Client http.Client

var Currencies = []Currency{
	{"AUD"},
	{"BGN"},
	{"BRL"},
	{"CAD"},
	{"CHF"},
	{"CNY"},
	{"CZK"},
	{"DKK"},
	{"GBP"},
	{"HKD"},
	{"HRK"},
	{"HUF"},
	{"IDR"},
	{"ILS"},
	{"INR"},
	{"JPY"},
	{"KRW"},
	{"MXN"},
	{"MYR"},
	{"NOK"},
	{"NZD"},
	{"PHP"},
	{"PLN"},
	{"RON"},
	{"RUB"},
	{"SEK"},
	{"SGD"},
	{"THB"},
	{"TRY"},
	{"USD"},
	{"ZAR"},
}

func main() {

	//Create the router
	r := mux.NewRouter()

	//Endpoint handlers
	r.HandleFunc("/status", ApiStatusHandler).Methods("GET")
	r.HandleFunc("/convert", ConvertCurrencyHandler).Methods("GET")

	//Start the router on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))

}

func ApiStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("API is up and running"))
}

func ConvertCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	currency := r.URL.Query().Get("currency")
	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		err = CheckValidCurrency(currency)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			LatestRates, err := MakeAPICall(currency)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				response, err := ConvertToCurrencies(LatestRates, currency, amount)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					if r.Header.Get("Accept") == "application/xml" {
						payload, err := xml.Marshal(response)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						w.Header().Set("Content-Type", "application/xml")
						w.Write([]byte(payload))
					} else {
						payload, err := json.Marshal(response)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						w.Header().Set("Content-Type", "application/json")
						w.Write([]byte(payload))
					}
				}
			}

		}
	}

}

func MakeAPICall(currency string) (FixerRates, error) {

	var latestCurrencyRates FixerRates

	req, err := http.NewRequest("GET", "http://api.fixer.io/latest?base="+currency, nil)
	if err != nil {
		return latestCurrencyRates, err
	}
	resp, err := Client.Do(req)
	if err != nil {
		return latestCurrencyRates, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&latestCurrencyRates); err != nil {
		return latestCurrencyRates, err
	}

	return latestCurrencyRates, nil
}

func ConvertToCurrencies(LatestCurrencyRates FixerRates, currency string, amount float64) (Converted, error) {

	//Initiate the struct that will hold the converted values
	var convertedToRates Converted
	convertedToRates.Amount = amount
	convertedToRates.Currency = currency

	//Traverse over the LatestRates struct and fill the new struct with values for each currency
	counter := reflect.ValueOf(LatestCurrencyRates.Rates)
	valuePointerLatestRates := reflect.ValueOf(&LatestCurrencyRates.Rates).Elem()
	valuePointerConvertedToRates := reflect.ValueOf(&convertedToRates.Converted).Elem()
	for i := 0; i < counter.Type().NumField(); i++ {
		fieldName := valuePointerLatestRates.Field(i)
		fieldValue := fieldName.Interface().(float64)
		newFieldValue := valuePointerConvertedToRates.Field(i)
		newFieldValue.SetFloat(limitDecimals(fieldValue*amount, 2))
	}
	return convertedToRates, nil

}

func CheckValidCurrency(currency string) error {

	for _, elem := range Currencies {
		if elem.curr == currency {
			return nil
		}
	}

	return errors.New("Invalid currency")
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func limitDecimals(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
