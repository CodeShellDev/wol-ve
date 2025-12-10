package ve

import (
	"errors"
	"os/exec"
)

func StartVirsh(id string) error {
	if existsVirshVM(id) {
		startVirshVM(id)
	}

	if existsVirshLXC(id) {
		startVirshLXC(id)
	}

	return errors.New("Could not find vm/lxc with " + id)
}

func startVirshVM(id string) error {
	cmd := exec.Command("virsh", "start", id)

	return cmd.Run()
}

func startVirshLXC(id string) error {
	cmd := exec.Command("virsh", "-c", "lxc:///", "start", id)

	return cmd.Run()
}

func existsVirshVM(id string) bool {
	cmd := exec.Command("virsh", "dominfo", id)

	return cmd.Run() != nil
}

func existsVirshLXC(id string) bool {
	cmd := exec.Command("virsh", "-c", "lxc:///", "dominfo", id)

	return cmd.Run() != nil
}

func IsVirsh() bool {
	_, err := exec.LookPath("virsh")
	if err == nil {
		return true
	}

	return false
}