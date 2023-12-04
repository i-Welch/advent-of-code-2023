package day2

// Example game line
// Game 1: 3 blue, 7 green, 10 red; 4 green, 4 red; 1 green, 7 blue, 5 red; 8 blue, 10 red; 7 blue, 19 red, 1 green

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func DoPartOne() {
	total := 0
	file, err := os.Open("./input2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for true {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		game := strings.Split(line, ":")

		gameId, err := strconv.Atoi(strings.Split(game[0], " ")[1])

		if err != nil {
			panic(err)
		}

		rounds := strings.Split(game[1], ";")

		isValid := true
		for _, round := range rounds {
			colorCounts := strings.Split(round, ",")

			for _, colorCount := range colorCounts {
				colorAndCount := strings.Split(strings.Trim(colorCount, " \n"), " ")

				count, _ := strconv.Atoi(colorAndCount[0])
				color := colorAndCount[1]

				if color == "red" && count <= 12 {
					continue
				}

				if color == "green" && count <= 13 {
					continue
				}

				if color == "blue" && count <= 14 {
					continue
				}

				isValid = false
				break
			}

			if !isValid {
				break
			}
		}

		if isValid {
			total += gameId
		}
	}

	fmt.Println("The solution for part 1 is:", total)
}

func DoPartTwo() {
	total := 0
	file, err := os.Open("./input2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for true {
		line, readErr := reader.ReadString('\n')

		if readErr != nil && readErr != io.EOF {
			panic(readErr)
		}

		game := strings.Split(line, ":")

		gameId, err := strconv.Atoi(strings.Split(game[0], " ")[1])

		if err != nil {
			panic(err)
		}

		rounds := strings.Split(game[1], ";")

		minColors := map[string]int{}

		for _, round := range rounds {
			colorCounts := strings.Split(round, ",")

			for _, colorCount := range colorCounts {
				colorAndCount := strings.Split(strings.Trim(colorCount, " \n"), " ")

				count, _ := strconv.Atoi(colorAndCount[0])
				color := colorAndCount[1]

				if minColors[color] < count {
					minColors[color] = count
				}
			}
		}

		fmt.Printf("gameId: %d - power: %d\n", gameId, minColors["red"]*minColors["blue"]*minColors["green"])

		total += minColors["red"] * minColors["blue"] * minColors["green"]

		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("The solution for part 2 is:", total)
}
