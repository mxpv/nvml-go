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

func TestSystemGetCudaDriverVersion(t *testing.T) {
	w := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetCudaDriverVersion()
	require.NoError(t, err)
	require.True(t, version > 0)
}

func TestSystemGetDriverVersion(t *testing.T) {
	w := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetDriverVersion()
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestSystemGetNVMLVersion(t *testing.T) {
	w := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetNVMLVersion()
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestSystemGetProcessName(t *testing.T) {
	w := create(t)
	defer w.Shutdown()

	name, err := w.SystemGetProcessName(1336)
	require.NoError(t, err)
	require.NotEmpty(t, name)
}

func create(t *testing.T) *wrapper {
	w, err := New("")
	require.NoError(t, err)

	err = w.Init()
	require.NoError(t, err)

	return w
}
