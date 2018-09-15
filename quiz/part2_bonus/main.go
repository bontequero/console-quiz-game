package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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

	answers := make(chan struct{})
	go func() {
		var answer string
		for _, record := range content {
			fmt.Printf("%s = ", record[0])
			fmt.Scanln(&answer)
			if strings.ToLower(strings.TrimSpace(answer)) ==
				strings.ToLower(strings.TrimSpace(record[1])) {
				answers <- struct{}{}
			}
		}
		close(answers)
	}()

	correctCount := make(chan int)
	go func() {
		score := 0
		for {
			select {
			case <-time.After(time.Duration(*timeLimit) * time.Second):
				correctCount <- score
				return
			case _, received := <-answers:
				if !received {
					correctCount <- score
					return
				}
				score++
			}
		}
	}()

	select {
	case score := <-correctCount:
		fmt.Printf("\nYour score %d of %d\n", score, len(content))
	}
}
