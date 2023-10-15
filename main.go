/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/s-alexsan/studybuddy/cmd"
	"github.com/s-alexsan/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
