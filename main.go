package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Wost struct {
	Start       time.Time
	End         time.Time
	Description string
}

func (wost Wost) String() string {
	return fmt.Sprintf("Start: %s\nEnd: %s\nDescription: %s\nTime Spent: %s", wost.Start, wost.End, wost.Description, wost.timeSpent())
}

func (wost Wost) timeSpent() time.Duration {
	return wost.End.Sub(wost.Start)
}

func WostTotTime(wosts []Wost) time.Duration {
	var tot time.Duration
	for _, wost := range wosts {
		tot += wost.timeSpent()
	}
	return tot
}

func WostNew(start time.Time, end time.Time, description string) Wost {
	return Wost{start, end, description}
}

func ReadWost(file string) []Wost {
	lines, err := ReadCsv(file)
	if err != nil {
		panic(err)
	}

	var wosts []Wost
	for _, line := range lines {
		start, err := time.Parse(time.UnixDate, fmt.Sprintf("%s %s %s %s", line[0], line[3], line[1], line[2]))
		if err != nil {
			panic(err)
		}
		end, err := time.Parse(time.UnixDate, fmt.Sprintf("%s %s %s %s", line[0], line[4], line[1], line[2]))
		if err != nil {
			panic(err)
		}

		data := WostNew(start, end, line[5])
		wosts = append(wosts, data)
	}
	return wosts
}

func ReadCsv(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	// skip header
	reader.Read()
	return reader.ReadAll()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wost <csv_file>")
		os.Exit(1)
	}
	wosts := ReadWost(os.Args[1])
	for _, wost := range wosts {
		fmt.Println("------")
		fmt.Println(wost)
	}
	fmt.Println("Total time spent:", WostTotTime(wosts))
}
