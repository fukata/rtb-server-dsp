package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
    "flag"
)

func handler(w http.ResponseWriter, r *http.Request) {
    id               := r.FormValue("id")
    sleep_seconds, _ := strconv.Atoi( r.FormValue("t") )
    status, _        := strconv.Atoi( r.FormValue("s") )
    price            := 0
    if r.FormValue("p") != "" {
        price, _ = strconv.Atoi( r.FormValue("p") )
    }

    if sleep_seconds > 0 {
        time.Sleep( time.Second * time.Duration( sleep_seconds ) ) // milliseconds -> seconds
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"id\":\"%s\",\"status\":%d,\"price\":%d}", id, status, price)
}

func main() {
    port := flag.Int("port", 8080, "PORT")
    flag.Parse()

    http.HandleFunc("/ad", handler)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
