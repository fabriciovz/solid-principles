package employee_solution

import "fmt"

type EmployeeRepository struct {
}

func (r *EmployeeRepository) Save(e *Employee) {
	//BD implementation here
	fmt.Println("Saving in the Terminal .... " + e.GetEmployeeID())
}
