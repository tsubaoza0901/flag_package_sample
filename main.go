package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flags := flag.NewFlagSet("sub", flag.ContinueOnError)
	filePath := flags.String(
		"f",            // name（設定するflag name）
		"/tmp/",        // value（flagが指定されなかった場合のdefault value）
		"path to file", // usage（使用方法説明）
	)

	flags.Parse(os.Args[2:])
	fmt.Println("file path ", *filePath)
}
