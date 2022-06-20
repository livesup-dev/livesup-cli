/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/livesup-dev/livesup-cli/cmd"
	"github.com/livesup-dev/livesup-cli/internal/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
