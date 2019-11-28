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

var inputRegexp = regexp.MustCompile(`(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)`)

type lights struct {
	lights          [lightsX][lightsY]int
	totalLitCount   int
	totalBrightness int
}

// flickLights in inclusive coords with given cmd and keep track of lit count while we go
func (l *lights) flickLights(x1, y1, x2, y2 uint, cmd command) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if l.lights[x][y] == 1 && (cmd == off || cmd == toggle) {
				// light is on already - and we want to switch it off
				l.totalLitCount--
				l.lights[x][y] = 0
			} else if l.lights[x][y] == 0 && (cmd == on || cmd == toggle) {
				// light is off - and we want to switch it on
				l.totalLitCount++
				l.lights[x][y] = 1
			}
		}
	}
}

// twiddleLights in inclusive coords with given cmd and keep track of total brightness count while we go
func (l *lights) twiddleLights(x1, y1, x2, y2 uint, cmd command) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			var change int
			switch cmd {
			case off:
				if l.lights[x][y] > 0 {
					change = -1
				}
			case on:
				change = 1
			case toggle:
				change = 2
			}
			l.lights[x][y] += change
			l.totalBrightness += change
		}
	}
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

	part1Lights := lights{}
	part2Lights := lights{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		x1, y1, x2, y2, cmd := parseInput(txt)
		part1Lights.flickLights(x1, y1, x2, y2, cmd)
		part2Lights.twiddleLights(x1, y1, x2, y2, cmd)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part1: Lights lit: %d\n", part1Lights.totalLitCount)
	fmt.Printf("Part2: Total brightness: %d\n", part2Lights.totalBrightness)
}
