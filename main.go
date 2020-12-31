package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(
		&name,                          // flagのValueを格納する変数
		"name",                         // name（設定するflag name）
		"Tom",                          // value（flagが指定されなかった場合のdefault value）
		"Please specify -name <Value>", // usage（使用方法説明）
	)
	flag.Parse()
	fmt.Println("The name flag is", name)
}
