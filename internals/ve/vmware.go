package ve

import (
	"errors"
	"os/exec"
)

func StartVMWare(id string) error {
	if existsVMWareVM(id) {
		cmd := exec.Command("vim-cmd", "vmsvc/power.on", id)

		return cmd.Run()
	}

	if existsVMWareFusionVM(id) {
		cmd := exec.Command("vmrun", "start", id)

		return cmd.Run()
	}

	return errors.New("Could not find vm with " + id)
}

func existsVMWareVM(id string) bool {
	cmd := exec.Command("vim-cmd", "vmsvc/getallvms", "|", "grep", "-w", id)

	return cmd.Run() != nil
}

func existsVMWareFusionVM(id string) bool {
	cmd := exec.Command("test", "-f", id)

	return cmd.Run() != nil
}

func IsVMWare() bool {
	_, err := exec.LookPath("vim-cmd")
	if err == nil {
		return true
	}

	_, err = exec.LookPath("vmrun")
	if err == nil {
		return true
	}

	return false
}