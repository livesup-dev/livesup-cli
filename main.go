/*
Copyright © 2022 Emiliano Jankowski

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
