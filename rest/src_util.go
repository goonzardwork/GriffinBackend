package rest

import (
	"GriffinBackend/rdb"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(ALLOW_ORIGIN, ALLOW_ORIGIN_VALUE)
		c.Header(ALLOW_CREDENTIALS, ALLOW_CREDENTIALS_VALUE)
		c.Next()
	}
}

func processEmployee(es []rdb.Employee) []rdb.Employee {
	var result []rdb.Employee
	for _, r := range es {
		result = append(result, calcTimeLeft(r))
	}
	return result
}

func calcTimeLeft(w rdb.Employee) rdb.Employee {
	var payTime time.Time
	d := w.Date
	curTime := time.Now()
	payTime = time.Date(curTime.Year(), curTime.Month(), d, 10, 0, 0, 0, curTime.Location())
	if curTime.After(payTime) {
		payTime = payTime.AddDate(0, 1, 0)
	}
	fmt.Println(payTime, curTime)
	timeDiff := payTime.Sub(curTime)
	w.TimeLeft = timeDiff.Seconds()
	fmt.Println(w)
	return w
}
