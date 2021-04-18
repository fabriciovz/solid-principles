package problem

type InsurancePremiumDiscountCalculator struct {
}

func (i *InsurancePremiumDiscountCalculator) calculatePremiumDiscountPercent(customer CustomerProfile) int{
	if customer.IsLoyalCustomer() {
		return 20
	}
	return 0
}