package main

import (
  "bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var userfile string
	fmt.Print("File name:")
	fmt.Scanln(&userfile)
	fileLineNumber(userfile)
}

func fileLineNumber(userfile string) {
	file := userfile + ".zfln"
	fin, err := os.Open(userfile)
	fout, errO := os.Create(file)
	defer fout.Close()
	defer fin.Close()
	if err != nil {
		fmt.Println(userfile, err)
	}
	if errO != nil {
		fmt.Println(file, errO)
	}
	buff := bufio.NewReader(fin)
	fout.WriteString("#########\t\t" + userfile + "\t\t#########\n\n")
	lineNum := 1
	for {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//fmt.Println(lineNum, "\t", line)
		fout.WriteString(strconv.Itoa(lineNum) + "\t" + line)
		lineNum += 1
	}
}
