package misc

import (
	"github.com/pawski/go-xchange/internal/logger"
)

func Check(e error) {
	if e != nil {
		logger.Get().Panic(e)
		panic(e)
	}
}
