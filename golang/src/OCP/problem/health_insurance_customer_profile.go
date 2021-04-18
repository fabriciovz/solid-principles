package problem_ocp

import (
	"math/rand"
	"time"
)

type HealthInsuranceCustomerProfile struct {
}

func NewHealthInsuranceCustomerProfile() *HealthInsuranceCustomerProfile {
	return &HealthInsuranceCustomerProfile{}
}

func (h *HealthInsuranceCustomerProfile) IsLoyalCustomer() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
