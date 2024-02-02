package utils

import (
	"io"
	"net/http"
	"time"
)

var httpClient *http.Client

func HTTPClient() *http.Client {
	return httpClient
}

func init() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient = &http.Client{
		Timeout:   time.Minute,
		Transport: t,
	}
}

func GetHTTP(url string) []byte {

	client := HTTPClient()
	req, err := http.NewRequest("GET", url, nil)
	HandleErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	HandleErr(err)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	HandleErr(err)

	return body
}
