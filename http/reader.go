package http

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/pawski/go-xchange/misc"
	"time"
	"strconv"
)

var buff []byte

func GetUrl(url string) ([]byte) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		log.Print(err)
	} else {
		//writeToFile(body)
		buff = body
	}

	log.Printf("%d bytes", len(body))
	log.Print(resp.Status)

	return body
}

func FlushBufferToFile()  {
	writeToFile(buff)
}

func writeToFile(body []byte) {
	t := time.Now()

	err := ioutil.WriteFile("./cache/"+strconv.FormatInt(t.Unix(), 10)+".html", body, 0644)
	misc.Check(err)
}
