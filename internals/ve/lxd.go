package ve

import (
	"errors"
	"os/exec"
)

func StartLXD(id string) error {
	if !existsLXD(id) {
		return errors.New("Could not find lxc with " + id)
	}

	cmd := exec.Command("lxc", "start", id)

	return cmd.Run()
}

func existsLXD(id string) bool {
	cmd := exec.Command("lxc", "info", id)

	return cmd.Run() == nil
}

func IsLXD() bool {
	_, err := exec.LookPath("lxc")
	if err == nil {
		return true
	}

	_, err = exec.LookPath("incus")
	if err == nil {
		return true
	}

	return false
}