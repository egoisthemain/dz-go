package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func average(array []int) float64 {
	summ := 0
	n := 0
	if len(array) == 0 {
		return 0
	}
	for _, value := range array {
		summ += value
		n += 1
	}
	return float64(summ) / float64(n)
}

func summ(array []int) int {
	summ := 0
	for _, value := range array {
		summ += value
	}
	return summ
}

func median(array []int) float64 {
	sort.Ints(array)
	l := len(array)
	if l%2 == 0 {
		return float64(array[int(l/2.0)-1]+array[int(l/2.0)]) / 2
	} else {
		return float64(array[int((l-1)/2)])
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var operation string
	for {
		fmt.Print("Введите операцию(AVG/SUM/MED): ")
		opLine, _ := reader.ReadString('\n')
		operation = strings.TrimSpace(opLine)
		if strings.EqualFold(operation, "AVG") || strings.EqualFold(operation, "SUM") || strings.EqualFold(operation, "MED") {
			break
		}
		fmt.Println("Неверный формат операции")
	}

	fmt.Print("Введите числа через запятую: ")
	line, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line)
	parts := strings.Split(s, ",")
	sliceOfNums := make([]int, 0, len(parts))
	for i := range parts {
		sTemp := strings.TrimSpace(parts[i])
		numTemp, err := strconv.Atoi(sTemp)
		if err != nil {
			fmt.Println(errors.New("convertation error"))
			os.Exit(1)
		}
		sliceOfNums = append(sliceOfNums, numTemp)
	}
	fmt.Println(sliceOfNums)

	switch {
	case strings.EqualFold(operation, "AVG"):
		fmt.Printf("Average: %.2f", average(sliceOfNums))
	case strings.EqualFold(operation, "SUM"):
		fmt.Printf("Summ: %d", summ(sliceOfNums))
	case strings.EqualFold(operation, "MED"):
		fmt.Printf("Median: %.2f", median(sliceOfNums))
	}
}
