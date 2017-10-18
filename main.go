package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	port      = 8080
	sizeParam = "size"
	unitParam = "unit"
	usage     = `Usage: curl <host>:<port>/bytes?size=<uint>&unit=<unit>
	Where unit = b|k|m

Example: curl localhost:8080/bytes?size=1&unit=k > 1k`
)
const (
	B       = iota
	KB uint = 1 << (10 * iota)
	MB
)

func main() {
	fmt.Println(usage)
	listenOn := fmt.Sprintf(":%d", port)
	http.HandleFunc("/bytes", randomBits)
	log.Fatal(http.ListenAndServe(listenOn, nil))
}

func randomBits(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	size := q.Get(sizeParam)
	unit := q.Get(unitParam)
	n, err := numBytes(size, unit)
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

func numBytes(size, unit string) (uint, error) {
	var base uint
	switch unit {
	case "b":
		base = B
	case "k":
		base = KB
	case "m":
		base = MB
	}

	m, err := strconv.ParseUint(size, 10, 32)
	if err != nil {
		return 0, errors.New("unable to get determine size")
	}
	n := base * uint(m)
	if n > 512*MB {
		return 0, errors.New("too many bytes requested")
	}
	return n, nil
}
