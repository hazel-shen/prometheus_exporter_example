package main
import (
    "fmt"
    "net/http"
    "log"
)
type Counter struct{
    count int64
}

var counter  = new(Counter)

func (c *Counter) Add(count int64) int64 {
    c.count += count
    return c.count
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "lexporter_request_count{user=\"admin\"} %d", counter.Add(10))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You are the %d guests", counter.Add(10))

}

func main () {
    http.HandleFunc("/", MetricsHandler)
    http.HandleFunc("/hello", HelloHandler)
    log.Print("Start webserver..")
    http.ListenAndServe(":8000", nil)
}

