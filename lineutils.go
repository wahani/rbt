package main

import (
	"bufio"
	"log"
	"os"
)

func appendLine(file, line string) {
	log.Println("Adding:\n'" + line + "' to '" + file + "'")
	f := fileOpen(file)
	fileWrite(f, line+"\n")
	fileClose(f)
}

func lineInFile(file string, line string) bool {
	f := fileOpenRead(file)
	defer fileClose(f)
	scanner := bufio.NewScanner(f)
	var ind bool
	ind = false
	for scanner.Scan() {
		tmp := scanner.Text()
		if tmp == line {
			ind = true
			break
		}
	}
	return ind
}

func fileOpenRead(file string) *os.File {
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	check(err)
	return f
}

func fileOpen(file string) *os.File {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_RDONLY
	check(err)
	return f
}

func fileWrite(f *os.File, line string) {
	_, err := f.WriteString(line)
	check(err)
}

func fileClose(f *os.File) {
	err := f.Close()
	check(err)
}
