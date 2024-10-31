package printing_test

import (
	"parsing-template/printing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerTag(t *testing.T) {
	output := printing.HandlerTag("<C10>ajdhaf")
	require.Equal(t, output, "C10")
	output = printing.HandlerTag("<C> asdwoeiur")
	require.Equal(t, output, "C")
	output = printing.HandlerTag("<F> jkadshfl")
	require.Equal(t, output, "F")
}

func TestHandlerBar(t *testing.T) {
	output := printing.HandlerBar("-")
	require.Equal(t, output, printing.Dash)
	output = printing.HandlerBar("=")
	require.Equal(t, output, printing.DoubleDash)
	output = printing.HandlerBar(" ")
	require.Equal(t, output, printing.Space)
}

func TestHandlerHeader(t *testing.T) {
	tag, body, height, width, err := printing.HandlerHeader("<C10>ajdhaf")
	require.NoError(t, err)
	require.Equal(t, "C", tag)
	require.Equal(t, "ajdhaf", body)
	require.Equal(t, 1, height)
	require.Equal(t, 0, width)

	tag, body, height, width, err = printing.HandlerHeader("<C>ajdhaf")
	require.NoError(t, err)
	require.Equal(t, "C", tag)
	require.Equal(t, "ajdhaf", body)
	require.Equal(t, -1, height)
	require.Equal(t, -1, width)

	_, _, _, _, err = printing.HandlerHeader("<Cxy>ajdhaf")
	require.EqualError(t, err, "invalid number")
}