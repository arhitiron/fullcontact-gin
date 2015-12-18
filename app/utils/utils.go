package utils
import (
	"net/http"
	"log"
	"net/url"
)

func SendToFullcontact(email string, webhookUrl string, apiKey string) {
	var apiUrl *url.URL
	apiUrl, err := url.Parse("https://api.fullcontact.com")
	if err != nil {
		panic("wtf?")
	}

	apiUrl.Path += "/v2/person.json"
	parameters := url.Values{}
	parameters.Add("email", email)
	parameters.Add("webhookUrl", webhookUrl)
	parameters.Add("apiKey", apiKey)
	apiUrl.RawQuery = parameters.Encode()

	_, httpErr := http.Get(apiUrl.String())
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
