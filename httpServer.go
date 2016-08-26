package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	var keepAliveMins int
	if args[1] != "" {
		keepAliveMins, _ = strconv.Atoi(args[1])
	} else {
		keepAliveMins = 10
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))
	fmt.Fprintf(os.Stdout, "server url is %s", ts.URL)
	time.Sleep(time.Duration(keepAliveMins) * time.Minute)
	defer ts.Close()
}
