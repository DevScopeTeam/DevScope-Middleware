package main

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"
	"fmt"

	"github.com/google/uuid"
)

func main2() {
	domains := []string{
		"artificial-intelligence",
		"machine-learning",
		"data-science",
		"software-development",
		"web-development",
		"mobile-development",
		"game-development",
		"blockchain",
		"cybersecurity",
		"cloud-computing",
		"devops",
		"database",
		"internet-of-things",
		"embedded-systems",
		"robotics",
		"quantum-computing",
	}

	for _, domain := range domains {
		tag := model.Tag{
			UUID: uuid.NewString(),
			Name: domain,
		}
		method.AddTag(tag)
	}

	fmt.Println(domains)
}
