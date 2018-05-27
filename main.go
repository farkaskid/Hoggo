package main

import (
	"Hoggo/types"
	"fmt"
)

func main() {
	text1 := types.Text{types.New([]byte("this is a "))}
	text2 := types.Text{types.New([]byte("car"))}

	fmt.Println(text1.Concat(text2).StringView())
}
