package http

import (
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/misc"
	"io/ioutil"
	"net/http"
	"time"
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

	logger.Get().Infof("%d bytes received, http code %v", len(body), resp.StatusCode)

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
