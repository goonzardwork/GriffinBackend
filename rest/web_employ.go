package rest

import (
	"GriffinBackend/rdb"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getEmployee(c *gin.Context, db *redis.Client) {
	var (
		employeeP [][]rdb.Employee
		employeeF [][]rdb.Employee

		p []rdb.Employee
		f []rdb.Employee
	)

	eid, err := handleGetEmployee(c)
	if err != nil {
		return
	}
	// unravel database
	err0 := rdb.JsonGet(
		db,
		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, eid),
		PERMANENT_EMPLOYER_PATH,
		&employeeP,
	)
	err1 := rdb.JsonGet(
		db,
		fmt.Sprintf(FREELANCE_EMPLOYER_KEY, eid),
		FREELANCE_EMPLOYER_PATH,
		&employeeF,
	)
	switch {
	case err0 != nil:
		p, f = []rdb.Employee{}, employeeF[0]
	case err1 != nil:
		p, f = employeeP[0], []rdb.Employee{}
	case (err0 != nil) && (err1 != nil):
		p, f = []rdb.Employee{}, []rdb.Employee{}
	default:
		p, f = employeeP[0], employeeF[0]
	}
	// send it to webserver
	totalEmp := rdb.EmployeePlural{
		Total:     len(employeeF[0]) + len(employeeP[0]),
		Permanent: p,
		Freelance: f,
	}
	c.JSON(http.StatusOK, totalEmp)
}

func handleGetEmployee(c *gin.Context) (string, error) {
	// check for employer id
	q, ok := c.GetQuery(EMPLOYER_ID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameter",
		})
		return "", errors.New("missing necessary parameter")
	}
	return q, nil
}

func postEmployeePermanent(c *gin.Context, db *redis.Client) {
	newEmployee, err := handleNewEmployee(c)
	if err != nil {
		return
	}
	eid, err := handleGetEmployee(c)
	if err != nil {
		return
	}
	l, err := rdb.JsonArrLen(
		db,
		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, eid),
		PERMANENT_EMPLOYER_PATH,
	)
	switch {
	case err != nil:
		newEmployee.Id = 1
		_ = rdb.JsonSet(
			db,
			fmt.Sprintf(PERMANENT_EMPLOYER_KEY, eid),
			PERMANENT_EMPLOYER_PATH,
			&newEmployee,
		)
		c.JSON(http.StatusOK, gin.H{
			"message": "new employer appended with new worker",
		})
	default:
		newEmployee.Id = l + 1
		err = rdb.JsonArrAppend(
			db,
			fmt.Sprintf(PERMANENT_EMPLOYER_KEY, eid),
			PERMANENT_EMPLOYER_PATH,
			&newEmployee,
		)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"message": "database appending failed",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "new permanent worker appended",
			})
		}
	}

}

func postEmployeeFreelance(c *gin.Context, db *redis.Client) {

}

func handleNewEmployee(c *gin.Context) (rdb.Employee, error) {
	needQuery := []string{
		EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_POSITION,
		EMPLOYEE_ACCOUNT, EMPLOYEE_PAYROLL, EMPLOYEE_PAYDAY,
	}
	needValue := []string{}
	for _, qs := range needQuery {
		q, ok := c.GetQuery(qs)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing parameter",
			})
			return rdb.Employee{}, errors.New("missing necessary parameter")
		}
		needValue = append(needValue, q)
	}

	payroll, err := strconv.Atoi(needValue[EMPLOYEE_PAYROLL_INDEX])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "payroll should be integer",
		})
		return rdb.Employee{}, errors.New("wrong typed parameter")
	}
	date, err := strconv.Atoi(needValue[EMPLOYEE_PAYDAY_INDEX])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "date should be integer(UNIX)",
		})
		return rdb.Employee{}, errors.New("wrong typed parameter")
	}

	return rdb.Employee{
		Name:     needValue[EMPLOYEE_NAME_INDEX],
		Email:    needValue[EMPLOYEE_EMAIL_INDEX],
		Position: needValue[EMPLOYEE_POSITION_INDEX],
		Account:  needValue[EMPLOYEE_ACCOUNT_INDEX],
		Payroll:  payroll,
		Date:     date,
	}, nil
}

func deleteEmployee(c *gin.Context, db *redis.Client) {

}

func updateEmployee(c *gin.Context, db rdb.DataBase) {

}
