package practice4

import (
	"io/ioutil"
	"net/http"
)

func buildGetRequest(url string, params map[string]string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("appid", API_KEY)
	q.Add("lang", "ru")
	q.Add("units", "metric")
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req
}

func executeRequest(request *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
