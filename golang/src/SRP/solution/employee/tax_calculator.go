package solution_srp

type TaxCalculator struct {
}

func (r *TaxCalculator) CalculateTax(e *Employee) {
	if e.GetEmployeeType() == "fulltime" {
		// Task calc for full time employee
	}
	if e.GetEmployeeType() == "contract" {
		// Task calc for contractors
	}
}
