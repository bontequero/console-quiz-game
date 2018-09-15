package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	fileDefault    = "problems.csv"
	timerDefault   = 30
	shuffleDefault = false
)

func main() {
	filename := flag.String(
		"csv",
		fileDefault,
		"Specify name of the .csv file, where contains questions and answers.",
	)
	timeLimit := flag.Int(
		"timer",
		timerDefault,
		"Specify time in seconds to answer quiz questions.",
	)
	shouldShuflle := flag.Bool(
		"shuffle",
		shuffleDefault,
		"Specify this to shuffle order of quiz questions.",
	)
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *shouldShuflle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(content), func(i, j int) {
			content[i], content[j] = content[j], content[i]
		})
	}

	fmt.Print("Press Enter to start quiz")
	fmt.Scanf("\n")

	score := 0
	end := make(chan int)
	go func() {
		var answer int
		for _, record := range content {
			fmt.Printf("%s = ", record[0])
			fmt.Scanln(&answer)
			if strconv.Itoa(answer) == record[1] {
				score++
			}
		}
		end <- score
	}()

	select {
	case <-time.After(time.Duration(*timeLimit) * time.Second):
	case <-end:
	}

	fmt.Printf("\nYour score %d of %d\n", score, len(content))
}
