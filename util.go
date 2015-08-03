package keybase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func encodeParams(params interface{}) string {
	v := reflect.ValueOf(params)
	vals := url.Values{}
	var value string

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("url")
		var name string
		var opts string

		name = tag
		if idx := strings.Index(tag, ","); idx != -1 {
			name = tag[:idx]
			opts = tag[idx+1:]
		}

		if name == "-" {
			continue
		}

		switch v.Field(i).Type().Kind() {
		case reflect.String:
			value = v.Field(i).String()
		case reflect.Int:
			value = strconv.Itoa(int(v.Field(i).Int()))
			fmt.Println(value)
		}

		if value == "" && strings.Contains(opts, "omitempty") {
			continue
		}

		vals.Set(name, value)
	}

	return vals.Encode()
}

func client(urlStr string) *http.Client {
	client := new(http.Client)

	if session != "" {
		jar, _ := cookiejar.New(nil)
		u, _ := url.Parse(urlStr)

		jar.SetCookies(
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
	urlStr := fmt.Sprintf("%s/%s.json?%s", kbURL, method, encodeParams(params))

	res, err := client(urlStr).Get(urlStr)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}

func post(method string, params url.Values, resp interface{}) error {
	urlStr := fmt.Sprintf("%s/%s.json", kbURL, method)

	res, err := client(urlStr).PostForm(urlStr, params)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}
