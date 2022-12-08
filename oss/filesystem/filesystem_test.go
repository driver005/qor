package filesystem

import (
	"testing"

	"github.com/driver005/commerce/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}
