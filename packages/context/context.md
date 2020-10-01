# `context`

`http.Request` などに様々な情報を付け加える.  
e.g. デッドラインやタイムアウト、キャンセルシグナル、その他リクエスト内の値

## 例
```go
// username = "foo" を context を用いて http.Request に添付する.
func addCtx(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    username := "foo"
    
    // r の context へ username を追加する
    ctx := context.WithValue(r.Context(), "username", username)
    
    // r.Context を新しい context である ctx にする.
    r := r.WithContext(ctx)

    // 次の処理へ送る
    next(w, r)
}
```

## `context.Background`
基本的な `context`.  
新たに `context` をつくるときに使う

## 