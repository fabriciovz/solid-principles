package employee_problem

import "fmt"

type Employee struct {
	employeeID      string
	employeeName    string
	employeeAddress string
	contactNumber   string
	employeeType    string
}

func (e *Employee) Save() {
	//BD implementation here
	fmt.Println("Saving in the Terminal .... " + e.GetEmployeeID())
}

func (e *Employee) CalculateTax() {
	if e.GetEmployeeType() == "fulltime" {
		// Task calc for full time employee
	}
	if e.GetEmployeeType() == "contract" {
		// Task calc for contractors
	}
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
