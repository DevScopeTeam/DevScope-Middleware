package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"
	github_model "DevScope-Middleware/model/github"

	"github.com/gofiber/fiber/v2"
)

// GetDeveloperInfoHandler godoc
//
//	@Summary		获取开发者基本信息
//	@Description	获取开发者基本信息
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	github_model.UserResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/github/user/info [get]
func GetDeveloperInfoHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	// 调用方法
	user_info, err := method.GetGithuUserInfo(username)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	return c.JSON(github_model.UserResp{
		Code:     200,
		UserInfo: user_info,
	})
}

// GetUserEventsHandler godoc
//
//	@Summary		获取开发者活动数据
//	@Description	获取开发者活动数据
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	github_model.EventListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/github/user/events [get]
func GetUserEventsHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	// 调用方法
	events, err := method.GetUserEvents(username)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	return c.JSON(github_model.EventListResp{
		Code: 200,
		List: events,
	})
}

// GetUserReposHandler godoc
//
//	@Summary		获取用户的仓库列表
//	@Description	获取用户的仓库列表
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	github_model.RepoListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/github/user/repos [get]
func GetUserReposHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	// 调用方法
	repos, err := method.GetUserRepos(username)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	return c.JSON(github_model.RepoListResp{
		Code: 200,
		List: repos,
	})
}

// GetUserRepoHandler godoc
//
//	@Summary		获取用户的仓库信息
//	@Description	获取用户的仓库信息
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"仓库拥有者"
//	@Param			repo	query		string	true	"仓库名"
//	@Success		200		{object}	github_model.RepoResp
//	@Failure		400		{object}	model.OperationResp
//	@Router			/github/user/repo [get]
func GetUserRepoHandler(c *fiber.Ctx) error {
	owner := c.Query("owner")
	repo := c.Query("repo")
	if owner == "" || repo == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "owner or repo is empty",
		})
	}

	repo_info, err := method.GetRepo(owner, repo)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(github_model.RepoResp{
		Code: 200,
		Repo: repo_info,
	})
}
