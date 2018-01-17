package nvml

import (
	"C"
	"fmt"
	"syscall"
	"unsafe"
)

type wrapper struct {
	dll *syscall.DLL
	nvmlInit,
	nvmlShutdown,
	nvmlErrorString *syscall.Proc
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

func New(path string) (*wrapper, error) {
	if path == "" {
		path = "C:\\Program Files\\NVIDIA Corporation\\NVSMI\\nvml.dll"
	}

	dll, err := syscall.LoadDLL(path)
	if err != nil {
		return nil, err
	}

	bindings := &wrapper{
		dll:             dll,
		nvmlInit:        dll.MustFindProc("nvmlInit"),
		nvmlShutdown:    dll.MustFindProc("nvmlShutdown"),
		nvmlErrorString: dll.MustFindProc("nvmlErrorString"),
	}

	return bindings, nil
}
