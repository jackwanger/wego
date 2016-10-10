package main

import "fmt"

var githash string
var buildstamp string

func version() {
	fmt.Printf("Git Hash   : %s\n", githash)
	fmt.Printf("Build Time : %s\n", buildstamp)
}
