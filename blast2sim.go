package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func mustParseInt(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	r := csv.NewReader(os.Stdin)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// 読み込みエラー発生
			fmt.Println("Read error: ", err)
			break
		}

		query :=

			fmt.Printf("日本語名[%s] 英語名[%s] 科分類[%s]\n", record[0], record[1], record[2])
	}
}
