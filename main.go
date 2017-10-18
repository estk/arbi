package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	port      = 8080
	sizeParam = "size"
	usage     = `Usage: curl <host>:<port>/bytes?size=<size>
	Where size = 1k|32k|128k|512k|1m|8m|32m

Example: curl localhost:8080/bytes?size=1k > 1k`
)
const (
	_      = iota
	kb int = 1 << (10 * iota)
	mb
)

func main() {
	fmt.Println(usage)
	listenOn := fmt.Sprintf(":%d", port)
	http.HandleFunc("/bytes", randomBits)
	log.Fatal(http.ListenAndServe(listenOn, nil))
}

func randomBits(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get(sizeParam)
	n, err := numBytes(size)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resBytes := make([]byte, n)
	rand.Read(resBytes)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(resBytes)
	if err != nil {
		log.Println("Unable to write random bytes")
	}
}

func numBytes(size string) (int, error) {
	switch size {
	case "1k":
		return kb, nil
	case "32k":
		return kb * 32, nil
	case "128k":
		return kb * 128, nil
	case "512k":
		return kb * 512, nil
	case "1m":
		return mb * 1, nil
	case "8m":
		return mb * 8, nil
	case "32m":
		return mb * 32, nil
	}
	return 0, errors.New("unable to get determine size")
}
