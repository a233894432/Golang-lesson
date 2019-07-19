/**
* curl http://localhost:8080/
* curl --data "users=eden, eden2" http://localhost:8080/import
* curl http://localhost:8080/lucky
 */

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var userList []string

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()

	userList = []string{}

	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("total: %d\n", count)
}

// POST http://localhost:8080/import
// params: users
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	beforeImportedUsersCount := len(userList)

	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	afterImportedUsersCount := len(userList)
	return fmt.Sprintf("total: %d, success import: %d\n", beforeImportedUsersCount, afterImportedUsersCount)
}

func (c *lotteryController) GetLucky() string {
	userCount := len(userList)

	if userCount > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(userCount))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)

		return fmt.Sprintf("lucky user is %s, remining %d", user, userCount-1)
	} else if userCount == 1 {
		user := userList[0]

		return fmt.Sprintf("lucky user is %s, remaining %d", user, userCount-1)
	} else {

		return fmt.Sprint("no more user")
	}
}
