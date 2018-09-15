package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const fileDefault = "problems.csv"

func main() {
	filename := flag.String(
		"csv",
		fileDefault,
		"Specify name of the .csv file, where contains questions and answers.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	content := csv.NewReader(file)
	total, score := 0, 0
	for {
		record, err := content.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s = ", record[0])
		var answer int
		fmt.Scanln(&answer)
		if strconv.Itoa(answer) == record[1] {
			score++
		}

		total++
	}

	fmt.Printf("Your score %d of %d\n", score, total)
}
