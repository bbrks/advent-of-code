package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	val, err := partOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part one: %d\n", val)

	val, err = partTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part two: %d\n", val)
}

func partOne(input string) (int, error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
updates:
	for _, pages := range updates {
		//pages:
		for i, page := range pages {
			if i == 0 {
				// skip first
				continue
			}
			prevPage := pages[i-1]
			for _, rule := range rules {
				if rule.a == prevPage && rule.b == page {
					// correct order, move on to next page
					break
				} else if rule.a == page && rule.b == prevPage {
					// incorrect order - skip update completely
					continue updates
				}
			}
		}
		middlePageIdx := len(pages) / 2
		val += pages[middlePageIdx]
	}
	return val, nil
}

func partTwo(input string) (int, error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
	for _, pages := range updates {
		updateCorrect := true
	redoRules:
		for i := range pages {
			if i == 0 {
				// skip first
				continue
			}
			page := pages[i]
			prevPage := pages[i-1]
			for _, rule := range rules {
				if rule.a == prevPage && rule.b == page {
					// correct order - skip
					break
				} else if rule.a == page && rule.b == prevPage {
					updateCorrect = false
					pages[i], pages[i-1] = prevPage, page
					goto redoRules // ... but does it work?
				}
			}
		}
		if updateCorrect {
			continue
		}
		middlePageIdx := len(pages) / 2
		val += pages[middlePageIdx]
	}
	return val, nil
}

type rule struct {
	a, b int
}

func parseInput(input string) ([]rule, [][]int, error) {
	rules := make([]rule, 0)
	updates := make([][]int, 0)

	parts := strings.SplitN(input, "\n\n", 2)
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid input")
	}

	ruleScanner := bufio.NewScanner(strings.NewReader(parts[0]))
	for ruleScanner.Scan() {
		var a, b int
		_, err := fmt.Sscanf(ruleScanner.Text(), "%d|%d", &a, &b)
		if err != nil {
			return nil, nil, err
		}
		rules = append(rules, rule{a, b})
	}

	pageScanner := bufio.NewScanner(strings.NewReader(parts[1]))
	for pageScanner.Scan() {
		pages := make([]int, 0)
		for _, num := range strings.Split(pageScanner.Text(), ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
			pages = append(pages, n)
		}
		updates = append(updates, pages)
	}

	return rules, updates, nil
}
