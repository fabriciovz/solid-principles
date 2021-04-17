package SRP.solution;

public class Employee {
	private String employeeID;         
	private String employeeName;
	private String employeeAddress;
	private String contactNumber;
	private String employeeType;

	public void save() {
		new EmployeeRepository().save(this);
	}
	public void calculateTax() {
		new TaxCalculator().calculateTax(this);
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
	public void setEmployeeAddress(String employeeAddress) {
		this.employeeAddress = employeeAddress;
	}
	public String getContactNumber() {
		return this.contactNumber;
	}
	public void setContactNumber(String contactNumber) {
		this.contactNumber = contactNumber;
	}
	public String getEmployeeType() {
		return this.employeeType;
	}
	public void setEmployeeType(String employeeType) {
		this.employeeType = employeeType;
	}
}