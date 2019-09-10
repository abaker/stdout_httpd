package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	port := flag.Int("port", 80, "port")
	debug := flag.Bool("debug", false, "debug logging")
	flag.Parse()
	mutex := &sync.Mutex{}
	var last string

	if !*debug {
		log.SetOutput(ioutil.Discard)
	}

	go func() {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			line := scan.Text()
			mutex.Lock()
			last = line
			mutex.Unlock()
			log.Println(line)
		}
		log.Fatal("stdin closed")
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		body := last
		mutex.Unlock()
		if body == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, err := fmt.Fprintln(w, body)
			if err != nil {
				log.Println(err)
			}
		}
		log.Println(r.RemoteAddr, r.Method, r.RequestURI, "\u21a9", body)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
