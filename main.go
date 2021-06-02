package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setHeaders(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST,GET")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		h(c)
		return nil
	}
}

type User struct {
	ID         int    `json:"id" form:"id"`
	FirstName  string `json:"firstName" form:"firstName"`
	LastName   string `json:"lastName" form:"lastName"`
	Email      string `json:"email" form:"email"`
	CreateDate string `json:"createDate" form:"createDate"`
}

func addUser(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	email := c.Param("email")

	response := &User{
		ID:         142,
		Email:      email,
		CreateDate: "2019-08-24",
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.GET("/users/:email", getUser)
	e.POST("/users", addUser)
	e.Logger.Fatal(e.Start(":3000"))
}
