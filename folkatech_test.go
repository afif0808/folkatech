package folkatech

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	shares := []int{13, 10, 14, 3, 13, 7, 18, 4, 19}
	profit := MaxProfit(shares, 3)
	assert.Equal(t, 36, profit, "Wrong answer")
}
