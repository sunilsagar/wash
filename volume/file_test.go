package volume

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/puppetlabs/wash/plugin"
	"github.com/stretchr/testify/assert"
)

func TestVolumeFile(t *testing.T) {
	now := time.Now()
	initialAttr := plugin.EntryAttributes{}
	initialAttr.SetCtime(now)
	vf := NewFile("mine", initialAttr, func(ctx context.Context, path string) (plugin.SizedReader, error) {
		assert.Equal(t, "my path", path)
		return strings.NewReader("hello"), nil
	}, "my path")

	attr := plugin.Attributes(vf)
	expectedAttr := plugin.EntryAttributes{}
	expectedAttr.SetCtime(now)
	assert.Equal(t, expectedAttr, attr)
	rdr, err := vf.Open(context.Background())
	assert.Nil(t, err)
	if assert.NotNil(t, rdr) {
		buf := make([]byte, rdr.Size())
		n, err := rdr.ReadAt(buf, 0)
		assert.Nil(t, err)
		assert.Equal(t, int64(n), rdr.Size())
		assert.Equal(t, "hello", string(buf))
	}
}

func TestVolumeFileErr(t *testing.T) {
	vf := NewFile("mine", plugin.EntryAttributes{}, func(ctx context.Context, path string) (plugin.SizedReader, error) {
		return nil, errors.New("fail")
	}, "my path")

	rdr, err := vf.Open(context.Background())
	assert.Nil(t, rdr)
	assert.Equal(t, errors.New("fail"), err)
}
