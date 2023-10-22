package fetchQuery

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Instance struct {
	BaseURL string
	Headers http.Header
	Timeout time.Duration
}

func NewInstance(url string, headers http.Header, timeout time.Duration) *Instance {
	return &Instance{
		BaseURL: url,
		Headers: headers,
		Timeout: timeout,
	}
}

func (i *Instance) Post(url string, data map[string]interface{}, query ...url.Values) (resp *Response, err error) {
	return postX(fmt.Sprintf("%s/%s", i.BaseURL, url), data, i.Timeout, i.Headers, query...)
}

func (i *Instance) Put(url string, data map[string]interface{}, query ...url.Values) (resp *Response, err error) {
	return putX(fmt.Sprintf("%s/%s", i.BaseURL, url), data, i.Timeout, i.Headers, query...)
}

func (i *Instance) Del(url string, data map[string]interface{}, query ...url.Values) (resp *Response, err error) {
	return delX(fmt.Sprintf("%s/%s", i.BaseURL, url), data, i.Timeout, i.Headers, query...)
}

func (i *Instance) Get(url string, query ...url.Values) (resp *Response, err error) {
	return getX(fmt.Sprintf("%s/%s", i.BaseURL, url), i.Timeout, i.Headers, query...)
}

func Post(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	return postX(url, data, timeout, headers, query...)
}

func Put(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	return putX(url, data, timeout, headers, query...)
}

func Del(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	return delX(url, data, timeout, headers, query...)
}

func Get(url string, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	return getX(url, timeout, headers, query...)
}
