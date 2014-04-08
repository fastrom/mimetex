package main

import (
	"flag"
	"log"
	"os"

	"github.com/xlab/mimetex-go"
)

var size = flag.Int("s", 2, "set desired size in points (1-12)")
var out = flag.String("o", "out.gif", "set output file name")

func main() {
	flag.Parse()
	expr := flag.Arg(0)
	data, err := mimetex.RenderGifBytes(expr, *size) // render
	if err != nil {
		log.Fatalln("mimetex error: " + err.Error())
	}
	log.Printf("rendered image (%d bytes)\n", len(data))

	f, err := os.Create(*out)
	if err != nil {
		log.Fatalln("error creating file: " + *out)
	}
	defer f.Close()

	_, err = f.Write(data) // save gif
	if err != nil {
		log.Fatalln("error writing to file: " + *out)
	}
}
