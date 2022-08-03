package rest

import (
	"GriffinBackend/rdb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getEmployee(c *gin.Context, db *redis.Client) {
	var (
		employeeP []rdb.Employee
		employeeF []rdb.Employee
	)
	err0 := rdb.JsonGet(
		db,
		PERMANENT_EMPLOYEE_KEY,
		PERMANENT_EMPLOYEE_PATH,
		&employeeP,
	)
	err1 := rdb.JsonGet(
		db,
		FREELANCE_EMPLOYEE_KEY,
		FREELANCE_EMPLOYEE_PATH,
		&employeeF,
	)
	if (err0 != nil) || (err1 != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong request parameters",
		})
	}

	totalEmp := rdb.EmployeePlural{
		Total:     len(employeeF) + len(employeeP),
		Permanent: employeeP,
		Freelance: employeeF,
	}
	c.JSON(http.StatusOK, totalEmp)
}

func postEmployee(c *gin.Context, db rdb.DataBase) {

}

func deleteEmployee(c *gin.Context, db rdb.DataBase) {

}

func updateEmployee(c *gin.Context, db rdb.DataBase) {

}
