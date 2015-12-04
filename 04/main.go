package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

func adventCoin(inpt *string, p string) int {
	var val int
	var hash string

	for !strings.HasPrefix(hash, p) {
		val = val + 1
		data := []byte(fmt.Sprintf("%s%d", *inpt, val))
		hash = fmt.Sprintf("%x", md5.Sum(data))
	}

	return val
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		ret := adventCoin(&line[0], line[1])
		fmt.Printf("Advent Coin [%s] Number for %s: %d\n", line[1], line[0], ret)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
