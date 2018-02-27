package nvml

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeviceGetAPIRestriction(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetAPIRestriction(device, RestrictedAPISetApplicationClocks)
	require.NoError(t, err)
}

func TestDeviceGetApplicationsClock(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clock, err := w.DeviceGetApplicationsClock(device, ClockGraphics)
	require.NoError(t, err)
	require.NotZero(t, clock)
}

func TestDeviceGetAutoBoostedClocksEnabled(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetAutoBoostedClocksEnabled(device)
	require.NoError(t, err)
}

func TestDeviceGetBAR1MemoryInfo(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	mem, err := w.DeviceGetBAR1MemoryInfo(device)
	require.NoError(t, err)
	require.NotZero(t, mem.Used)
	require.NotZero(t, mem.Total)
	require.NotZero(t, mem.Free)
}

func TestDeviceGetBoardID(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	id, err := w.DeviceGetBoardID(device)
	require.NoError(t, err)
	require.NotEqual(t, 0, id)
}

func TestDeviceGetBoardPartNumber(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	number, err := w.DeviceGetBoardPartNumber(device)
	require.NoError(t, err)
	require.NotEmpty(t, number)
}

func TestDeviceGetBrand(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	brand, err := w.DeviceGetBrand(device)
	require.NoError(t, err)
	require.NotEqual(t, BrandUnknown, brand)
}

func TestDeviceGetBridgeChipInfo(t *testing.T) {
	t.Skip()
}

func TestDeviceGetClock(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clock, err := w.DeviceGetClock(device, ClockMem, ClockIDCurrent)
	require.NoError(t, err)
	require.NotZero(t, clock)
}

func TestDeviceGetClockInfo(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	// Core Clock
	graphics, err := w.DeviceGetClockInfo(device, ClockGraphics)
	require.NoError(t, err)
	require.True(t, graphics > 0)

	// Memory Clock
	mem, err := w.DeviceGetClockInfo(device, ClockMem)
	require.NoError(t, err)
	require.True(t, mem > 0)
}

func TestDeviceGetComputeMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetComputeMode(device)
	require.NoError(t, err)
}

func TestDeviceGetComputeRunningProcesses(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	list, err := w.DeviceGetComputeRunningProcesses(device)
	require.NoError(t, err)
	require.NotEmpty(t, list)
}

func TestDeviceGetCount(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	count, err := w.DeviceGetCount()
	require.NoError(t, err)
	require.True(t, count > 0)
}

func TestDeviceGetCudaComputeCapability(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetCudaComputeCapability(device)
	require.NoError(t, err)
}

func TestDeviceGetCurrPcieLinkGeneration(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	gen, err := w.DeviceGetCurrPcieLinkGeneration(device)
	require.NoError(t, err)
	require.NotZero(t, gen)
}

func TestDeviceGetCurrPcieLinkWidth(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	gen, err := w.DeviceGetCurrPcieLinkWidth(device)
	require.NoError(t, err)
	require.NotZero(t, gen)
}

func TestDeviceGetCurrentClocksThrottleReasons(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	reasons, err := w.DeviceGetCurrentClocksThrottleReasons(device)
	require.NoError(t, err)
	require.NotZero(t, reasons)
}

func TestDeviceGetDecoderUtilization(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetDecoderUtilization(device)
	require.NoError(t, err)
}

func TestDeviceGetDefaultApplicationsClock(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clock, err := w.DeviceGetDefaultApplicationsClock(device, ClockMem)
	require.NoError(t, err)
	require.NotZero(t, clock)
}

func TestDeviceGetDetailedECCErrors(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetDetailedECCErrors(device, MemoryErrorTypeCorrected, VolatileECC)
	require.NoError(t, err)
}

func TestDeviceGetDisplayActive(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	connected, err := w.DeviceGetDisplayActive(device)
	require.NoError(t, err)
	require.True(t, connected)
}

func TestDeviceGetDisplayMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	connected, err := w.DeviceGetDisplayMode(device)
	require.NoError(t, err)
	require.True(t, connected)
}

func TestDeviceGetDriverModel(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetDriverModel(device)
	require.NoError(t, err)
}

func TestDeviceGetECCMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetECCMode(device)
	require.NoError(t, err)
}

func TestDeviceGetEncoderCapacity(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	cap, err := w.DeviceGetEncoderCapacity(device, EncoderTypeQueryH264)
	require.NoError(t, err)
	require.NotZero(t, cap)
}

func TestDeviceGetEncoderSessions(t *testing.T) {
	t.Skip()
}

func TestDeviceGetEncoderStats(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, _, err := w.DeviceGetEncoderStats(device)
	require.NoError(t, err)
}

