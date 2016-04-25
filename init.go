package main

import (
	"os/exec"
)

func Init() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"init")

	return cmd
}
