// +build linux,cgo

package nvml

import "unsafe"

// DeviceGetCPUAffinity retrieves an array of unsigned ints (sized to cpuSetSize) of bitmasks with the ideal CPU
// affinity for the device. For example, if processors 0, 1, 32, and 33 are ideal for the device and
// cpuSetSize == 2, result[0] = 0x3, result[1] = 0x3
func (a API) DeviceGetCPUAffinity(device Device, cpuSetSize uint32) (cpuSet uint32, err error) {
	err = a.call(a.nvmlDeviceGetCpuAffinity, uintptr(device), uintptr(cpuSetSize), uintptr(unsafe.Pointer(&cpuSet)))
	return
}

// DeviceSetCpuAffinity sets the ideal affinity for the calling thread and device using the guidelines given in
// DeviceGetCPUAffinity(). Note, this is a change as of version 8.0. Older versions set the affinity for a calling
// process and all children. Currently supports up to 64 processors.
func (a API) DeviceSetCpuAffinity(device Device) error {
	return a.call(a.nvmlDeviceSetCpuAffinity, uintptr(device))
}

// DeviceClearCpuAffinity clears all affinity bindings for the calling thread. Note, this is a change as of version 8.0
// as older versions cleared the affinity for a calling process and all children.
func (a API) DeviceClearCpuAffinity(device Device) (err error) {
	err = a.call(a.nvmlDeviceClearCpuAffinity, uintptr(device))
	return
}

// DeviceGetPersistenceMode eetrieves the persistence mode associated with this device.
// When driver persistence mode is enabled the driver software state is not torn down when the last client disconnects.
// By default this feature is disabled.
func (a API) DeviceGetPersistenceMode(device Device) (enabled bool, err error) {
	var state int32
	err = a.call(a.nvmlDeviceGetPersistenceMode, uintptr(device), uintptr(unsafe.Pointer(&state)))
	if err != nil {
		return
	}

	if state > 0 {
		enabled = true
	} else {
		enabled = false
	}

	return
}

// DeviceSetPersistenceMode sets the persistence mode for the device.
// For Linux only. Requires root/admin permissions.
// The persistence mode determines whether the GPU driver software is torn down after the last client exits.
// This operation takes effect immediately. It is not persistent across reboots.
// After each reboot the persistence mode is reset to "Disabled".
func (a API) DeviceSetPersistenceMode(device Device, mode bool) error {
	var modeInt int32 = 0
	if mode {
		modeInt = 1
	}

	return a.call(a.nvmlDeviceSetPersistenceMode, uintptr(device), uintptr(modeInt))
}
