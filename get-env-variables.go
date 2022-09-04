package main

import (
	"fmt"
	"os"
)

// ! Remove the funtion name to main before run
func main1() {
	//Get all environment variables
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	//Get spesific environment variables
	userName := os.Getenv("USER")
	homePath := os.Getenv("HOME")

	fmt.Println("User Name: " + userName)
	fmt.Println("Home Path: " + homePath)
}
