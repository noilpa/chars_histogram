package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("missing path argument")
		os.Exit(1)
	}

	root := os.Args[1]
	info, err := os.Stat(root);
	if os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(2)
	}

	if !info.IsDir() {
		fmt.Println("path must contain dir")
		os.Exit(3)
	}

	fileList, err := getFileList(root)
	if err != nil {
		fmt.Printf("creating list of files is fail. err: %v\n", err)
		os.Exit(4)
	}
	time1 := time.Since(start)
	resultChan := processFileList(fileList)
	resultFreqMap := make(map[string]int)

	time2 := time.Since(start)
	var sum int
	for fm := range resultChan {
		for k, v := range fm {
			sum += v
			if _, ok := resultFreqMap[k]; ok {
				resultFreqMap[k] += v
				continue
			}
			resultFreqMap[k] = v
		}
	}

	prettyRes, err := json.MarshalIndent(resultFreqMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	fmt.Println(string(prettyRes))
	fmt.Println("Sum:", sum)
	fmt.Println("Get file list time:", time1)
	fmt.Println("Process file list time:", time2)
	fmt.Println("All time:", time.Since(start))
}

func getFileList(root string) ([]string, error) {
	fileList := make([]string, 0)
	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				fileList = append(fileList, path)
			}

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func processFileList(fileList []string) chan map[string]int {
	resultChan := make(chan map[string]int, len(fileList))
	wg := new(sync.WaitGroup)
	for _, file := range fileList {
		wg.Add(1)
		go do(file, resultChan, wg)
	}
	wg.Wait()
	close(resultChan)

	return resultChan
}

func processFile(path string) (map[string]int, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return calculateCharsFrequency(file), nil
}

func calculateCharsFrequency(file []byte) map[string]int {
	charsFrequency := make(map[string]int, 0)
	for _, bChar := range file {
		char := string(bChar)
		if _, ok := charsFrequency[char]; !ok {
			charsFrequency[char] = 0
		}
		charsFrequency[char]++
	}
	return charsFrequency
}

func do(path string, resultChan chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	freqMap, err :=  processFile(path)
	if err != nil {
		// if get corrupted file - skip it
		fmt.Printf("processing %v fail. err: %v\n", path, err)
		return
	}
	resultChan <- freqMap
}
