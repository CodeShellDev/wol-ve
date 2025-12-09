package ve

import "os/exec"

func StartVMWare(id string) error {
	cmd := exec.Command("vim-cmd", "vmsvc/power.on", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	cmd = exec.Command("vmrun", "start", id)

	err = cmd.Run()

	if err == nil {
		return nil
	}

	return err
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