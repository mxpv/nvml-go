package nvml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	w, err := New("")
	require.NoError(t, err)

	err = w.Init()
	require.NoError(t, err)

	err = w.Shutdown()
	require.NoError(t, err)
}

func TestErrorString(t *testing.T) {
	w, err := New("")
	require.NoError(t, err)
	defer w.Shutdown()

	require.Equal(t, "Uninitialized", w.ErrorString(1))
	require.Equal(t, "Invalid Argument", w.ErrorString(2))
	require.Equal(t, "Not Supported", w.ErrorString(3))
}

func create(t *testing.T) (*API, Device) {
	w, err := New("")
	require.NoError(t, err)

	err = w.Init()
	require.NoError(t, err)

	device, err := w.DeviceGetHandleByIndex(0)
	require.NoError(t, err)

	return w, device
}
