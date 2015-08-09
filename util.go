package keybase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/kladd/urlvalues"
)

func client(urlStr string) *http.Client {
	client := new(http.Client)

	if session != "" {
		client.Jar, _ = cookiejar.New(nil)
		u, _ := url.Parse(urlStr)

		client.Jar.SetCookies(
			u,
			append(
				[]*http.Cookie(nil),
				&http.Cookie{
					Name:  "session",
					Value: session,
				},
			),
		)
	}

	return client
}

func get(method string, params interface{}, resp interface{}) error {
	qs := make(url.Values)
	encoder := urlvalues.NewEncoder()

	err := encoder.Encode(params, qs)
	urlStr := fmt.Sprintf("%s/%s.json?%s", kbURL, method, qs.Encode())

	res, err := client(urlStr).Get(urlStr)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}

func post(method string, params interface{}, resp interface{}) error {
	qs := make(url.Values)
	encoder := urlvalues.NewEncoder()

	err := encoder.Encode(params, qs)

	urlStr := fmt.Sprintf("%s/%s.json", kbURL, method)

	res, err := client(urlStr).PostForm(urlStr, qs)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}
