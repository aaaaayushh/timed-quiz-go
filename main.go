package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Declare a string flag with specified name('csv'), default value('problems.csv'), and usage description.
	// The return value is the address of a string variable in which to store the value of the flag.
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// Parse() is called to actually parse the command line arguments. It must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

// exit prints the message and exits the program with status code 1.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
