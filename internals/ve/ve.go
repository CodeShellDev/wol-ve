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
		return StartVirtualBoxVM(id)
	}

	if IsLXD() {
		return StartLXD(id)
	}

	if IsRawLXC() {
		return StartRawLXC(id)
	}

	return errors.New("No virtual environment detected")
}