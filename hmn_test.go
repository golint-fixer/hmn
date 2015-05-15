package hmn

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	AB int
	BC float64
	CD string
	DE uint
	Date time.Time
}

func TestLoad(t *testing.T) {
	assert := assert.New(t)

	i := `a_b: 4 b_c: 8.0 c_d: hello date: 2011-01-21 d_e: 7`

	date, _ := time.Parse("2006-01-02", "2011-01-21")

	e := Example{}
	Load(&e, i)

	assert.Equal(4, e.AB, "e.AB is equal to 4")
	assert.Equal(8.0, e.BC, "e.BC is equal to 8.0")
	assert.Equal(8.0, e.BC, "e.BC is equal to 8.0")
	assert.EqualValues(7, e.DE, "e.DE is equal to 7")
	assert.Equal("hello", e.CD, "e.CD is equal to `hello`")
	assert.Equal(date, e.Date, "e.Date is equal to `2011-01-21`")
}
