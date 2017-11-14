package main

import (
	"github.com/axetroy/go-upload/config"
	"github.com/axetroy/go-upload/http"
)

func main() {

	var (
		err error
	)

	// init anything
	if err = config.Init(); err != nil {
		panic(err)
		return
	}

	if err = http.Init(); err != nil {
		panic(err)
		return
	}

	http.RunServer()
}
