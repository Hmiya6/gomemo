# `bufio`

## バッファとは
入出力を一時的に保存する部分.

サイズの大きいファイルは何度も入出力を行う必要がある.

そこで入出力回数が多すぎると、時間がかかる. = パフォーマンスの低下  
そのためにバッファを用意して**IOを削減**する.

なら、バッファは大きい方がいいのか? -> 一概にそうとは言えない.  
バッファが大きすぎると、大量の**メモリを消費**する.

バッファサイズは **パフォーマンスとメモリを利益均衡**をはかって見極める必要がある.  

## `bufio.Scanner`

## `bufio.Writer`
### `bufio.NewWriter`
既存の `io.Writer` からバッファ付き writer をつくる.
### `bufio.Write`
実装
```go
func (b *Writer) Write(p []byte) (nn int, err error)
```
バイト列 `p` から、バッファ付き writer ( = `bufio.Writer`) `b` の**バッファへ**書き込みを実行する.

## 参考
Golangのバッファってよくできてるよな  
https://note.crohaco.net/2019/golang-buffer/