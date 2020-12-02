package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
	options, _ := optionHandle(ctx, opts...)
	// 默认加一个30秒的超时
	url, err = generateUrlParams(url, options.params)
	if err != nil {
		return
	}

	req, err := newRequest(ctx, url, method, reqBody)

	generateHeader(req, options.headers)

	httpClient := newClient()

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp.StatusCode, body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, body, err
	}
	return resp.StatusCode, body, err
}

func generateUrlParams(url string, params map[string]string) (string, error) {
	paramsList := make([]string, 0)
	for k, v := range params {
		paramsList = append(paramsList, fmt.Sprintf("%s=%s", k, v))
	}

	if len(paramsList) == 0 {
		return url, nil
	}
	return fmt.Sprintf("%s?%s", url, strings.Join(paramsList, "&")), nil
}

func generateHeader(req *http.Request, headers map[string]string) {
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")

	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}
}

func newRequest(ctx context.Context, url string, method string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	return
}

func newClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return http.Client{Transport: tr}

}
