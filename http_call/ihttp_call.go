package http_call

import (
	"strings"
)

type HttpCall interface {
	Response(res interface{}) (err error)
}

type HttpCallBuilder interface {
	SetUrl(url string) HttpCallBuilder
	SetMethod(method string) HttpCallBuilder
	SetHeader(key, value string) HttpCallBuilder
	SetPayload(payload *strings.Reader) HttpCallBuilder
	Build() HttpCall
}
