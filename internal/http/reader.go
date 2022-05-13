package http

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pawski/go-xchange/internal/logger"
	"github.com/pawski/go-xchange/internal/misc"
)

var buff []byte

func GetUrl(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		logger.Get().Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		logger.Get().Print(err)
	} else {
		buff = body
	}

	logger.Get().Printf("%d bytes", len(body))
	logger.Get().Print(resp.Status)

	return body
}

func FlushBufferToFile() {
	writeToFile(buff)
}

func writeToFile(body []byte) {
	t := time.Now()

	err := ioutil.WriteFile("./cache/"+t.String()+".html", body, 0644)
	misc.Check(err)
}
