package misc

import "log"

func Check(e error) {
	if e != nil {
		log.Panic(e)
		panic(e)
	}
}
