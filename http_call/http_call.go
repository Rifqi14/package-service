package http_call

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpCallBuild struct {
	url     string
	method  string
	payload *strings.Reader
	header  []httpCallBuildHeader
}

type httpCallBuildHeader struct {
	key   string
	value string
}

func (hc *httpCallBuild) SetUrl(url string) HttpCallBuilder {
	hc.url = url
	return hc
}

func (hc *httpCallBuild) SetMethod(method string) HttpCallBuilder {
	hc.method = method
	return hc
}

func (hc *httpCallBuild) SetHeader(key, value string) HttpCallBuilder {
	hc.header = append(hc.header, httpCallBuildHeader{
		key:   key,
		value: value,
	})

	return hc
}

func (hc *httpCallBuild) SetPayload(payload *strings.Reader) HttpCallBuilder {
	hc.payload = payload
	return hc
}

func (hc *httpCallBuild) Build() HttpCall {
	return &httpCall{
		url:     hc.url,
		method:  hc.method,
		payload: hc.payload,
		header:  hc.header,
	}
}

func New() HttpCallBuilder {
	return &httpCallBuild{}
}

type httpCall struct {
	url     string
	method  string
	payload *strings.Reader
	header  []httpCallBuildHeader
}

func (hc *httpCall) Response(res interface{}) (err error) {

	// Init prerequisite
	urlAPICourierPrice := hc.url
	client := &http.Client{}
	var reqAPI *http.Request
	if hc.payload == nil {
		reqAPI, err = http.NewRequest(hc.method, urlAPICourierPrice, nil)
	} else {
		reqAPI, err = http.NewRequest(hc.method, urlAPICourierPrice, hc.payload)
	}

	if err != nil {
		return err
	}

	for _, header := range hc.header {
		reqAPI.Header.Add(header.key, header.value)
	}

	resAPI, err := client.Do(reqAPI)
	if err != nil {
		return err
	}
	defer resAPI.Body.Close()

	body, err := ioutil.ReadAll(resAPI.Body)
	if err != nil {
		return err
	}

	if resAPI.StatusCode != 200 {
		var respMap map[string]interface{}
		err = json.Unmarshal(body, &respMap)
		if err != nil {
			return err
		}
		return errors.New(respMap["message"].(string))
	}

	err = json.Unmarshal([]byte(string(body)), &res)
	if err != nil {
		return err
	}

	return err
}
