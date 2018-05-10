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

var gDeep, deep int

type DTree struct {
	el []Dirs
}

type Dirs struct {
	Name   string
	IsDir  bool
	Size   int64
	IsLast bool
	Deep   int
	Pos    []int
	Sep    []int
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
			deep++
			if gDeep < deep {
				gDeep = deep
			}
			deepCalc(path+"/"+files[i].Name(), printFiles)
			deep--
		}
	}
	return nil
}

func dirsCalc(files []os.FileInfo) int {
	num := 0
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			num = i
		}
	}
	return num
}

func fillStruct(tree DTree, path string, printFiles bool) (DTree, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return tree, err
	}

	if printFiles {
		for i := 0; i < len(files); i++ {

			var unit = Dirs{
				Pos: make([]int, gDeep, gDeep),
				Sep: make([]int, gDeep, gDeep),
			}

			if files[i].IsDir() {
				unit.Name = files[i].Name()
				unit.IsDir = true
				unit.Deep = deep
				if deep > 0 {
					// if tree.el[(len(tree.el)-1)].IsDir {

					// }
					unit.Pos[deep-1] = 1
				}
				if i == (len(files) - 1) {
					unit.IsLast = true
				}
				tree.el = append(tree.el, unit)
				deep++
				tree, _ = fillStruct(tree, path+"/"+unit.Name, printFiles)
				deep--
			} else {
				unit.Name = files[i].Name()
				unit.Size = files[i].Size()
				unit.Deep = deep
				if deep > 0 {
					unit.Pos[deep-1] = 1
				}
				if i == (len(files) - 1) {
					unit.IsLast = true
				}
				tree.el = append(tree.el, unit)
			}
		}
	} else {

		for i := 0; i < len(files); i++ {

			var unit = Dirs{
				Pos: make([]int, gDeep, gDeep),
				Sep: make([]int, gDeep, gDeep),
			}

			if files[i].IsDir() {
				numDirs := dirsCalc(files)
				unit.Name = files[i].Name()
				unit.IsDir = true
				unit.Deep = deep
				if deep > 0 {
					unit.Pos[deep-1] = 1
				}
				if i == numDirs {
					unit.IsLast = true
				}
				tree.el = append(tree.el, unit)
				deep++
				tree, _ = fillStruct(tree, path+"/"+unit.Name, printFiles)
				deep--
			}
		}
	}
	return tree, nil
}

func setLevels(tree DTree) DTree {
	for i := range tree.el {
		for k := range tree.el[i].Pos {
			if k > 0 && i > 0 && tree.el[i].Pos[k] == 1 {
				tree.el[i].Pos[k-1] = tree.el[i-1].Pos[k-1]
				if k > 1 {
					for j := k; j > 0; j-- {
						tree.el[i].Pos[j-1] = tree.el[i-1].Pos[j-1]
					}
				}
			}
		}
	}
	return tree
}

func setSep(tree DTree) DTree {
	for i := 0; i < len(tree.el); i++ {
		if (tree.el[i].Deep == 0) && (tree.el[i].IsDir) {
			// fmt.Println("-", tree.el[i].Name)
			key := i
			if !tree.el[i].IsLast {
				for tree.el[key+1].Pos[tree.el[i].Deep] != tree.el[i].Pos[tree.el[i].Deep] {
					tree.el[key+1].Sep[tree.el[i].Deep] = tree.el[key+1].Pos[tree.el[i].Deep]
					// fmt.Println(tree.el[key+1].Name, tree.el[key+1].Pos, tree.el[key+1].Sep)
					key++
				}
			}
		} else if (tree.el[i].Deep != 0) && (tree.el[i].IsDir) {
			// fmt.Println("-", tree.el[i].Name)
			key := i
			if !tree.el[i].IsLast { //deep = 1
				for tree.el[key+1].Pos[tree.el[i].Deep] != tree.el[i].Pos[tree.el[i].Deep] {
					tree.el[key+1].Sep[tree.el[i].Deep] = tree.el[key+1].Pos[tree.el[i].Deep]
					// fmt.Print(tree.el[key+1].Sep[tree.el[i].Deep], "+", tree.el[key+1].Pos[tree.el[i].Deep], "+", tree.el[i].Deep)
					// fmt.Println(tree.el[key+1].Name, tree.el[key+1].Pos, tree.el[key+1].Sep)
					key++
				}
			}
		}
	}
	return tree
}

func dirTree(path string, printFiles bool) error {

	deepCalc(path, printFiles)

	var err error
	var tree = DTree{}

	tree, err = fillStruct(tree, path, printFiles)
	if err != nil {
		fmt.Println("ERROR")
	}

	tree = setLevels(tree)
	tree = setSep(tree)

	for i := range tree.el {
		for num := range tree.el[i].Pos {
			// if i > 0 && (tree.el[i-1].IsLast) && (num == (len(tree.el[i].Pos)-1)) {
			//    fmt.Print(multiStr("    ", tree.el[i].Pos[num]))
			// } else {
			fmt.Print(multiStr("│", tree.el[i].Sep[num]))
			if tree.el[i].Sep[num] != 0 {
				fmt.Print(multiStr("   ", tree.el[i].Pos[num]))
			} else {
				fmt.Print(multiStr("    ", tree.el[i].Pos[num]))
			}
			// }
		}
		if tree.el[i].IsLast {
			if tree.el[i].IsDir {
				fmt.Print("└───", tree.el[i].Name, "\n")
			} else if tree.el[i].Size != 0 {
				fmt.Print("└───", tree.el[i].Name, " ( ", tree.el[i].Size, "b ) ", "\n")
			} else {
				fmt.Print("└───", tree.el[i].Name, " ( empty ) ", "\n")
			}
		} else {
			if tree.el[i].IsDir {
				fmt.Print("├───", tree.el[i].Name, "\n")
			} else if tree.el[i].Size != 0 {
				fmt.Print("├───", tree.el[i].Name, " ( ", tree.el[i].Size, "b ) ", "\n")
			} else {
				fmt.Print("├───", tree.el[i].Name, " ( empty ) ", "\n")
			}
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	// 	panic(err.Error())
	// }

	dirTree(path, printFiles)

}
