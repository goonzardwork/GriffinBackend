package rdb

type EmployeePlural struct {
	Total     int        `json:"total_length"`
	Permanent []Employee `json:"repeat_true"`
	Freelance []Employee `json:"repeat_false"`
}

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
	Account  string `json:"account"`
	Payroll  int    `json:"payroll"`
	Date     int    `json:"int"`
}
