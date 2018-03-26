package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	// "io"
	"os"
	// "path/filepath"
	// "strings"
)

func multiStr(s string, num int) string {
	var newStr string
	for i := 0; i < num; i++ {
		newStr = newStr + s
	}
	return newStr
}

// func shTree(path string, count int) error {
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}

// 	//fmt.Println(len(files))

// 	for _, file := range files {
// 		if file.IsDir() {
// 			fmt.Println("├───", file.Name())
// 			err := dirTree(path + "/" + file.Name())
// 			if err != nil {
// 				continue
// 			}
// 		} else {
// 			fmt.Println("│  └───", file.Name())
// 		}
// 	}
// 	return nil
// }

func printDir(sep string, level int) {
	fmt.Println(multiStr("a", level))
}

func getDir(path string) (names []os.FileInfo, count int, err error) {
	names, err = ioutil.ReadDir(path)
	if err != nil {
		return nil, 0, err
	}
	count = len(names)
	return names, count, err
}

func dirTree(path string) error {
	var level int
	//sep1, sep2 := "├───", "───"
	names, count, err := getDir(path)

	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		name := names[i]
		if name.IsDir() {
			fmt.Println(name)
			level++
			dirTree(path + "/" + name.Name())
		}
	}
	level = 0

	// sep1, sep2 := "├───", "└───"
	// var count int

	// shTree(path, count)
	// files, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	return err
	// }

	// //fmt.Println(len(files))

	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Println("d├───", file.Name())
	// 	} else {
	// 		fmt.Println("f├───", file.Name())
	// 	}

	// 	err := dirTree(path + "/" + file.Name())
	// 	if err != nil {
	// 		continue
	// 	}
	// }
	return nil
}

func main() {
	// out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	// 	panic(err.Error())
	// }
	fmt.Println(multiStr("a", 3))
	fmt.Println("pfiles", printFiles)
	dirTree(path)
}
