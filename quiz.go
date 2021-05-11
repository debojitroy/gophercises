package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Quiz game...")
	loadFile := flag.String("file", "payload.csv", "CSV File with input")
	flag.Parse()

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
	wrong :=0

	for i:=0; i < len(dataArray); i++ {
		fmt.Println(dataArray[i][0])
		var answer string
		fmt.Scanln(&answer)

		if answer == dataArray[i][1] {
			correct++
		} else {
			wrong++
		}
	}

	fmt.Println("Correct: ", correct)
	fmt.Println("Wrong: ", wrong)
}