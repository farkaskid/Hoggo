package writer

import (
	"Hoggo/types"
	"bufio"
	"bytes"
	"os"
)

func Write(value types.Type) int {
	workingDirectoryPath, _ := os.Getwd()

	file, _ := os.Create(workingDirectoryPath + "/db.hg")

	var buffer bytes.Buffer

	buffer.Write(value.Size().WriterView())
	buffer.Write(value.WriterView())

	writer := bufio.NewWriter(file)

	count, _ := writer.Write(buffer.Bytes())

	writer.Flush()

	return count
}
