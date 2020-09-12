# path/filepath のメモ

## 1. パスの一部を抜き出す
* `filepath.Base()`
* `filepath.Dir()`
* `filepath.Ext()`

### `filepath.Base()`
`path` の最後の要素を返す。  
```go
func Base(path string) string
```

e.g.
```go
filepath.Base("foo/bar/baz.js")  // == baz.js
```

### `filepath.Dir()`
`path` のディレクトリ部分(最後以外の要素)を返す
```go
func Dir(path string) string
```

e.g.  
// は / として返される
```go
filepath.Dir("/foo//bar/baz.js") // == /foo/bar
```

### `filepath.Ext()`
拡張子を返す
```go
func Ext(path string) string
```

e.g.
```go
filepath.Ext("/foo/bar/baz.js") // == .js
```
---
