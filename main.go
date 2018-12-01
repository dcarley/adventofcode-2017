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
		fmt.Fprintf(os.Stderr, "usage: %s <YEAR_NUMBER> <DAY_NUMBER>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	yearNumber, dayNumber := flag.Arg(0), flag.Arg(1)
	if yearNumber == "" || dayNumber == "" {
		flag.Usage()
		os.Exit(2)
	}

	destDir := fmt.Sprintf("%s/day%s", yearNumber, dayNumber)
	err := os.MkdirAll(destDir, 0755)
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

		contents = bytes.Replace(contents, []byte("{{year}}"), []byte(yearNumber), -1)
		contents = bytes.Replace(contents, []byte("{{day}}"), []byte(dayNumber), -1)

		fmt.Println(sourcePath, "->", destPath)
		err = ioutil.WriteFile(destPath, contents, 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
