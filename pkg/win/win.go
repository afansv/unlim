package win

import (
	"golang.org/x/sys/windows/registry"
)

var (
	ipv4ParametersPath = `SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`
	defaultTTLKey      = "DefaultTTL"

	defaultTTL uint32 = 65

	ipv6ParametersPath    = `SYSTEM\CurrentControlSet\services\TCPIP6\Parameters`
	disabledComponentsKey = "DisabledComponents"

	disabledComponents uint32 = 255 // IPv6 disabled
)

func SetIPv4TTL() error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, ipv4ParametersPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetDWordValue(defaultTTLKey, defaultTTL)
	if err != nil {
		return err
	}

	return nil
}

func SetIPv6TTL() error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, ipv6ParametersPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetDWordValue(defaultTTLKey, defaultTTL)
	if err != nil {
		return err
	}

	return nil
}

func DisableIPv6() error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, ipv6ParametersPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetDWordValue(disabledComponentsKey, disabledComponents)
	if err != nil {
		return err
	}

	return nil
}
