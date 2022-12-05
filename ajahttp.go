package ajahttp

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

var DefaultClient = NewAjaClient()

func Request(opt *AjaOption) (*http.Response, error) {
	return DefaultClient.Request(opt)
}

func Get(url string) (*http.Response, error) {
	return DefaultClient.Get(url)
}

func Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	return DefaultClient.Post(url, contentType, body)
}

// 发送一个 "application/x-www-form-urlencoded" 请求
func PostForm(url string, data url.Values) (*http.Response, error) {
	return DefaultClient.PostForm(url, data)
}

// 发送一个 "application/json" 请求
func PostJson(url string, data any) (*http.Response, error) {
	return DefaultClient.PostJson(url, data)
}

// 发送一个 "multipart/form-data" 请求
func PostFormData(url string, data *formData, opt AjaOption) (*http.Response, error) {
	return DefaultClient.PostFormData(url, data)
}

func Head(url string) (*http.Response, error) {
	return DefaultClient.Head(url)
}

// 从resp中读取所有字节，并关闭resp.Body
func ByteResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

/*

从resp中提取 json 到 data，并关闭resp.Body

```go
	resp, _ := clint.Get("/", ajahttp.AjaOption{})
	var data map[string]any
	ajahttp.JsonResponse(resp, &data)
```
*/
func JsonResponse(resp *http.Response, data any) error {
	bs, err := ByteResponse(resp)

	if err != nil {
		return err
	}

	return json.Unmarshal(bs, data)
}

// 从resp中提取 xml 到 data，并关闭resp.Body
func XmlResponse(resp *http.Response, data any) error {
	bs, err := ByteResponse(resp)

	if err != nil {
		return err
	}

	return xml.Unmarshal(bs, data)
}
