package config

import "os"

type GlobalPaths struct {
	root string
	cwd  string
}

var Paths GlobalPaths

/**
get current work dir
 */
func getCwd() (pwd string, err error) {
	pwd, err = os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	return
}

func InitPaths() {
	var (
		cwd, _ = getCwd()
	)

	Paths = GlobalPaths{
		cwd: cwd,
	}
}
