package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	skelDir := flag.String("d", ".SKELETON", "Skeleton directory")
	placeHolder := flag.String("p", "SKELETON", "Placeholder string")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <DAY_NUMBER>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	dayNumber := flag.Arg(0)
	if dayNumber == "" {
		flag.Usage()
		os.Exit(2)
	}

	destDir := "day" + dayNumber
	err := os.Mkdir(destDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	files, err := ioutil.ReadDir(*skelDir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		var (
			sourcePath = filepath.Join(*skelDir, file.Name())
			destPath   = filepath.Join(
				destDir,
				strings.Replace(file.Name(), *placeHolder, dayNumber, -1),
			)
		)

		contents, err := ioutil.ReadFile(sourcePath)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(sourcePath, "->", destPath)
		err = ioutil.WriteFile(
			destPath,
			bytes.Replace(contents, []byte(*placeHolder), []byte(dayNumber), -1),
			0644,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
