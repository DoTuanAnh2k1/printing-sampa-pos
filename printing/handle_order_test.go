package printing_test

import (
	"fmt"
	"printing-sampa-pos/printing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleOrder(t *testing.T) {
	output, err := printing.HandleOrder(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}
