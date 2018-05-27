package main

import (
	"Hoggo/types"
	"Hoggo/writer"
	"fmt"
)

func main() {
	text := types.Text{"a "}

	count := writer.Write(text)

	fmt.Println(count, "bytes written")
}
