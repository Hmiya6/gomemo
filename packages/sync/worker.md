# worker をつくって並行処理

## サンプルコード
`worker` を 10 つくって、`ch` の処理を行うサンプル.
```go
func worker(ch *chan int, wg *sync.WaitGroup) {
    defer wg.Done() // ... (1)
    for i := range ch {
        // Do something
    }
}

func main() {

    wg := &sync.WaitGroup{}
    ch := make(chan int, 10)
    for i := 0; i<10; i++ {
        wg.Add(1) // ... (2)
        go worker(&ch, wg)
    }

    for i:=0;i<100000; i++ {
        ch <- i // ... (3)
    }
    close(ch) // ... (4)
    wg.Wait() // ... (5)
}
```
(1) 必ず `Done` する. でないと (5) が待ち続ける.  
(2) goroutine を起動する前に `Add` する. 他の場所(e.g. `worker` 内、`go worker` の後)だと、動作の順序が保障されず危険 (e.g. `Add` 前に `Done` が起こる. (5) の後に `Add` がスケジューリングされる)  
(3) 処理対象を worker へ渡すための `ch`.  
(4) 必ず `close` する. でないと `worker` が `ch` の受信を待ち続ける.  
(5) すべてで `Done` されるまで待つ.
