package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// the 'commands' we can execute
	addrepo := flag.Bool("addrepo", false, "add repository to Rprofile.site")
	install := flag.Bool("install", false, "install and compile and R")
	// arguments used by commands
	rversion := flag.String("rversion", "3.5.2", "the R version")
	repo := flag.String("repo", "https://inwt-vmeh2.inwt.de/r-repo", "a cran like repository")
	mran := flag.String("mran", "", "the date of the MRAN mirror to use")
	flag.Parse()
	if *addrepo {
		cmdaddrepo(*rversion, *repo)
	}
	if *install {
		cmdinstall(*rversion, *mran)
	}
}

func cmdinstall(rversion, mran string) {
	mran, err := makemran(rversion, mran)
	if err != nil {
		log.Fatal(err)
	}
	if err := downloadFile("/tmp/rbt-install-r.sh", "https://raw.githubusercontent.com/INWTlab/r-config/master/bionic/install-r.sh"); err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/rbt-install-r.sh")
	execute("bash", "/tmp/rbt-install-r.sh", rversion, mran)
}

func cmdaddrepo(rversion string, repo string) {
	file := newRfiles(rversion).rprofile.site
	options := "options(repos = c(options(\"repos\"), \"" + repo + "\"))"
	if lineInFile(file, options) {
		log.Println("line already in file: not appending")
	} else {
		appendLine(file, options)
	}
}
