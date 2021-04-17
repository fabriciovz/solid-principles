package problem

type InsurancePremiumDiscountCalculator struct {
}

func (i *InsurancePremiumDiscountCalculator) calculatePremiumDiscountPercent(customer HealthInsuranceCustomerProfile) int{
	if customer.IsLoyalCustomer() {
		return 20
	}
	return 0
}
// Go not support overloading of methods - see: https://golang.org/doc/faq#overloading
func (i *InsurancePremiumDiscountCalculator) calculatePremiumDiscountPercentVehicle(customer VehicleInsuranceCustomerProfile) int{
	if customer.IsLoyalCustomer() {
		return 20
	}
	return 0
}