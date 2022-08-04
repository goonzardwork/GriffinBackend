package rest

const (
	// function get employer query
	EMPLOYER_ID = "employerId"
	// function new employer query
	EMPLOYEE_NAME     = "name"
	EMPLOYEE_EMAIL    = "email"
	EMPLOYEE_POSITION = "position"
	EMPLOYEE_ACCOUNT  = "account"
	EMPLOYEE_PAYROLL  = "payroll"
	EMPLOYEE_PAYDAY   = "date"
	// function new employer query index
	EMPLOYEE_NAME_INDEX     = 0
	EMPLOYEE_EMAIL_INDEX    = 1
	EMPLOYEE_POSITION_INDEX = 2
	EMPLOYEE_ACCOUNT_INDEX  = 3
	EMPLOYEE_PAYROLL_INDEX  = 4
	EMPLOYEE_PAYDAY_INDEX   = 5
	// database access keys
	PERMANENT_EMPLOYER_KEY  = "employee_perma:%v"
	PERMANENT_EMPLOYER_PATH = "$"
	FREELANCE_EMPLOYER_KEY  = "employee_free:%v"
	FREELANCE_EMPLOYER_PATH = "$"
)
