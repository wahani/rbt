package main

import "errors"

func makemran(rversion, mran string) (string, error) {
	if mran != "" {
		return mran, nil
	}
	if rversion == "3.5.2" {
		return "2019-03-10", nil
	}
	return "", errors.New("unable to determine mran mirror")
}
