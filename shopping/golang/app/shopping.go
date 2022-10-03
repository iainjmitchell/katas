package shopping

type Discount struct {
	quantity int
	amount   int
}

type Till struct {
	prices    map[string]int
	discounts map[string]Discount
}

func (till *Till) Total(cart []string) int {
	var total int
	itemCount := map[string]int{}
	for _, sku := range cart {
		itemCount[sku] += 1
		total += till.prices[sku]
	}
	for sku, count := range itemCount {
		discount, containsDiscount := till.discounts[sku]
		if containsDiscount && count%discount.quantity == 0 {
			numberOfDiscounts := count / discount.quantity
			total -= (discount.amount * numberOfDiscounts)
		}
	}
	return total
}
