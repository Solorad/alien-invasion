package invasion

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAlienNamesArray(t *testing.T) {
	names := BuildAlienNamesArray(10)
	assert.Equal(t, 10, len(names))
	result, err := json.Marshal(names)
	assert.Nil(t, err)
	fmt.Println(string(result))
}
