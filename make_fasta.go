package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
//open file for reading
func openFile(fname string) *os.File{
	var file,err = os.Open(fname)
	if err != nil {
		panic(err)
	}
	return file
}
//create file in the directory
func createFile(fname string){
	var file, err = os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

//open file for write
func writeFile(ofname string) *os.File {
	var ofile, err = os.OpenFile(ofname, os.O_RDWR, 0644)
	if  err != nil {
		panic(err)
	}
	return ofile
}

//main
func main() {
	start := time.Now()
	var ofname = "fullsequence.fa" //write a fasta file
	var infile = openFile("cafaUniprotBPOIfull.csv") //input file of full sequences
	createFile(ofname) //create output file
	var ofile = writeFile(ofname) //open file for writing
	fscanner := bufio.NewScanner(infile) //fscanner reads file line by line
	for fscanner.Scan() {
		var tokens = strings.Split(fscanner.Text(),",")//tokenize
		ofile.WriteString(">"+tokens[0]+"\n")
		ofile.WriteString(tokens[len(tokens)-1]+"\n")
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Print(elapsed)
}
