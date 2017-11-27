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

// func newLine(path string) []In {

// 	mStr := []In{}
// 	inFile, _ := os.Open(path)
// 	defer inFile.Close()
// 	scanner := bufio.NewScanner(inFile)
// 	scanner.Split(bufio.ScanLines)

// 	for scanner.Scan() {
// 		mStr = append(mStr, In{string(scanner.Text()[:strings.IndexRune(scanner.Text(), '*')]), string(scanner.Text()[strings.IndexRune(scanner.Text(), '*')+2:])})
// 	}

// 	return mStr
// }

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

func ReplaceFile(target []In) {
	inFile, _ := os.Open("./url.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var targetArray []string

	for scanner.Scan() {
		tmp_orig := scanner.Text()
		for i := 0; i < len(target); i++ {
			tmp_new := strings.Replace(string(scanner.Text()), target[i].Val, target[i].Name, 1)
			if strings.Compare(tmp_new, tmp_orig) != 0 {
				targetArray = append(targetArray, tmp_new)
			}
			// if strings.Compare(tmp_new, tmp_orig) == 0 {
			// 	targetArray = append(targetArray, tmp_orig)
			// }
		}
	}

	for k := 0; k < len(targetArray); k++ {
		fmt.Println(targetArray[k])
	}
	fmt.Println(len(targetArray))
}

func main() {
	// CurSt := newLine("./url.txt")
	NewSt := mapLine("./mapping.txt")

	// len_NewSt := 0
	// for k := 0; k < len(NewSt); k++ {
	// 	len_NewSt++
	// }

	ReplaceFile(NewSt)
	// println(len_NewSt)

	// for k := 0; k < len(CurSt); k++ {
	// 	for m := 0; m < len(NewSt); m++ {
	// 		if CurSt[k].Name == NewSt[m].Name {
	// 			CurSt[k].Name = NewSt[m].Val
	// 		}
	// 	}
	// }

	// for i := 0; i < len(CurSt); i++ {
	// 	fmt.Println(CurSt[i].Name + "*" + CurSt[i].Val)
	// }

}
