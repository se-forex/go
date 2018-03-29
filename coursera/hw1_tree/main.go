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

var deep, ddeep int
var regSep, lastSep string = "├───", "└───"
var gSep []int

type TreeUnit struct {
	pos  []int
	last bool
	name string
}

type Tree struct {
	unit []TreeUnit
}

func multiStr(s string, num int) string {
	var newStr string
	for i := 0; i < num; i++ {
		newStr = newStr + s
	}
	return newStr
}

func deepCalc(path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			if ddeep < deep {
				ddeep = deep
			}
			deep++
			deepCalc(path+"/"+files[i].Name(), printFiles)
			deep--
		}
	}
	return nil
}

// func myTempFunc() {
// 	var sep string
// 	gSep = append(gSep, deep)

// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}

// 	for i := 0; i < len(files); i++ {
// 		if i == (len(files) - 1) {
// 			sep = lastSep
// 			gSep[deep] = 0
// 		} else {
// 			sep = regSep
// 			gSep[deep] = 1
// 		}

// 		if files[i].IsDir() {
// 			fmt.Print(multiStr("│   ", deep), sep, files[i].Name(), deep, " ", gSep[deep], "\n")
// 			//fmt.Print(gSep)
// 			deep++
// 			dirTree(path+"/"+files[i].Name(), printFiles)
// 			deep--
// 		}

// 		if !files[i].IsDir() && printFiles {
// 			fmt.Print(multiStr("│   ", deep), sep, files[i].Name(), deep, " ", gSep[deep], "\n")
// 			//fmt.Print(gSep)
// 		}
// 	}

// 	// fmt.Print(gSep)
// }

func fillTree(tree Tree, path string, printFiles bool) {
	fmt.Print(tree)
}

func dirTree(path string, printFiles bool) error {
	deepCalc(path, printFiles)

	var treeUnit TreeUnit = TreeUnit{
		pos: [1,2,3],
		
	}

	var tree Tree = Tree{}
	tree.unit = append(tree.unit,)

	fillTree(tree, path, printFiles)

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
	deepCalc(path, printFiles)
	deep = 0
	fmt.Println(ddeep)
	dirTree(path, printFiles)
}
