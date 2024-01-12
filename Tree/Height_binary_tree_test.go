package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeight(t *testing.T) {
	root := Initialize()
	actual := Height(&root)
	assert.Equal(t, 3, actual)

}
