package printing_test

import (
	"fmt"
	"printing-sampa-pos/printing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleDiscount(t *testing.T) {
	output, err := printing.HandleDiscount(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}
