package nvml

import (
	"log"
)

func ExampleNew() {
	nvml, err := New("")
	if err != nil {
		panic(err)
	}

	defer nvml.Shutdown()

	err = nvml.Init()
	if err != nil {
		panic(err)
	}

	driverVersion, err := nvml.SystemGetDriverVersion()
	if err != nil {
		panic(err)
	}

	log.Printf("Driver version:\t%s", driverVersion)

	nvmlVersion, err := nvml.SystemGetNVMLVersion()
	if err != nil {
		panic(err)
	}

	log.Printf("NVML version:\t%s", nvmlVersion)

	deviceCount, err := nvml.DeviceGetCount()
	if err != nil {
		panic(err)
	}

	for i := uint32(0); i < deviceCount; i++ {
		handle, err := nvml.DeviceGetHandleByIndex(i)
		if err != nil {
			panic(err)
		}

		name, err := nvml.DeviceGetName(handle)
		log.Printf("Product name:\t%s", name)

		brand, err := nvml.DeviceGetBrand(handle)
		if err != nil {
			panic(err)
		}

		log.Printf("Product Brand:\t%s", brand)

		uuid, err := nvml.DeviceGetUUID(handle)
		if err != nil {
			panic(err)
		}

		log.Printf("GPU UUID:\t\t%s", uuid)

		fan, err := nvml.DeviceGetFanSpeed(handle)
		if err != nil {
			panic(err)
		}

		log.Printf("Fan Speed:\t\t%d", fan)
	}
}
