package ve

import "errors"

func StartVirtualHost(id string) error {
	if IsPVE() {
		return StartPVE(id)
	}

	if IsVirsh() {
		return StartVirsh(id)
	}

	if IsVMWare() {
		return StartVMWare(id)
	}

	if IsVirtualBox() {
		return StartVirtualBox(id)
	}

	if IsLXD() {
		return StartLXD(id)
	}

	if IsRawLXC() {
		return StartRawLXC(id)
	}

	if IsHyperV() {
		return StartHyperV(id)
	}

	return errors.New("No virtual environment detected")
}