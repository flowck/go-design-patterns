package publication

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPublication(t *testing.T) {
	mag, err := NewPublication(Magazine, "Vogue", 10, "Vogue")

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s", mag), "This is a magazine named Vogue")

	pub, err := NewPublication(Newspaper, "NY Times", 10, "NY Times")

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s", pub), "This is a newspaper named NY Times")
}
