package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimter
	reps, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer reps.Body.Close()

	if reps.StatusCode != http.StatusOK {
		// fmt.Println("Error: Status Code: ", reps.StatusCode)
		return nil, fmt.Errorf("wrong status Code: %d", reps.StatusCode)
	}

	e := determineEncoding(reps.Body)
	utf8Reader := transform.NewReader(reps.Body,
		e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
