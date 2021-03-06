package journal

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPIDTOID(t *testing.T) {
	id := PIDToID(500000)
	assert.Equal(t, "500000", id)

	id = PIDToID(os.Getpid())
	expected := strconv.FormatInt(int64(os.Getpid()), 10) + "-journal.test"
	assert.Equal(t, expected, id)
}
