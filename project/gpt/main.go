/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "learner/project/gpt/router"

//go:generate swag init -o ./docs
func main() {
	router.Server()
}
