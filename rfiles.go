package main

import (
	"log"
	"os"
)

type rfiles struct {
	rversion string
	lib      string
	rprofile rprofile
}

type rprofile struct {
	site string
}

func newRprofile(lib string) rprofile {
	site := lib + "R/etc/Rprofile.site"
	if _, err := os.Stat(site); os.IsNotExist(err) {
		log.Fatal(err)
	}
	return rprofile{site}
}

func newRfiles(rversion string) rfiles {
	lib := "/usr/local/lib/R/" + rversion + "/lib/"
	if _, err := os.Stat(lib); os.IsNotExist(err) {
		log.Fatal(err)
	}
	return rfiles{rversion, lib, newRprofile(lib)}
}
