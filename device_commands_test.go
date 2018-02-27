package nvml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeviceClearECCErrorCounts(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceClearECCErrorCounts(device, VolatileECC)
	require.NoError(t, err)
}

func TestDeviceSetAPIRestriction(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceSetAPIRestriction(device, RestrictedAPISetAutoBoostedClocks, true)
	require.NoError(t, err)
}

func TestDeviceSetApplicationsClocks(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	mem, err := w.DeviceGetApplicationsClock(device, ClockMem)
	require.NoError(t, err)

	graphics, err := w.DeviceGetApplicationsClock(device, ClockGraphics)
	require.NoError(t, err)

	err = w.DeviceSetApplicationsClocks(device, mem, graphics)
	require.NoError(t, err)
}

func TestDeviceSetComputeMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	mode, err := w.DeviceGetComputeMode(device)
	require.NoError(t, err)

	err = w.DeviceSetComputeMode(device, mode)
	require.NoError(t, err)
}

func TestDeviceSetDriverModel(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, pending, err := w.DeviceGetDriverModel(device)
	require.NoError(t, err)

	err = w.DeviceSetDriverModel(device, pending, 0)
	require.NoError(t, err)
}

func TestDeviceSetECCMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceSetECCMode(device, true)
	require.NoError(t, err)
}

func TestDeviceSetGPUOperationMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, pending, err := w.DeviceGetGPUOperationMode(device)
	require.NoError(t, err)

	err = w.DeviceSetGPUOperationMode(device, pending)
	require.NoError(t, err)
}

func TestDeviceSetPowerManagementLimit(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	limit, err := w.DeviceGetPowerManagementLimit(device)
	require.NoError(t, err)

	err = w.DeviceSetPowerManagementLimit(device, limit)
	require.NoError(t, err)
}
