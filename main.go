/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	// "fmt"
	//
	// "github.com/mobiance/vimsy/utils"
	"github.com/mobiance/vimsy/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load environment variables from .env file
	cmd.Execute()
}

// func main() {
// 	err := utils.AddConfig("default"); if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Config added successfully! 🎉")
// }
