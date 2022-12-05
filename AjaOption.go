package ajahttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

const (
	kJson = "application/json"
	kXml  = "application/xml"
	kForm = "application/x-www-form-urlencoded"
)

type AjaOption struct {
	Method  string
	Url     string
	Params  url.Values
	Headers map[string]string
	Body    io.Reader

	contentType string
}

/*
```go
	var body struct {
		XMLName xml.Name `xml:"xml"`
		Name    string   `xml:"name"`
	}
	opt := &ajahttp.AjaOption{}
	body.Name = "ajahttp"
	opt.XmlBody(&body)
```
*/
func (my *AjaOption) XmlBody(data any) error {
	bs, err := xml.Marshal(data)
	if err != nil {
		return err
	}

	my.Body = bytes.NewBuffer(bs)
	my.contentType = kXml
	return nil
}

/*
```go
	body := map[string]any{"name": "ajahttp"}
	opt := &ajahttp.AjaOption{}
	opt.JsonBody(&body)
```
*/
func (my *AjaOption) JsonBody(data any) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	my.Body = bytes.NewBuffer(bs)
	my.contentType = kJson
	return nil
}

/*
```go
	body := url.Values{}
	body.Set("name", "ajahttp")
	opt := &ajahttp.AjaOption{}
	opt.FormBody(body)
```
*/
func (my *AjaOption) FormBody(data url.Values) {
	my.Body = strings.NewReader(data.Encode())
	my.contentType = kForm
}

/*
```go
	fd := ajahttp.NewFormData()

	fd.Append("name", "ajahttp")

	f, _ := os.Open("/root/a.jpg")
	defer f.Close()
	fd.AppendFile("file", f, f.Name())
	fd.Close()

	opt := &ajahttp.AjaOption{}
	opt.FormDataBody(fd)
```
*/
func (my *AjaOption) FormDataBody(data *formData) {
	my.Body = data.buf
	my.contentType = data.w.FormDataContentType()
	data.w.Close()
}

func (my *AjaOption) MultipartBody(body io.Reader, w *multipart.Writer) {
	my.Body = body
	my.contentType = w.FormDataContentType()
}
