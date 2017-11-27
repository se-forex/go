package main

import (
	"bufio"
	"os"
	"strings"
)

type In struct {
	Name string
}

func mapLine(path string) []In {

	mStr := []In{}
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		mStr = append(mStr, In{string(scanner.Text()[:strings.IndexRune(scanner.Text(), '_')])})
	}

	return mStr
}

func Unify(target []In) {
	// inFile, _ := os.Open("./url.txt")
	// defer inFile.Close()
	// scanner := bufio.NewScanner(inFile)
	// scanner.Split(bufio.ScanLines)

	var targetArray []string

	// for scanner.Scan() {
	// 	tmp_orig := scanner.Text()
	// 	for i := 0; i < len(target); i++ {
	// 		tmp_new := strings.Replace(string(scanner.Text()), target[i].Val, target[i].Name, 1)
	// 		if strings.Compare(tmp_new, tmp_orig) != 0 {
	// 			targetArray = append(targetArray, tmp_new)
	// 		}
	// 		// if strings.Compare(tmp_new, tmp_orig) == 0 {
	// 		// 	targetArray = append(targetArray, tmp_orig)
	// 		// }
	// 	}
	// }

	// for k := 0; k < len(targetArray); k++ {
	// 	fmt.Println(targetArray[k])
	// }
	// fmt.Println(len(targetArray))

	key := 0

	for i := 0; i < (len(target) - 1); i++ {

		if key != 1 {
			targetArray = append(targetArray, target[i].Name)
		}

		if target[i].Name == target[i+1].Name {
			key = 1
		} else {
			key = 0
		}

		// i = key
	}

	for i := 0; i < len(targetArray); i++ {
		println(targetArray[i])
	}

}

func main() {
	NewSt := mapLine("./url.txt")

	// len_NewSt := 0
	// for k := 0; k < len(NewSt); k++ {
	// 	println(NewSt[k].Name)
	// }

	Unify(NewSt)
	// println(Index(NewSt, ""))

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
