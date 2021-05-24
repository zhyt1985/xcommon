package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number struct {
	OrderAmt string  `fake:"func(RandDate(2016-05-06,2020-05-02,date))"`
	Count    int64   `fake:"func(RandIntRangeBetween(30000,50000))"`
	Percent  float64 `fake:"func(RandFloatRangeRand(4))"`
	Code     string  `fake:"func(RandEnum(110000|120000|130000|150000|170000))"`
}
type Default struct {
	Count int `default:"1"`
}

func init() {
	Seed(0)
	RegisterFakes([]Func{
		{
			TagName: "RandIntRangeBetween",
			Call:    RandIntRangeBetween,
		},
		{
			TagName: "RandDate",
			Call:    RandDate,
		},
		{
			TagName: "RandIntRangeRand",
			Call:    RandIntRangeRand,
		},
		{
			TagName: "RandFloatRangeRand",
			Call:    RandFloatRangeRand,
		},
		{
			TagName: "RandEnum",
			Call:    RandEnum,
		},
		{
			TagName: "RandDate",
			Call:    RandDate,
		},
	})
}
func TestFake(t *testing.T) {
	assert := assert.New(t)
	order := Number{}
	Fake(&order)
	fmt.Println(order)
	assert.NotNil(order.Count)
}
func TestDefault(t *testing.T) {
	assert := assert.New(t)
	order := Default{}
	Fake(&order)
	fmt.Println(order)
	assert.NotNil(order.Count)
}
