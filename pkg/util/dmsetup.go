package util

import (
	"regexp"
)

const (
	dmsetupBinary = "dmsetup"
)

// DmsetupCreate creates a device mapper device with the given name and table
func DmsetupCreate(dmDeviceName, table string, executor Executor) error {
	opts := []string{
		"create", dmDeviceName, "--table", table,
	}
	_, err := executor.Execute(dmsetupBinary, opts)
	return err
}

// DmsetupSuspend suspends the device mapper device with the given name
func DmsetupSuspend(dmDeviceName string, executor Executor) error {
	opts := []string{
		"suspend", dmDeviceName,
	}
	_, err := executor.Execute(dmsetupBinary, opts)
	return err
}

// DmsetupResume removes the device mapper device with the given name
func DmsetupResume(dmDeviceName string, executor Executor) error {
	opts := []string{
		"resume", dmDeviceName,
	}
	_, err := executor.Execute(dmsetupBinary, opts)
	return err
}

// DmsetupReload reloads the table of the device mapper device with the given name and table
func DmsetupReload(dmDeviceName, table string, executor Executor) error {
	opts := []string{
		"reload", dmDeviceName, "--table", table,
	}
	_, err := executor.Execute(dmsetupBinary, opts)
	return err
}

// DmsetupRemove removes the device mapper device with the given name
func DmsetupRemove(dmDeviceName string, force, deferred bool, executor Executor) error {
	opts := []string{
		"remove", dmDeviceName,
	}
	if force {
		opts = append(opts, "--force")
	}
	if deferred {
		opts = append(opts, "--deferred")
	}
	_, err := executor.Execute(dmsetupBinary, opts)
	return err
}

// DmsetupDeps returns the dependent devices of the device mapper device with the given name
func DmsetupDeps(dmDeviceName string, executor Executor) ([]string, error) {
	opts := []string{
		"deps", dmDeviceName, "-o", "devname",
	}

	outputStr, err := executor.Execute(dmsetupBinary, opts)
	if err != nil {
		return nil, err
	}

	return parseDependentDevicesFromString(outputStr), nil
}

func parseDependentDevicesFromString(str string) []string {
	re := regexp.MustCompile(`\(([\w-]+)\)`)
	matches := re.FindAllStringSubmatch(str, -1)

	devices := make([]string, 0, len(matches))

	for _, match := range matches {
		devices = append(devices, match[1])
	}

	return devices
}