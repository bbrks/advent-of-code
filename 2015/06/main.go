package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	lightsX = 1000
	lightsY = 1000
)

var (
	lights [lightsX][lightsY]bool
	lit    int // keep track of how many lights are on

	inputRegexp = regexp.MustCompile(`(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)`)
)

// flickLights in inclusive coords with given cmd and keep track of lit count while we go
func flickLights(x1, y1, x2, y2 uint, cmd command) int {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if lights[x][y] && (cmd == off || cmd == toggle) {
				// light is on already - and we want to switch it off
				lit--
				lights[x][y] = false
			} else if !lights[x][y] && (cmd == on || cmd == toggle) {
				// light is off - and we want to switch it on
				lit++
				lights[x][y] = true
			}
		}
	}
	return lit
}

func parseInput(s string) (x1, y1, x2, y2 uint, cmd command) {
	result := inputRegexp.FindAllStringSubmatch(s, 1)
	return strToUint(result[0][2]), // x1
		strToUint(result[0][3]), // y1
		strToUint(result[0][4]), // x2
		strToUint(result[0][5]), // y2
		strToCommand(result[0][1]) // cmd
}

func strToUint(s string) uint {
	resInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint(resInt)
}

type command uint

const (
	off command = iota
	on
	toggle
)

func strToCommand(s string) command {
	switch s {
	case "turn off":
		return off
	case "turn on":
		return on
	case "toggle":
		return toggle
	}
	panic("unknown command: " + s)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lastLitCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		startX, startY, endX, endY, cmd := parseInput(txt)
		lastLitCount = flickLights(startX, startY, endX, endY, cmd)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lights lit: %d\n", lastLitCount)
}
