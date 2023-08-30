package ajahttp

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrAjaOptionNil = errors.New("AjaOption is nil")
)

// 提取path中的query参数，返回没有query的path
func extractQueryInUrl(path string, query *url.Values) (string, error) {
	n := strings.Index(path, "?")

	if n >= 0 {
		// path中携带了query参数
		p := path[0:n]

		if query != nil {
			q := path[n+1:]
			qv, err := url.ParseQuery(q)
			if err != nil {
				return "", err
			}
			// 将pqth query添加进query
			for k, v := range qv {
				for _, v2 := range v {
					query.Add(k, v2)
				}
			}
		}

		return p, nil
	}

	return path, nil
}

// 向req添加hedaer,如果key重复则覆盖
func setRequestHeaders(req *http.Request, headerList ...map[string]string) {
	if req != nil {
		for _, h := range headerList {
			for k, v := range h {
				req.Header.Set(k, v)
			}
		}
	}
}

type AjaClient struct {
	httpClient *http.Client
	BaseURL    *url.URL
	Headers    map[string]string
	SendBefore func(request *http.Request) error
	SendAfter  func(response *http.Response) error
}

func NewAjaClient() *AjaClient {
	return &AjaClient{
		httpClient: &http.Client{},
	}
}

func (my *AjaClient) SetBaseURL(u any /* string | *url.URL */) error {
	switch uu := u.(type) {
	case string:
		rev, err := url.Parse(uu)
		if err != nil {
			return err
		}
		my.BaseURL = rev
	case *url.URL:
		my.BaseURL = uu
	default:
		return errors.New("BaseURL type error")
	}
	return nil
}

func (my *AjaClient) Request(opt *AjaOption) (*http.Response, error) {
	if opt == nil {
		return nil, ErrAjaOptionNil
	}

	query := url.Values{}

	urlNoQuery, err := extractQueryInUrl(opt.Url, &query)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	if opt.Params != nil {
		for k, v := range opt.Params {
			for _, v2 := range v {
				query.Add(k, v2)
			}
		}
	}

	// 如果scheme非空，则不拼接BaseURL
	tmpUrl, err := url.Parse(urlNoQuery)
	if err != nil {
		return nil, err
	}

	var u *url.URL
	if my.BaseURL == nil || tmpUrl.IsAbs() {
		u = tmpUrl
	} else {
		path, err := url.JoinPath(my.BaseURL.Path, urlNoQuery)
		if err != nil {
			return nil, err
		}

		rel := &url.URL{Path: path}
		u = my.BaseURL.ResolveReference(rel)
	}

	u.RawQuery = query.Encode()

	req, err := http.NewRequest(opt.Method, u.String(), opt.Body)

	if err != nil {
		return nil, err
	}

	// 根据post的数据类型设置 Content-Type
	if opt.contentType != "" {
		req.Header.Set("Content-Type", opt.contentType)
	}

	setRequestHeaders(req, my.Headers, opt.Headers)

	if my.SendBefore != nil {
		err := my.SendBefore(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := my.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if my.SendAfter != nil {
		err := my.SendAfter(resp)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func (my *AjaClient) Get(url string) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method: http.MethodGet,
		Url:    url,
	})
}

func (my *AjaClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method:      http.MethodPost,
		Url:         url,
		Body:        body,
		contentType: contentType,
	})
}

// 发送一个 "application/x-www-form-urlencoded" 请求
func (my *AjaClient) PostForm(url string, data url.Values) (*http.Response, error) {
	opt := AjaOption{
		Method: http.MethodPost,
		Url:    url,
	}
	opt.FormBody(data)
	return my.Request(&opt)
}

// 发送一个 "application/json" 请求
func (my *AjaClient) PostJson(url string, data any) (*http.Response, error) {
	opt := AjaOption{
		Method: http.MethodPost,
		Url:    url,
	}
	opt.JsonBody(data)
	return my.Request(&opt)
}

// 发送一个 "multipart/form-data" 请求
func (my *AjaClient) PostFormData(url string, data *formData) (*http.Response, error) {
	opt := AjaOption{
		Method: http.MethodPost,
		Url:    url,
	}
	opt.FormDataBody(data)
	return my.Request(&opt)
}

func (my *AjaClient) Put(url string) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method: http.MethodPut,
		Url:    url,
	})
}

func (my *AjaClient) Patch(url string) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method: http.MethodPatch,
		Url:    url,
	})
}

func (my *AjaClient) Delete(url string) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method: http.MethodDelete,
		Url:    url,
	})
}

func (my *AjaClient) Head(url string) (*http.Response, error) {
	return my.Request(&AjaOption{
		Method: http.MethodDelete,
		Url:    url,
	})
}
