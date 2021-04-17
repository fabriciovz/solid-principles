package employee_solution

type Employee struct {
	employeeID         string
	employeeName       string
	employeeAddress    string
	contactNumber      string
	employeeType       string
	employeeRepository EmployeeRepository
	taxCalculator      TaxCalculator
}

func (e *Employee) Save() {
	e.employeeRepository.Save(e)
}
func (e *Employee) CalculateTax() {
	e.taxCalculator.CalculateTax(e)
}
func (e *Employee) GetEmployeeID() string {
	return e.employeeID
}
func (e *Employee) SetEmployeeID(ID string) {
	e.employeeID = ID
}
func (e *Employee) GetEmployeeName() string {
	return e.employeeName
}
func (e *Employee) SetEmployeeName(employeeName string) {
	e.employeeName = employeeName
}
func (e *Employee) GetEmployeeAddress() string {
	return e.employeeAddress
}
func (e *Employee) SetEmployeeAddress(employeeAddress string) {
	e.employeeAddress = employeeAddress
}
func (e *Employee) GetContactNumber() string {
	return e.contactNumber
}
func (e *Employee) SetContactNumber(contactNumber string) {
	e.contactNumber = contactNumber
}
func (e *Employee) GetEmployeeType() string {
	return e.employeeType
}
func (e *Employee) SetEmployeeType(employeeType string) {
	e.employeeType = employeeType
}
