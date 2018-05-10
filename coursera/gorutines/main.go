package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	itNums = 7
	goNums = 5
)

func doSomeWork(in int) {
	for i := 0; i < itNums; i++ {
		fmt.Println(formatWork(in, i))
		runtime.Gosched()
	}
}

func formatWork(in, i int) string {
	return fmt.Sprintln(
		strings.Repeat(" ", in),
		"*",
		strings.Repeat(" ", goNums-in),
		"th", in,
		"iter", i, strings.Repeat("*", i))
}

func main() {
	for i := 0; i < goNums; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}