func TestDeviceGetEncoderUtilization(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetEncoderUtilization(device)
	require.NoError(t, err)
}

func TestDeviceGetEnforcedPowerLimit(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	limit, err := w.DeviceGetEnforcedPowerLimit(device)
	require.NoError(t, err)
	require.True(t, limit > 0)
}

func TestDeviceGetFanSpeed(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	speed, err := w.DeviceGetFanSpeed(device)
	require.NoError(t, err)
	require.True(t, speed > 0)
}

func TestDeviceGetGPUOperationMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, _, err := w.DeviceGetGPUOperationMode(device)
	require.NoError(t, err)
}

func TestDeviceGetGraphicsRunningProcesses(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	list, err := w.DeviceGetGraphicsRunningProcesses(device)
	require.NoError(t, err)
	require.NotEmpty(t, list)
}

func TestDeviceGetHandleByIndex(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	device, err := w.DeviceGetHandleByIndex(0)
	require.NoError(t, err)
	require.NotNil(t, device)
}

func TestDeviceGetHandleByPciBusId(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	info, err := w.DeviceGetPCIInfo(device)
	require.NoError(t, err)

	handle, err := w.DeviceGetHandleByPCIBusID(info.BusID)
	require.NoError(t, err)
	require.Equal(t, device, handle)
}

func TestDeviceGetHandleBySerial(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	serial, err := w.DeviceGetSerial(device)
	require.NoError(t, err)

	result, err := w.DeviceGetHandleBySerial(serial)
	require.NoError(t, err)
	require.NotNil(t, err)
	require.Equal(t, device, result)
}

func TestDeviceGetHandleByUUID(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	uuid, err := w.DeviceGetUUID(device)
	require.NoError(t, err)

	result, err := w.DeviceGetHandleByUUID(uuid)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, device, result)
}

func TestDeviceGetIndex(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	index, err := w.DeviceGetIndex(device)
	require.NoError(t, err)
	require.Equal(t, uint32(0), index)
}

func TestDeviceGetInforomConfigurationChecksum(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	checksum, err := w.DeviceGetInforomConfigurationChecksum(device)
	require.NoError(t, err)
	require.True(t, checksum != 0)
}

