package ve

import (
	"errors"
	"os/exec"
)

func StartVirtualBoxVM(id string) error {
	if !existsVirtualBoxVM(id) {
		return errors.New("Could not find vm with " + id)
	}

	cmd := exec.Command("VBoxManage", "startvm", id, "--type headless")

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func existsVirtualBoxVM(id string) bool {
	cmd := exec.Command("VBoxManage", "showvminfo", id)

	return cmd.Run() != nil
}

func IsVirtualBox() bool {
	_, err := exec.LookPath("VBoxManage")
	if err == nil {
		return true
	}

	return false
}