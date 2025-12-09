package ve

import "os/exec"

func StartLXD(id string) error {
	cmd := exec.Command("lxc", "start", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
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