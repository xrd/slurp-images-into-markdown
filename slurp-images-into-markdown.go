package main

import (
	"fmt"
	"path/filepath"
	"os"
	"flag"
	"strconv"
	"image/png"
	"github.com/nfnt/resize"
	"log"
)


func resizePng( f string, i int, size int ) {
	// open "test.jpg"
	file, err := os.Open( f )
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
	m := resize.Resize( uint(size), 0, img, resize.Lanczos3)
	
	output := "./" + strconv.Itoa(i) + ".resized.png"
	fmt.Println( "Writing: ", output )
	out, err := os.Create( output )
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	
	// write new image to file
	png.Encode(out, m )
}


func main() {
	if 3 > len(os.Args)  {
		fmt.Println( "Need to specify a directory+pattern to use and output size." )
	} else {
		format := flag.String( "--format", nil, "Filename glob, e.g.: ~/Desktop/Screen\ Shot\ 2016-12-01" ) // os.Args[1]
		size := flat.Int( "--size", 800, "Resize width, defaults to 800" )
		// size, err := strconv.Atoi( os.Args[2] )
		if nil != err {
			fmt.Println( "Second argument conversion error: ", err )
		} else {
			files, _ := filepath.Glob(format+"*")
			if 4 == len(os.Args) {
				if "--channels" == os.Args[3] {
					fmt.Println( "Using channels" )
				} else if "--goroutine" == os.Args[3] {
					fmt.Println( "Using goroutine" )
					for index,f := range files {
						go resizePng( f, index, size )
					}
				}
			} else {
				files, _ := filepath.Glob(format+"*")
				for index,f := range files {
					resizePng( f, index, size )
				}
			}
		}
		// fmt.Println(files) // contains a list of all files in the current directory
	}
	fmt.Println( "Done" );
}

