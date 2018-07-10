package readtypes_test

import (
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	"github.com/amonks/typeshift/readtypes"
)

func TestReadtypes(t *testing.T) {
	snapshotter := snapshotter.New(t)
	defer snapshotter.Verify()
	testcases := []string{
		"arrays",
		"nullables",
		"objects",
		"scalars",
	}
	for _, testcase := range testcases {
		s, err := readtypes.ReadPackageDirectoryTypes("../testpackages/" + testcase)
		assert.NoError(t, err)
		snapshotter.Snapshot(testcase, s)
	}
}
