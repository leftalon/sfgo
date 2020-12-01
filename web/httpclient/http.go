package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
)

func Get(ctx context.Context, url string, opts ...Option) (code int, body []byte, err error) {

	return doRequest(ctx, url, "GET", nil, opts...)
}
func GetJson(ctx context.Context, url string, respData interface{}, opts ...Option) (code int, err error) {
	code, body, err := doRequest(ctx, url, "GET", nil, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}
	return
}

func Post(ctx context.Context, url string, reqBody []byte, opts ...Option) (code int, body []byte, err error) {
	return doRequest(ctx, url, "POST", bytes.NewReader(reqBody), opts...)
}

func PostJson(ctx context.Context, url string, reqBody []byte, respData interface{}, opts ...Option) (code int, err error) {
	code, body, err := doRequest(ctx, url, "POST", bytes.NewReader(reqBody), opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}
	return
}

func Put(ctx context.Context, url string, reqBody []byte, opts ...Option) (code int, body []byte, err error) {
	return doRequest(ctx, url, "PUT", bytes.NewReader(reqBody), opts...)
}

func PutJson(ctx context.Context, url string, reqBody []byte, respData interface{}, opts ...Option) (code int, err error) {
	code, body, err := doRequest(ctx, url, "PUT", bytes.NewReader(reqBody), opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}
	return
}

func Delete(ctx context.Context, url string, opts ...Option) (code int, body []byte, err error) {
	return doRequest(ctx, url, "DELETE", nil, opts...)
}

func DeleteJson(ctx context.Context, url string, respData interface{}, opts ...Option) (code int, err error) {
	code, body, err := doRequest(ctx, url, "DELETE", nil, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}
	return
}

func doRequest(ctx context.Context, url string, method string, reqBody io.Reader, opts ...Option) (code int, body []byte, err error) {
	return
}
