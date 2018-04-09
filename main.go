package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

const fileDefault = "problems.csv"

var operators = map[string]func(int, int) int{
	"+": func(a, b int) int {
		return a + b
	},
}

func main() {
	filename := flag.String(
		"csv",
		fileDefault,
		"Specify name of the .csv file, where contains questions and answers.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	log.Print(file.Name())

	reg := regexp.MustCompile(`(\d+)(.)(\d+)`)

	content := csv.NewReader(file)
	for {
		record, err := content.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		opers := reg.FindAllStringSubmatch(record[0], -1)
		firstOperand, err := strconv.Atoi(opers[0][1])
		if err != nil {
			panic(err)
		}
		secondOperand, err := strconv.Atoi(opers[0][3])
		if err != nil {
			panic(err)
		}
		resultFunction := operators[opers[0][2]]
		log.Println(resultFunction(firstOperand, secondOperand))
	}
}
