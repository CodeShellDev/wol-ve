package ve

import "os/exec"

func StartPVE(id string) error {
	cmd := exec.Command("qm", "start", id)

	err := cmd.Run()

	if err == nil {
		return nil
	}

	cmd = exec.Command("pct", "start", id)

	err = cmd.Run()

	if err == nil {
		return nil
	}

	return err
}

func IsPVE() bool {
	_, err := exec.LookPath("pveversion")
	if err == nil {
		return true
	}

	return false
}