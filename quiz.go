package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	loadFile := flag.String("file", "payload.csv", "CSV File with input")
	timeout := flag.Int("timeout", 30, "Timeout for the game")
	flag.Parse()

	fmt.Println("Welcome to the Quiz Game")
	fmt.Println("Timeout for the game is: ", *timeout)

	fmt.Println("Input Filename: ", *loadFile)

	fileData, err := ioutil.ReadFile(*loadFile)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		panic(err)
	}

	csvReader := csv.NewReader(strings.NewReader(string(fileData)))
	dataArray := make([][]string, 0, 1)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		dataArray = append(dataArray, record)
	}

	correct := 0
	wrong := 0

	fmt.Println("Press Enter to start....")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Time up !!!")
		fmt.Println("Correct: ", correct)
		fmt.Println("Wrong: ", wrong)
		os.Exit(0)
	}()

	for i := 0; i < len(dataArray); i++ {
		fmt.Println(dataArray[i][0])
		var answer string
		fmt.Scanln(&answer)

		if answer == dataArray[i][1] {
			correct++
		} else {
			wrong++
		}
	}

	timer.Stop()
	fmt.Println("Correct: ", correct)
	fmt.Println("Wrong: ", wrong)
}
