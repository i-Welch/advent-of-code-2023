package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var validDigits = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func isValidDigit(maybeChar rune) bool {
	for i := 0; i < 10; i++ {
		if validDigits[i] == maybeChar {
			return true
		}
	}

	return false
}

func DoPartOne() {
	total := 0
	file, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var first, last rune
	newline := true

	addFirstLastToTotal := func() {
		digit, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic(err)
		}
		total += digit
	}

	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			addFirstLastToTotal()
			break
		}

		if char == '\n' {
			addFirstLastToTotal()
			newline = true
			continue
		}

		if isValidDigit(char) && newline {
			first = char
			last = char
			newline = false
			continue
		}

		if isValidDigit(char) {
			last = char
		}
	}

	fmt.Println("The final value for part 1 is:", total)
}

var digitMap = map[string]rune{
	"0":     '0',
	"1":     '1',
	"2":     '2',
	"3":     '3',
	"4":     '4',
	"5":     '5',
	"6":     '6',
	"7":     '7',
	"8":     '8',
	"9":     '9',
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func DoPartTwo() {
	total := 0
	file, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	rgx, err := regexp.Compile("\\d|zero|one|two|three|four|five|six|seven|eight|nine")

	if err != nil {
		panic(err)
	}

	revRgx, err := regexp.Compile("\\d|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")

	if err != nil {
		panic(err)
	}

	for true {
		line, err := reader.ReadBytes('\n')

		if err != nil && err != io.EOF {
			panic(err)
		}

		firstDigit := rgx.FindString(string(line))

		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}

		lastDigit := revRgx.Find(line)

		for i, j := 0, len(lastDigit)-1; i < j; i, j = i+1, j-1 {
			lastDigit[i], lastDigit[j] = lastDigit[j], lastDigit[i]
		}

		digit, _ := strconv.Atoi(string(digitMap[firstDigit]) + string(digitMap[string(lastDigit)]))
		total += digit

		if err == io.EOF {
			break
		}
	}

	fmt.Println("The final value for part 2 is:", total)
}