func TestDeviceGetInforomImageVersion(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	version, err := w.DeviceGetInfoROMImageVersion(device)
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestDeviceGetInfoROMVersion(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	version, err := w.DeviceGetInfoROMVersion(device, InfoROMObjectPower)
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestDeviceGetMaxClockInfo(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clock, err := w.DeviceGetMaxClockInfo(device, ClockMem)
	require.NoError(t, err)
	require.NotZero(t, clock)
}

func TestDeviceGetMaxCustomerBoostClock(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clock, err := w.DeviceGetMaxCustomerBoostClock(device, ClockGraphics)
	require.NoError(t, err)
	require.NotZero(t, clock)
}

func TestDeviceGetMaxPcieLinkGeneration(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	maxLinkGen, err := w.DeviceGetMaxPcieLinkGeneration(device)
	require.NoError(t, err)
	require.NotZero(t, maxLinkGen)
}

func TestDeviceGetMaxPcieLinkWidth(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	maxLinkWidth, err := w.DeviceGetMaxPcieLinkWidth(device)
	require.NoError(t, err)
	require.NotZero(t, maxLinkWidth)
}

func TestDeviceGetMemoryErrorCounter(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	counter, err := w.DeviceGetMemoryErrorCounter(device, MemoryErrorTypeCorrected, VolatileECC, MemoryLocationDeviceMemory)
	require.NoError(t, err)
	require.NotZero(t, counter)
}

func TestDeviceGetMemoryInfo(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	mem, err := w.DeviceGetMemoryInfo(device)
	require.NoError(t, err)

	require.True(t, mem.Free > 0)
	require.True(t, mem.Total > 0)
	require.True(t, mem.Used > 0)
}

func TestDeviceGetMinorNumber(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetMinorNumber(device)
	require.NoError(t, err)
}

func TestDeviceGetName(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	name, err := w.DeviceGetName(device)
	require.NoError(t, err)
	require.NotEmpty(t, name)
}

func TestDeviceGetP2PStatus(t *testing.T) {
	t.Skip()
}

func TestDeviceGetPciInfo(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	info, err := w.DeviceGetPCIInfo(device)
	require.NoError(t, err)
	require.NotNil(t, info)
}

func TestDeviceGetPcieReplayCounter(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetPcieReplayCounter(device)
	require.NoError(t, err)
}

func TestDeviceGetPCIeThroughput(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	value, err := w.DeviceGetPCIeThroughput(device, PCIeUtilTXBytes)
	require.NoError(t, err)
	require.NotZero(t, value)
}

func TestDeviceGetPerformanceState(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetPerformanceState(device)
	require.NoError(t, err)
}

func TestDeviceGetPowerManagementDefaultLimit(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	limit, err := w.DeviceGetPowerManagementDefaultLimit(device)
	require.NoError(t, err)
	require.NotZero(t, limit)
}

func TestDeviceGetPowerManagementLimit(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	limit, err := w.DeviceGetPowerManagementLimit(device)
	require.NoError(t, err)
	require.True(t, limit > 0)
}

func TestDeviceGetPowerManagementLimitConstraints(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	min, max, err := w.DeviceGetPowerManagementLimitConstraints(device)
	require.NoError(t, err)
	require.True(t, min > 0)
	require.True(t, max > 0)
	require.True(t, max > min)
}

func TestDeviceGetPowerManagementMode(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetPowerManagementMode(device)
	require.NoError(t, err)
}

func TestDeviceGetPowerState(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetPowerState(device)
	require.NoError(t, err)
}

func TestDeviceGetPowerUsage(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	power, err := w.DeviceGetPowerUsage(device)
	require.NoError(t, err)
	require.True(t, power > 0)
}

func TestDeviceGetRetiredPages(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetRetiredPages(device, PageRetirementCauseDoubleBitECCError)
	require.NoError(t, err)
}

func TestDeviceGetRetiredPagesPendingStatus(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetRetiredPagesPendingStatus(device)
	require.NoError(t, err)
}

func TestDeviceGetSamples(t *testing.T) {
	t.Skip()
}

func TestDeviceGetSerial(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	serial, err := w.DeviceGetSerial(device)
	require.NoError(t, err)
	require.NotEmpty(t, serial)
}

func TestDeviceGetSupportedClocksThrottleReasons(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	reasons, err := w.DeviceGetSupportedClocksThrottleReasons(device)
	require.NoError(t, err)
	require.NotZero(t, reasons)
}

func TestDeviceGetSupportedGraphicsClocks(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	mem, err := w.DeviceGetSupportedMemoryClocks(device)
	require.NoError(t, err)

	graphics, err := w.DeviceGetSupportedGraphicsClocks(device, mem[0])
	require.NoError(t, err)
	require.NotEmpty(t, graphics)
}

func TestDeviceGetSupportedMemoryClocks(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	clocks, err := w.DeviceGetSupportedMemoryClocks(device)
	require.NoError(t, err)
	require.NotEmpty(t, clocks)

	for i, x := range clocks {
		log.Printf("%d: %d", i, x)
	}
}

func TestDeviceGetTemperature(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	temp, err := w.DeviceGetTemperature(device, TemperatureGPU)
	require.NoError(t, err)
	require.True(t, temp > 0)
}

func TestDeviceGetTemperatureThreshold(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	temp, err := w.DeviceGetTemperatureThreshold(device, TemperatureThresholdShutdown)
	require.NoError(t, err)
	require.True(t, temp > 0)
}

func TestDeviceGetTopologyCommonAncestor(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetTopologyCommonAncestor(device, device)
	require.NoError(t, err)
}

func TestDeviceGetTopologyNearestGpus(t *testing.T) {
	t.Skip()
}

func TestDeviceGetTotalECCErrors(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	_, err := w.DeviceGetTotalECCErrors(device, MemoryErrorTypeCorrected, VolatileECC)
	require.NoError(t, err)
}

func TestDeviceGetTotalEnergyConsumption(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	energy, err := w.DeviceGetTotalEnergyConsumption(device)
	require.NoError(t, err)
	require.NotZero(t, energy)
}

func TestDeviceGetUUID(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	uuid, err := w.DeviceGetUUID(device)
	require.NoError(t, err)
	require.NotEmpty(t, uuid)
}

func TestDeviceGetUtilizationRates(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	rates, err := w.DeviceGetUtilizationRates(device)
	require.NoError(t, err)
	require.NotZero(t, rates.GPU)
	require.NotZero(t, rates.Memory)
}

func TestDeviceGetVbiosVersion(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	version, err := w.DeviceGetVbiosVersion(device)
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestDeviceGetViolationStatus(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	status, err := w.DeviceGetViolationStatus(device, PerfPolicyPower)
	require.NoError(t, err)
	require.NotZero(t, status.ViolationTime)
	require.NotZero(t, status.ReferenceTime)
}

func TestDeviceOnSameBoard(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	s, err := w.DeviceOnSameBoard(device, device)
	require.NoError(t, err)
	require.True(t, s)
}

func TestDeviceSetAutoBoostedClocksEnabled(t *testing.T) {
	w, device := create(t)
	defer w.Shutdown()

	err := w.DeviceSetAutoBoostedClocksEnabled(device, false)
	require.NoError(t, err)
}
