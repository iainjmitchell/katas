package shopping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var prices = map[string]int{
	"A": 50,
	"B": 30,
	"C": 20,
	"D": 15,
}

var discounts = map[string]Discount{
	"A": {quantity: 3, amount: 30},
	"B": {quantity: 2, amount: 15},
}

func FuzzSingleItemTest(fuzzyTest *testing.F) {
	for testCase := range prices {
		fuzzyTest.Add(testCase)
	}
	fuzzyTest.Fuzz(func(test *testing.T, sku string) {
		cart := []string{}
		cart = append(cart, sku)
		till := Till{prices: prices, discounts: discounts}
		assert.Equal(test, prices[sku], till.Total(cart))
	})
}

func TestTwoAIs100(test *testing.T) {
	const SKU = "A"
	cart := []string{}
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	till := Till{prices: prices, discounts: discounts}
	assert.Equal(test, 100, till.Total(cart))
}

func TestSpecialOfferOfThreeAFor120(test *testing.T) {
	const SKU = "A"
	cart := []string{}
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	till := Till{prices: prices, discounts: discounts}
	assert.Equal(test, 120, till.Total(cart))
}

func TestNoSpecialOfferThreeCFor60(test *testing.T) {
	const SKU = "C"
	cart := []string{}
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	till := Till{prices: prices, discounts: discounts}
	assert.Equal(test, 60, till.Total(cart))
}

func TestMultipleSpecialOffersOfSixAFor240(test *testing.T) {
	const SKU = "A"
	cart := []string{}
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	cart = append(cart, SKU)
	till := Till{prices: prices, discounts: discounts}
	assert.Equal(test, 240, till.Total(cart))
}

func TestMixedBasket(test *testing.T) {
	cart := []string{}
	cart = append(cart, "A")
	cart = append(cart, "A")
	cart = append(cart, "B")
	cart = append(cart, "B")
	cart = append(cart, "A")
	cart = append(cart, "D")
	cart = append(cart, "C")
	till := Till{prices: prices, discounts: discounts}
	assert.Equal(test, 200, till.Total(cart))
}
