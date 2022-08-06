package rest

import (
	"GriffinBackend/rdb"
	"fmt"
	"reflect"
	"strings"
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

func structIter[T any](ct T) [][]string {
	v := reflect.ValueOf(ct)
	typeOfS := v.Type()

	params := [][]string{}
	for i := 0; i < v.NumField(); i++ {
		name := typeOfS.Field(i).Name
		d := v.Field(i).Interface().(string)
		params = append(params, []string{strings.ToLower(name), d})
	}
	return params
}

func isRegistered(info []rdb.Login, username, password string) (bool, bool) {
	usrResult := false
	pwdResult := false
	for _, usr := range info {
		regUser := username == usr.Username
		regPwd := password == usr.Password
		if regUser {
			usrResult = true
			if regPwd {
				pwdResult = true
			}
		}
	}
	return usrResult, pwdResult
}
