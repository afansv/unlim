package win

import (
	"golang.org/x/sys/windows/registry"
)

var (
	defaultTTL uint32 = 65

	keyPath = `SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`
	keyName = "DefaultTTL"
)

func SetTTL() error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetDWordValue(keyName, defaultTTL)
	if err != nil {
		return err
	}

	return nil
}
