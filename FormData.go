package ajahttp

import (
	"bytes"
	"io"
	"mime/multipart"
)

type formData struct {
	buf *bytes.Buffer
	w   *multipart.Writer
}

func NewFormData() *formData {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	return &formData{buf, w}
}

func (my *formData) Append(fieldname, value string) {
	my.w.WriteField(fieldname, value)
}

func (my *formData) AppendFile(fieldname string, value io.Reader, filename string) error {
	part, err := my.w.CreateFormFile(fieldname, filename)
	if err != nil {
		return err
	}
	fileBytes, _ := io.ReadAll(value)
	part.Write(fileBytes)
	return nil
}
