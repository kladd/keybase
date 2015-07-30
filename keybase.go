package keybase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

const (
	kbURL = "https://keybase.io/_/api/1.0"
)

type status struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

func encodeParams(params interface{}) string {
	v := reflect.ValueOf(params).Type()
	vals := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		vals.Set(
			v.Field(i).Tag.Get("url"),
			reflect.ValueOf(params).Field(i).String(),
		)
	}

	return vals.Encode()
}

func call(method string, params interface{}, resp interface{}) error {
	res, err := http.Get(fmt.Sprintf("%s/%s.json?%s", kbURL, method, encodeParams(params)))

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, resp)

	return err
}
