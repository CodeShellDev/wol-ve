package ve

import "os/exec"

func StartRawLXC(id string) error {
	cmd := exec.Command("lxc-start", "-n", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func IsRawLXC() bool {
	_, err := exec.LookPath("lxc-start")
	if err == nil {
		return true
	}

	return false
}