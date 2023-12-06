package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type key struct {
	x, y int
}

var directions = [8][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func DoPartOne() {
	total := 0

	file, err := os.Open("./input3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numbers := make(map[key]*int)
	numberFinder := regexp.MustCompile("\\d+")

	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		indexes := numberFinder.FindAllIndex(line, len(line))

		for _, idx := range indexes {
			start := idx[0]
			end := idx[1]

			calibrationValue, err := strconv.Atoi(string(line[start:end]))

			if err != nil {
				panic(err)
			}

			for i := start; i < end; i++ {
				numbers[key{x: index, y: i}] = &calibrationValue
			}
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)

	index = 0
	for scanner.Scan() {
		line := scanner.Bytes()

		for y, b := range line {
			if b != '.' && !numberFinder.MatchString(string(b)) {
				for _, direction := range directions {
					cal, ok := numbers[key{x: index + direction[0], y: y + direction[1]}]

					if ok {
						fmt.Println(string(b), *cal)
						total = *cal + total
						*cal = 0
					}
				}

			}
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("The answer for day 3 is: %d\n", total)
}
