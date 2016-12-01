package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}

func resizePng(f string, size *int, output string) {
	// open "test.jpg"
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(*size), 0, img, resize.Lanczos3)

	// fmt.Println("Writing: ", output)
	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

const formatUsage = "Filename glob, e.g.: ~/Desktop/Screen\\ Shot\\ 2016-12-01"

func main() {
	format := flag.String("format", "", formatUsage)
	size := flag.Int("size", 800, "Resize width")
	convert := flag.Bool("convert", false, "Resize files")
	flag.Parse()

	if "" != *format {
		files, _ := filepath.Glob(*format + "*")
		for index, f := range files {
			short := strconv.Itoa(index) + ".resized.png"
			output := "./" + short
			if *convert {
				resizePng(f, size, output)
			} else {
				copy(f, output)
			}
			// Print out the files
			fmt.Printf( "## %d\n\n[](%v)\n\n", index, short )
		}
	} else {
		fmt.Println("Use -format: " + formatUsage)
	}
	// fmt.Println("Done")
}
