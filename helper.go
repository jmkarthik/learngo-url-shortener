package main

var mapURLs map[string] string

func init() {
	mapURLs = map[string]string {
		"goog": "https://google.com/",
		"msft": "https://microsoft.com/",
	}
}

func GetURL(shortCode string) (string, bool) {
	url, found :=  mapURLs[shortCode]
	return url, found
}