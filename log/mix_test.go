package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	log.Println("hello")

	log.SetPrefix("aaa:")

	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	log.Println("hello")

	//log.Panicln("hel")

	log.Fatalln("faa")

}

func TestNewLog(t *testing.T) {
	bf := &bytes.Buffer{}
	nLog := log.New(bf, "", log.Ltime)

	nLog.Println("hello new self log")

	fmt.Println("print:" + bf.String())
}
func TestNewLog2(t *testing.T) {

	create, err := os.Create("./testlog.txt")

	if err != nil {
		return
	}

	bf := &bytes.Buffer{}
	nLog := log.New(io.MultiWriter(bf, create, os.Stdout), "", log.Ltime)

	nLog.Println("hello new self log")

	fmt.Println("print:" + bf.String())
}
