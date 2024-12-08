package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func numTimes(nums []int, num int) int {
	numTimes := 0
	for _, i := range nums {
		if i == num {
			numTimes += 1
		}
	}
	return numTimes
}

func part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	var sum int
	for i := range left {
		sum = sum + absDiffInt(right[i], left[i])
	}
	return sum
}

func part2(left, right []int) int {
	score := 0
	for _, i := range left {
		score = score + i*numTimes(right, i)
	}
	return score
}

func main() {
	// open file
	file, err := os.Open("./day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		num0, _ := strconv.Atoi(parts[0])
		num1, _ := strconv.Atoi(parts[1])

		left = append(left, num0)
		right = append(right, num1)
	}

	// left := []int{3, 4, 2, 1, 3, 3}
	// right := []int{4, 3, 5, 3, 9, 3}

	//fmt.Printf("sum: %d", part1(left, right))
	fmt.Printf("score: %d", part2(left, right))
}
