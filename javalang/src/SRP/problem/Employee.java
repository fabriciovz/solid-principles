package SRP.problem;

public class Employee {
	private String employeeID;         
	private String employeeName;
	private String employeeAddress;
	private String contactNumber;
	private String employeeType;

	public void save() {
		//BD implementation here
		System.out.println("Saving in the Terminal .... " + this.getEmployeeID());
	}
	public void calculateTax() {

		if (this.getEmployeeType().equals("fulltime")) {
			// Task calc for full time employee
		}
		if (this.getEmployeeType().equals("contract")) {
			// Task calc for contractors
		}
	}
	public String getEmployeeID() {
		return this.employeeID;
	}	
	public void setEmployeeID(String ID) {
		this.employeeID=ID;
	}
	public String getEmployeeName() {
		return this.employeeName;
	}
	public void setEmployeeName(String employeeName) {
		this.employeeName = employeeName;
	}
	public String getEmployeeAddress() {
		return this.employeeAddress;
	}
	public void SetEmployeeAddress(String employeeAddress) {
		this.employeeAddress = employeeAddress;
	}
	public String GetContactNumber() {
		return this.contactNumber;
	}
	public void SetContactNumber(String contactNumber) {
		this.contactNumber = contactNumber;
	}
	public String getEmployeeType() {
		return this.employeeType;
	}
	public void SetEmployeeType(String employeeType) {
		this.employeeType = employeeType;
	}
}