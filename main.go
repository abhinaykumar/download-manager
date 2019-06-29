package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"./downloader"
	"./stats"
	"github.com/gorilla/mux"
)

func main() {
	// Log configs
	logPath := "development.log"
	openLogFile(logPath)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()
	mount(router, "/", downloader.Router())
	mount(router, "/stats", stats.Router())
	log.Fatal(http.ListenAndServe(":3000", RequestLogger(router)))
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}
		log.SetOutput(lf)
	}
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
		log.Printf("pid:[%d]\t\t%s\t\t%s\t\t%s\t\t%v\n", os.Getpid(), r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}
