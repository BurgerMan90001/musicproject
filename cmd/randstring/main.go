package main

import (
	"fmt"

	"musicproject.com/pkg/randutil"
)

func main() {
	rand := randutil.RandRead{}
	fmt.Println(rand.New(39))
	//fmt.Println(randutil.FNV("testasdasdas"))
}
