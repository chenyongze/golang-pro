package main
import (
    "net/http"
    "log"
    "fmt"
    "runtime"
)
func handler512(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "keep-alive")
    a := []byte("aaaaa...512...aaa")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", len(a)))
    w.Write(a)
}
func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    http.HandleFunc("/512b", handler512)
    log.Fatal(http.ListenAndServe(":9080", nil))
}
