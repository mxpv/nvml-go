package nvml

import (
	"C"
	"unsafe"
)

// SystemGetCudaDriverVersion retrieves the version of the CUDA driver.
// The returned CUDA driver version is the same as the CUDA API cuDriverGetVersion() would return on the system.
func (w wrapper) SystemGetCudaDriverVersion() (cudaDriverVersion int, err error) {
	err = w.call(w.nvmlSystemGetCudaDriverVersion, uintptr(unsafe.Pointer(&cudaDriverVersion)))
	return
}

// SystemGetDriverVersion retrieves the version of the system's graphics driver.
func (w wrapper) SystemGetDriverVersion() (string, error) {
	buffer := [nvmlSystemDriverVersionBufferSize]C.char{}
	if err := w.call(w.nvmlSystemGetDriverVersion, uintptr(unsafe.Pointer(&buffer[0])), nvmlSystemDriverVersionBufferSize); err != nil {
		return "", err
	}

	return C.GoString(&buffer[0]), nil
}

// SystemGetNVMLVersion retrieves the version of the NVML library.
func (w wrapper) SystemGetNVMLVersion() (string, error) {
	buffer := [nvmlSystemDriverVersionBufferSize]C.char{}
	if err := w.call(w.nvmlSystemGetNVMLVersion, uintptr(unsafe.Pointer(&buffer[0])), nvmlSystemDriverVersionBufferSize); err != nil {
		return "", err
	}

	return C.GoString(&buffer[0]), nil
}

// SystemGetProcessName gets name of the process with provided process id
func (w wrapper) SystemGetProcessName(pid uint) (string, error) {
	const maxLength = 256

	buffer := [maxLength]C.char{}
	if err := w.call(w.nvmlSystemGetProcessName, uintptr(pid), uintptr(unsafe.Pointer(&buffer[0])), maxLength); err != nil {
		return "", err
	}

	return C.GoString(&buffer[0]), nil
}
