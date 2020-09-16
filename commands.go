package idm

import (
	"os/exec"
)

func runCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}
func mainQueueStart() error {
	return runCommand(idmPath, []string{"/s"})
}

func addToQueue(url string) error {
	return runCommand(idmPath, []string{"/d", url, "/a"})
}

func startDownload(args []string) error {
	return runCommand(idmPath, args)
}
