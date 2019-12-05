package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./data/result.csv"
	isExist, err := PathExists(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isExist)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
