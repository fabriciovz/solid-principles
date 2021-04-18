package problem

import (
	"math/rand"
	"time"
)

type HealthInsuranceCustomerProfile struct {
}

func NewHealthInsuranceCustomerProfile() CustomerProfile {
	return &HealthInsuranceCustomerProfile{}
}

func (h *HealthInsuranceCustomerProfile) IsLoyalCustomer() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}