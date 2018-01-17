package nvml

import (
	"C"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	nvmlSystemDriverVersionBufferSize = 80
	nvmlDeviceNameBufferSize          = 64
)

type wrapper struct {
	dll *syscall.DLL
	// Initialization and cleanup
	nvmlInit,
	nvmlShutdown,
	// Error reporting
	nvmlErrorString,
	// System Queries
	nvmlSystemGetCudaDriverVersion,
	nvmlSystemGetDriverVersion,
	nvmlSystemGetNVMLVersion,
	nvmlSystemGetProcessName *syscall.Proc
}

func (w wrapper) call(p *syscall.Proc, a ...uintptr) error {
	ret, _, _ := p.Call(a...)
	if ret != 0 {
		return fmt.Errorf("%s call failed with error: %d %s", p.Name, ret, w.ErrorString(ret))
	}

	return nil
}

// Init initializes NVML, but don't initialize any GPUs yet.
func (w wrapper) Init() error {
	return w.call(w.nvmlInit)
}

// Shutdown shut downs NVML by releasing all GPU resources previously allocated with Init() and
// unloads nvml.dll via UnloadLibrary call.
func (w wrapper) Shutdown() error {
	err := w.call(w.nvmlShutdown)
	w.dll.Release()
	return err
}

// ErrorString returns a string representation of the error.
func (w wrapper) ErrorString(result uintptr) string {
	ret, _, _ := w.nvmlErrorString.Call(uintptr(result))
	buf := (*C.char)(unsafe.Pointer(ret))
	return C.GoString(buf)
}

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

// New creates nvml.dll wrapper
func New(path string) (*wrapper, error) {
	if path == "" {
		path = "C:\\Program Files\\NVIDIA Corporation\\NVSMI\\nvml.dll"
	}

	dll, err := syscall.LoadDLL(path)
	if err != nil {
		return nil, err
	}

	bindings := &wrapper{
		dll:                            dll,
		nvmlInit:                       dll.MustFindProc("nvmlInit"),
		nvmlShutdown:                   dll.MustFindProc("nvmlShutdown"),
		nvmlErrorString:                dll.MustFindProc("nvmlErrorString"),
		nvmlSystemGetCudaDriverVersion: dll.MustFindProc("nvmlSystemGetCudaDriverVersion"),
		nvmlSystemGetDriverVersion:     dll.MustFindProc("nvmlSystemGetDriverVersion"),
		nvmlSystemGetNVMLVersion:       dll.MustFindProc("nvmlSystemGetNVMLVersion"),
		nvmlSystemGetProcessName:       dll.MustFindProc("nvmlSystemGetProcessName"),
	}

	return bindings, nil
}
