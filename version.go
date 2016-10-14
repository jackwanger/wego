package main

import "fmt"

var (
	version    string
	githash    string
	buildstamp string
)

func init() {
	fmt.Printf("Version    : %s\n", version)
	fmt.Printf("Git Hash   : %s\n", githash)
	fmt.Printf("Build Time : %s\n", buildstamp)
}
