
golang http clint

## Example
```go
clint := ajahttp.NewAjaClient()
clint.SetBaseURL("http://127.0.0.1:7777")

// get request
clint.Get("/")

// post json
clint.PostJson("/", map[string]any{})


// post file
data := ajahttp.NewFormData()

data.Append("name", "ajahttp")

f, _ := os.Open("/root/a.jpg")
defer f.Close()

data.AppendFile("file", f, f.Name())
clint.PostFormData("/", data)


// get json response
resp, _ := clint.Get("/")
var data map[string]any
ajahttp.JsonResponse(resp, &data)


// orther method
opt := &ajahttp.AjaOption{
  Method: http.MethodDelete,
  Url:    "/",
}

opt.JsonBody(map[string]any{
  "id": 1,
})

clint.Request(opt)
```