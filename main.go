package main

import (
	//"fmt"
	//"io"
	"os"
	//"path/filepath"
	//"strings"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func dirTree(out *os.File, path string, printFiles bool) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("destination should be a directory")
	}

	listDir(path, 0, printFiles, false)
	return nil
}

func listDir(path string, level int, printFiles, islast bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	var tabs, tabsl string
	tabs = strings.Repeat("|\t", level)
	if level > 0 {
		tabsl = strings.Repeat("|\t", level-1)
		tabsl += "\t"
	}
	level += 1
	for i, f := range files {
		if islast && level > 0 {
			fmt.Print(tabsl)
		} else {
			fmt.Print(tabs)
		}

		fmt.Print(seg(i == len(files)-1))

		if f.IsDir() {
			islast = i == len(files)-1
			fmt.Printf("───%s\n", f.Name())
			listDir(path+string(os.PathSeparator)+f.Name(), level, printFiles, islast)
		} else if printFiles {
			fmt.Printf("───%s\n", f.Name())
		}
	}
	return nil
}

func seg(last bool) string {
	if last {
		return "└"
	} else {
		return "├"
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
