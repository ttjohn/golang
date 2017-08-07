/*
*  This program can reads the input from file/buffer and 
*  caluclate wiegthed computation time for various jobs.
*  you can add a new computation function and send it to initlizer.
*  schedulur filename <index type> 
*  "index type" is type for calculating the weight   
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type diff struct {
  wieght, ptime int
  wtIndex float32
 }

type  ByDiffWeight []diff
func (a ByDiffWeight) Len() int           { return len(a) }
func (a ByDiffWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDiffWeight) Less(i, j int) bool { 
	if a[i].wtIndex > a[j].wtIndex {
       return true
    }
    if a[i].wtIndex < a[j].wtIndex {
       return false
    }
    return a[i].wieght > a[j].wieght
}

type indexCalculater func(wieght, ptime int) (float32) 

func diffIndex(wieght, ptime int) (float32) {
  return   float32(wieght - ptime)
}

func ratioIndex(wieght, ptime int) (float32) {
  return  float32(wieght) / float32(ptime)
}

func initializeData(reader *bufio.Scanner,indexFunc indexCalculater)([] diff) {

	reader.Scan()
	textCount := reader.Text()
	count, e := strconv.Atoi(textCount);
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	
	line := 0
	var jobs = make([]diff, count) 

	for ; line < count; {
		reader.Scan()
		textData := reader.Text()
		words := strings.Fields(textData)
		jobs[line].wieght, e = strconv.Atoi(words[0]);
		jobs[line].ptime, e = strconv.Atoi(words[1]);
		jobs[line].wtIndex= indexFunc(jobs[line].wieght,jobs[line].ptime);
		if e != nil {
			fmt.Println(e)
			break
		}
		line++
	}
	return jobs
}

func findWieghtedCompletionTime(jobs []diff) int64{

	compTime := 0
	sum := int64(0);
	
	for _, job := range jobs {
		compTime += job.ptime;
		sum += int64((job.wieght * compTime))
	}
	return sum
}


func main() {

	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	
	file := os.Args[1]
	
	indexType := "default"
	if len(os.Args) > 2 {
		indexType = os.Args[2]
	}
	fmt.Println("Index Type : ", indexType)
	var indexFunc indexCalculater
	// Defult to diffrence based Index
	indexFunc = diffIndex

	if indexType == "ratio" {
		indexFunc = ratioIndex
	}
	
	f, err := os.Open(pwd + "\\"+ file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	
	jobs := initializeData(reader,indexFunc)
	sort.Sort(ByDiffWeight(jobs))
		
	fmt.Println("Weighted Sum : ", findWieghtedCompletionTime(jobs))
}
