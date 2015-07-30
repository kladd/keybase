package keybase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

const (
	kbURL = "https://keybase.io/_/api/1.0"
)

type status struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

func encodeParams(params interface{}) string {
	v := reflect.ValueOf(params)
	vals := url.Values{}
	var value string

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			value = v.Field(i).String()
		case reflect.Int:
			value = strconv.Itoa(int(v.Field(i).Int()))
			fmt.Println(value)
		}
		vals.Set(
			v.Type().Field(i).Tag.Get("url"),
			value,
		)
	}

	return vals.Encode()
}

func call(method string, params interface{}, resp interface{}) error {
	p := encodeParams(params)
	res, err := http.Get(fmt.Sprintf("%s/%s.json?%s", kbURL, method, p))

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}
