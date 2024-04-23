package eval

import (
	"fmt"
	"strings"
	"testing"
	"text/scanner"
)

func TestScanner(t *testing.T) {

	var scan scanner.Scanner
	scan.Mode = scanner.ScanInts | scanner.ScanIdents
	scan.Init(strings.NewReader("aaa(1)1.1  11 A '1' \"111111\" \\aa"))
	token := scan.Scan()
	for token != scanner.EOF {
		fmt.Printf("%v,%s,%s\n", token, scanner.TokenString(token), scan.TokenText())
		token = scan.Scan()
	}

}
