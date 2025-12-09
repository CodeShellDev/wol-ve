package ve

import "os/exec"

func StartVirtualBox(id string) error {
	cmd := exec.Command("VBoxManage", "startvm", id, "--type headless")

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func IsVirtualBox() bool {
	_, err := exec.LookPath("VBoxManage")
	if err == nil {
		return true
	}

	return false
}