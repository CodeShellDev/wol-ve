package ve

import "os/exec"

func StartVirsh(id string) error {
	cmd := exec.Command("virsh", "start", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	cmd = exec.Command("virsh", "-c", "lxc:///", "start", id)

	err = cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func IsVirsh() bool {
	_, err := exec.LookPath("virsh")
	if err == nil {
		return true
	}

	return false
}