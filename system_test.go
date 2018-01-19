package nvml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemGetCudaDriverVersion(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetCudaDriverVersion()
	require.NoError(t, err)
	require.True(t, version > 0)
}

func TestSystemGetDriverVersion(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetDriverVersion()
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestSystemGetNVMLVersion(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	version, err := w.SystemGetNVMLVersion()
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestSystemGetProcessName(t *testing.T) {
	w, _ := create(t)
	defer w.Shutdown()

	name, err := w.SystemGetProcessName(1336)
	require.NoError(t, err)
	require.NotEmpty(t, name)
}
