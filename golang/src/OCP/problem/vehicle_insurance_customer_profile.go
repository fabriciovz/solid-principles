package problem_ocp

import (
	"math/rand"
	"time"
)

type VehicleInsuranceCustomerProfile struct {
}

func NewVehicleInsuranceCustomerProfile() *VehicleInsuranceCustomerProfile {
	return &VehicleInsuranceCustomerProfile{}
}

func (h *VehicleInsuranceCustomerProfile) IsLoyalCustomer() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

