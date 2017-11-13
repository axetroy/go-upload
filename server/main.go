package main

import (
	"github.com/axetroy/go-upload/server/config"
	"github.com/axetroy/go-upload/server/http"
)

func main() {

	var (
		err error
	)

	// init anything
	if err = config.Init(); err != nil {
		panic(err)
	}

	if err = http.Init(); err != nil {
		panic(err)
	}

	http.RunServer()
}
