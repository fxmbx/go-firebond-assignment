package common

var (
	Bitcoin          = "bitcoin"
	Ethereum         = "ethereum"
	Binancecoin      = "binancecoin"
	Tether           = "tether"
	Dogecoin         = "dogecoin"
	Polkadot         = "polkadot"
	Litecoin         = "litecoin"
	Cryptocurrencies = []string{
		"bitcoin",
		"ethereum",
		"binancecoin",
		"tether",
		"dogecoin",
		"polkadot",
		"litecoin",
	}
	SupportedCrypto = map[string]string{
		"bitcoin":     "bitcoin",
		"ethereum":    "ethereum",
		"binancecoin": "binancecoin",
		"tether":      "tether",
		"dogecoin":    "dogecoin",
		"polkadot":    "polkadot",
		"litecoin":    "litecoin",
	}
)
var (
	NGN            = "ngn"
	JPY            = "jpy"
	USD            = "usd"
	EUR            = "eur"
	FiatCurrencies = []string{
		"ngn", "jpy", "usd", "eur",
	}

	SupportedFiat = map[string]string{
		"ngn": "ngn",
		"jpy": "jpy",
		"usd": "usd",
		"eur": "eur",
	}
)
