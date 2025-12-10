package ve

import (
	"errors"
	"os/exec"
)

func StartRawLXC(id string) error {
	if existsRawLXC(id) {
		return errors.New("Could not find lxc with " + id)
	}

	cmd := exec.Command("lxc-start", "-n", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func existsRawLXC(id string) bool {
	cmd := exec.Command("lxc-info", "-n", id)

	return cmd.Run() == nil
}

func IsRawLXC() bool {
	_, err := exec.LookPath("lxc-start")
	if err == nil {
		return true
	}

	return false
}