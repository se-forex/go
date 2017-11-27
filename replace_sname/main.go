package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type In struct {
	Name string
	Val  string
}

func newLine(path string) []In {

	mStr := []In{}
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		mStr = append(mStr, In{string(scanner.Text()[:strings.IndexRune(scanner.Text(), '*')]), string(scanner.Text()[strings.IndexRune(scanner.Text(), '*')+2:])})
	}

	return mStr
}

func mapLine(path string) []In {

	mStr := []In{}
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		mStr = append(mStr, In{string(scanner.Text()[:strings.IndexRune(scanner.Text(), '*')]), string(scanner.Text()[strings.IndexRune(scanner.Text(), '*')+1:])})
	}

	return mStr
}

func main() {
	CurSt := newLine("./url.txt")
	NewSt := mapLine("./mapping.txt")

	for k := 0; k < len(CurSt); k++ {
		for m := 0; m < len(NewSt); m++ {
			if CurSt[k].Name == NewSt[m].Name {
				CurSt[k].Name = NewSt[m].Val
			}
		}
	}

	for i := 0; i < len(CurSt); i++ {
		fmt.Println(CurSt[i].Name + "*" + CurSt[i].Val)
	}

}
