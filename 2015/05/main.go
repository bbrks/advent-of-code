package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func numOfVowels(inpt *string) int {
	var n int
	for _, c := range *inpt {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			n++
		}
	}
	return n
}

func twoInARow(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func noNaughtyStrings(inpt *string) bool {
	var n int
	naughtyStrings := []string{"ab", "cd", "pq", "xy"}
	for _, s := range naughtyStrings {
		n = n + strings.Count(*inpt, s)
	}
	return n == 0
}

func isNice(inpt *string) bool {
	ret := numOfVowels(inpt) >= 3 &&
		twoInARow(*inpt) &&
		noNaughtyStrings(inpt)
	return ret
}

func twoPairs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Count(s, s[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func middleRepeat(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func isNice2(inpt *string) bool {
	ret := twoPairs(*inpt) &&
		middleRepeat(*inpt)
	return ret
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var niceStrings, niceStrings2 int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if isNice(&line[0]) {
			niceStrings++
		}
		if isNice2(&line[0]) {
			niceStrings2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of nice strings: %d\n", niceStrings)
	fmt.Printf("Number of new nice strings: %d\n", niceStrings2)
}
