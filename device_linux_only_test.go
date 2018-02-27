// +build linux,cgo

package nvml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeviceClearCpuAffinity(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceClearCpuAffinity(device)
	require.NoError(t, err)
}

func TestDeviceSetPersistenceMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceSetPersistenceMode(device, true)
	require.NoError(t, err)
}

func TestDeviceGetCPUAffinity(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetCPUAffinity(device, 1)
	require.NoError(t, err)
}

func TestDeviceGetPersistenceMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetPersistenceMode(device)
	require.NoError(t, err)
}

func TestDeviceSetCpuAffinity(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceSetCpuAffinity(device)
	require.NoError(t, err)
}
