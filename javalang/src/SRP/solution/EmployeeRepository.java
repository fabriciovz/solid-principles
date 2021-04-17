package SRP.solution;

public class EmployeeRepository {

	public void save(Employee employee) {
		//BD implementation here
		System.out.println("Saving in the Terminal .... " + employee.getEmployeeID());
	}
}
