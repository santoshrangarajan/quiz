package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

//var buf bytes.Buffer

var file *os.File
var (
	logger *log.Logger
)

// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go

func initQuiz() {
	var err error
	file, err = os.OpenFile("quiz.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Hello world")
}

func parseCsv() (map[string]int64, error) {
	f, err := os.Open("questions.csv")
	if err != nil {
		println("error opening csv")
		return nil, err
	}

	logger.Println("parseCsv. file opened:")
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		println("error reading csv")
		return nil, err
	}
	logger.Println("parseCsv. lineCount:", len(lines))
	m := make(map[string]int64)

	var line1 int64
	for _, line := range lines {
		logger.Println("line1", line1)
		if line1, err = strconv.ParseInt(line[1], 10, 64); err != nil {
			return nil, err
		}
		m[line[0]] = line1
	}

	return m, err
}

////// initialize reader at start
func readFromTerminal() (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Answer: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	i, err := strconv.ParseInt(text, 10, 64)
	return i, err

}

func startQuiz() (int, int) {
	m, err := parseCsv()
	if err != nil {
		panic("error parsing csv")
	}

	correct, wrong := 0, 0

	for key, element := range m {
		fmt.Println("question:" + key)
		answer, err := readFromTerminal()

		if err != nil {
			fmt.Println("wrong, errror")
			wrong++
		} else if answer == element {
			fmt.Println("correct")
			correct++
		} else {
			fmt.Println("wrong comparison")
			wrong++
		}

	}
	return correct, wrong
}

func main() {
	println("starting quiz.......")
	initQuiz()
	correct, wrong := startQuiz()
	fmt.Println("Results. Correct answers:", correct, "Wrong answers:", wrong)
	println("ending quiz.......")

}
