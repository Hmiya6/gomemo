# net/http の HTTPリクエスト関連まとめ
go で HTTPリクエスト を送るときに関係する諸々についてまとめたもの  
主に扱うのは以下
* `http.Client`
* `http.Request`
* `http.Response`
* `http.Transport`

## 0. 最もシンプルな実装方法 `http.Get()`, `http.Post()`
  
デフォルトのクライアントでデフォルトのリクエストを送る.  
後述の `http.Response` 型(のポインタ)を返す.

### `http.Get()` について
例
```go
resp, err := http.Get("<URL>")
```

### `http.Post()` について
packageでの実装  
```go
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
```
送り先の URL, ヘッダの Content-Type, リクエストボディとなる `io.Reader` 型が引数となる.

`http.Post()` の例  
1. 画像ファイルを post する
```go
// set file reader
file, err := os.Open("<filename>")
if err != nil {
    /* handle error */
}

// set Content-Type
contentType := "image/jpeg"

// execute POST request
resp, err := http.Post("<URL>", contentType, file)
if err != nil {
    /* handle error */
}

/*
    Do something
*/
```
2. テキストを post する
```go
// set string reader for http.Post() at first
text := "Hello, http.Post()"
reader := strings.NewReader(text)

// set Content-Type
contentType := "text/plain"

// execute POST request
resp, err := http.Ppost("<URL>", contentType, reader)
if err != nil {
    /* handle error */
}

/*
    Do something
*/
```

## 1. `http.Client` 型
標準ライブラリでの実装 (一部)
```go
// A Client is an HTTP client. Its zero value (DefaultClient) is a

// usable client that uses DefaultTransport.

//

// A Client is higher-level than a RoundTripper (such as Transport)

// and additionally handles HTTP details such as cookies and

// redirects.

//

type Client struct {

	// Transport specifies the mechanism by which individual

	// HTTP requests are made.

	// If nil, DefaultTransport is used.

	Transport RoundTripper


	// CheckRedirect specifies the policy for handling redirects.

	// If CheckRedirect is not nil, the client calls it before

	// following an HTTP redirect. The arguments req and via are

	// the upcoming request and the requests made already, oldest

	// first. If CheckRedirect returns an error, the Client's Get

	// method returns both the previous Response (with its Body

	// closed) and CheckRedirect's error (wrapped in a url.Error)

	// instead of issuing the Request req.

	// As a special case, if CheckRedirect returns ErrUseLastResponse,

	// then the most recent response is returned with its body

	// unclosed, along with a nil error.

	//

	// If CheckRedirect is nil, the Client uses its default policy,

	// which is to stop after 10 consecutive requests.

	CheckRedirect func(req *Request, via []*Request) error


	// Jar specifies the cookie jar.

	//

	// The Jar is used to insert relevant cookies into every

	// outbound Request and is updated with the cookie values

	// of every inbound Response. The Jar is consulted for every

	// redirect that the Client follows.

	//

	// If Jar is nil, cookies are only sent if they are explicitly

	// set on the Request.

	Jar CookieJar


	// Timeout specifies a time limit for requests made by this

	// Client. The timeout includes connection time, any

	// redirects, and reading the response body. The timer remains

	// running after Get, Head, Post, or Do return and will

	// interrupt reading of the Response.Body.

	//

	// A Timeout of zero means no timeout.

	//

	// The Client cancels requests to the underlying Transport

	// as if the Request's Context ended.

	//

	// For compatibility, the Client will also use the deprecated

	// CancelRequest method on Transport if found. New

	// RoundTripper implementations should use the Request's Context

	// for cancellation instead of implementing CancelRequest.

	Timeout time.Duration

}
```
