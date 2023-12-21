package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const TIKA = "http://localhost:9998/meta"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/metadata", func(c echo.Context) error {
		fmt.Println("fishman")
		files, err := c.FormFile("form")
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("fishman2")

		src, err := files.Open()
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer src.Close()
		fmt.Println("fishman3")

		req, err := http.NewRequest("PUT", TIKA, src)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("fishman4")
		// req.Header.Set("Content-Type", "multipart/form")
		req.Header.Set("Accept", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("fishman5")

		defer resp.Body.Close()
		fmt.Println("fishman6")

		// print the body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusOK, string(body))

	})
	e.Logger.Fatal(e.Start(":1323"))
}

// curl -T README.md http://localhost:9998/meta/Content-Type --header "Accept: application/json"
