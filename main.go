package main

import (
	"fmt"
	"jira_api/controllers"
)

func main() {
	fmt.Println("Hello JIRA")
	controllers.StartServer()
}
