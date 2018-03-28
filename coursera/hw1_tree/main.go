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

// // func shTree(path string, count int) error {
// // 	files, err := ioutil.ReadDir(path)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	//fmt.Println(len(files))

// // 	for _, file := range files {
// // 		if file.IsDir() {
// // 			fmt.Println("├───", file.Name())
// // 			err := dirTree(path + "/" + file.Name())
// // 			if err != nil {
// // 				continue
// // 			}
// // 		} else {
// // 			fmt.Println("│  └───", file.Name())
// // 		}
// // 	}
// // 	return nil
// // }

// // func printDir(sep string, level int) {
// // 	fmt.Println(multiStr("a", level))
// // }

// // func getDir(path string) (names []os.FileInfo, count int, err error) {
// // 	names, err = ioutil.ReadDir(path)
// // 	if err != nil {
// // 		return nil, 0, err
// // 	}
// // 	count = len(names)
// // 	return names, count, err
// // }

// func dirTree(path string) error {
// 	// sep1, sep2 := "├───", "└───"
// 	var count int

// 	// shTree(path, count)
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}

// 	//fmt.Println(len(files))

// 	for _, file := range files {
// 		if file.IsDir() {
// 			fmt.Println("d├───", file.Name())
// 		} else {
// 			fmt.Println("f├───", file.Name())
// 		}

// 		err := dirTree(path + "/" + file.Name())
// 		if err != nil {
// 			continue
// 		}
// 	}
// 	return nil
// }

/* =============== OLD CODE START ===================
var key int
var sep1, sep2 string = "├───", "└───"

func multiStr(s string, num int) string {
	var newStr string
	for i := 0; i < num; i++ {
		newStr = newStr + s
	}
	return newStr
}

func printTree(fname, sep string, count int) {
	if count != 0 {
		fmt.Println(multiStr("│    ", count)+sep, fname, count)
	} else {
		fmt.Println(sep, fname, count)
	}
}

func dirTree(path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	i := 0
	key++

	for _, file := range files {
		if file.IsDir() && printFiles {
			if i < (len(files) - 1) {
				printTree(file.Name(), sep1, key-1)
				i++
			} else {
				printTree(file.Name(), sep2, key-1)
			}
		}

		err := dirTree(path+"/"+file.Name(), printFiles)
		if err != nil {
			continue
		}
	}
	key--
	return nil
}
=============== OLD CODE STOP =================== */

func dirTree(path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			dirTree(path+"/"+files[i].Name(), printFiles)
		}
		if i == (len(files) - 1) {
			fmt.Println("l", files[i].Name())
		} else {
			fmt.Println("r", files[i].Name())
		}
	}

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
	//fmt.Println(multiStr("a", 3))
	fmt.Println("pfiles", printFiles)
	dirTree(path, printFiles)
}
