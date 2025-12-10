package ve

import (
	"errors"
	"os/exec"
)

func StartPVE(id string) error {
	if existsPVEVM(id) {
		return startPVEVM(id)
	}

	if existsPVELXC(id) {
		return startPVELXC(id)
	}

	return errors.New("Could not find vm/lxc with " + id)
}

func existsPVEVM(id string) bool {
	cmd := exec.Command("qm", "status", id)

	return cmd.Run() != nil
}

func existsPVELXC(id string) bool {
	cmd := exec.Command("pct", "status", id)

	return cmd.Run() != nil
}

func startPVEVM(id string) error {
	cmd := exec.Command("qm", "start", id)

	return cmd.Run()
}

func startPVELXC(id string) error {
	cmd := exec.Command("pct", "start", id)

	return cmd.Run()
}

func IsPVE() bool {
	_, err := exec.LookPath("pveversion")
	if err == nil {
		return true
	}

	return false
}