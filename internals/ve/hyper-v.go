package ve

import "os/exec"

func StartHyperV(id string) error {
	cmd := exec.Command("powershell.exe", "-Command", "Start-VM", "-Name", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func IsHyperV() bool {
	cmd := exec.Command("powershell.exe", "-Command", "if (Get-Command Get-VM -ErrorAction SilentlyContinue) { exit 0 } else { exit 1 }")

	err := cmd.Run()
	if err == nil {
		return true
	}

	return false
}