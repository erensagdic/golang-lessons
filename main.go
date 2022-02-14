package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var e = echo.New()

//test
func main() {
	e.GET("/", mainHandler)
	e.GET("/users", usersHandler)
	e.POST("/login", loginHandler)
	runServer()
}

func loginHandler(context echo.Context) error {
	username := context.FormValue("username")
	nameSurname := context.FormValue("nameSurname")
	email := context.FormValue("email")
	password := context.FormValue("password")

	if (username == "erensagdicc" || email == "erensagdicc@gmail.com") && password == "erensagdicc123" {
		response := "Hoş Geldin " + nameSurname

		return context.String(http.StatusOK, response)
	}

	return context.String(http.StatusForbidden, "Buraya girmeye yetkiniz yok.")
}

func mainHandler(context echo.Context) error {
	return context.String(http.StatusOK, "Welcome to home page.")
}

func usersHandler(context echo.Context) error {
	username := context.QueryParam("username")
	nameSurname := context.QueryParam("nameSurname")
	email := context.QueryParam("email")
	password := context.QueryParam("password")

	if (username == "erensagdicc" || email == "erensagdicc@gmail.com") && password == "erensagdicc123" {
		response := "Hoş Geldin " + nameSurname

		return context.String(http.StatusOK, response)
	}

	return context.String(http.StatusForbidden, "Buraya girmeye yetkiniz yok.")

}

func runServer() {
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
