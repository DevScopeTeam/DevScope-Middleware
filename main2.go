package main

import (
	"DevScope-Middleware/method"
	"fmt"
)

func main() {
	// fmt.Println(method.CalculateWorkWeightInRepo("stulzq", "dotnetcore", "CanalSharp"))
	fmt.Println(method.CalculateDeveloperScore("zjy2414"))

	// fmt.Println(method.GetUserEvents("morri3"))
	fmt.Println("***********")
	// repos, err := method.GetUserRecentRepos("rsonghuster")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, repo := range repos {
	// 	fmt.Println(repo.Owner.Login, "|", repo.Name)
	// }
}
