# CC

This application converts an amount in a specific currency and return the amount in all currencies provided by the fixer.io API.

## Running it locally on your computer

1. Copy the application folder in to the directory where you run Go applications
2. This applications uses gorilla mux. Type "go get -u github.com/gorilla/mux" in a terminal to install it
3. To start the application, type "go run main.go" from the application directory
4. The application should now be running on "http://localhost:8000"

## Using the application

The application can now take requests as follows

GET http://localhost:8000/convert?amount={amount}&currency={currency}

Where input parameters are:

- amount (number) - the amount to be converted
- currency (string) - currency that is tied to the amount 

The response can be obtained in json or xml. To get the response in xml, a header `Accept: application/xml` should be present on the request.

Acceptable currencies are:
        AUD 
		BGN 
		BRL 
		CAD 
		CHF 
		CNY 
		CZK 
		DKK 		
        GBP 
		HKD 
		HRK 
		HUF 
		IDR 
		ILS 
		INR 
        JPY 
		KRW 
		MXN 
		MYR 
		NOK 
		NZD 
		PHP 
        PLN 
		RON 
		RUB 
		SEK 
		SGD 
		THB 
        TRY 
		USD 
		ZAR 


## Example

Request: http://localhost:8000/convert?amount=200&currency=SEK

Response:
{
    "amount": 200,
    "currency": "SEK",
    "converted": {
        "AUD": 31.21,
        "BGN": 41.01,
        "BRL": 77.79,
        "CAD": 31.04,
        "CHF": 23.83,
        "CNY": 164.44,
        "CZK": 546.88,
        "DKK": 155.94,
        "GBP": 19.23,
        "HKD": 193.16,
        "HRK": 155.25,
        "HUF": 6367.6,
        "IDR": 329300,
        "ILS": 89.31,
        "INR": 1581.96,
        "JPY": 2699,
        "KRW": 27992,
        "MXN": 435.92,
        "MYR": 105.65,
        "NOK": 195.04,
        "NZD": 33.89,
        "PHP": 1264.92,
        "PLN": 89.77,
        "RON": 96.17,
        "RUB": 1456.84,
        "SEK": 0,
        "SGD": 33.61,
        "THB": 820.36,
        "TRY": 86.32,
        "USD": 24.68,
        "ZAR": 325.84
    }
}

Request http://localhost:8000/convert?amount=200&currency=SEK
With a header `Accept: application/xml`

<Converted>
    <Amount>200</Amount>
    <Currency>SEK</Currency>
    <Converted>
        <AUD>31.21</AUD>
        <BGN>41.01</BGN>
        <BRL>77.79</BRL>
        <CAD>31.04</CAD>
        <CHF>23.83</CHF>
        <CNY>164.44</CNY>
        <CZK>546.88</CZK>
        <DKK>155.94</DKK>
        <GBP>19.23</GBP>
        <HKD>193.16</HKD>
        <HRK>155.25</HRK>
        <HUF>6367.6</HUF>
        <IDR>329300</IDR>
        <ILS>89.31</ILS>
        <INR>1581.96</INR>
        <JPY>2699</JPY>
        <KRW>27992</KRW>
        <MXN>435.92</MXN>
        <MYR>105.65</MYR>
        <NOK>195.04</NOK>
        <NZD>33.89</NZD>
        <PHP>1264.92</PHP>
        <PLN>89.77</PLN>
        <RON>96.17</RON>
        <RUB>1456.84</RUB>
        <SEK>0</SEK>
        <SGD>33.61</SGD>
        <THB>820.36</THB>
        <TRY>86.32</TRY>
        <USD>24.68</USD>
        <ZAR>325.84</ZAR>
    </Converted>
</Converted>
