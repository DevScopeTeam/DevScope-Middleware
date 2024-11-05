package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUserScoreHandler godoc
//
//	@Summary		获取开发者评分
//	@Description	获取开发者评分
//	@Tags			Rank
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	model.DevScoreResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/rank/score [get]
func GetUserScoreHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	score, err := method.GetRank(username)
	if err != nil {
		// 计算分数
		score, err = method.CalculateDeveloperScore(username)
		if err != nil {
			return c.JSON(model.OperationResp{
				Code: 400,
				Msg:  err.Error(),
			})
		}

		// 获取国籍
		score.Nation, err = method.GetUserNation(username)
		if err != nil {
			return c.JSON(model.OperationResp{
				Code: 400,
				Msg:  err.Error(),
			})
		}

		// 分析领域
		domains, err := method.AnalyzeUserDomainList(username)
		if err != nil {
			return c.JSON(model.OperationResp{
				Code: 400,
				Msg:  err.Error(),
			})
		}

		// 存储至数据库
		method.AddRank(score)
		for _, domain_name := range domains {
			domain_uuid, _ := method.GetDomainUUID(domain_name)
			method.AddDomain(model.Domain{
				UUID:     uuid.NewString(),
				Username: username,
				TagUUID:  domain_uuid,
			})
		}
	}

	// 返回结果
	return c.JSON(model.DevScoreResp{
		Code:  200,
		Score: score,
	})
}

// GetUserRankListHandler godoc
//
//	@Summary		获取开发者排名
//	@Description	获取开发者排名
//	@Tags			Rank
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int		false	"页码（默认：1）"
//	@Param			pageSize	query		int		false	"每页数量（默认：10）"
//	@Param			nation		query		string	false	"国籍"
//	@Success		200			{object}	model.RankListResp
//	@Failure		400			{object}	model.OperationResp
//
//	@Router			/rank/list [get]
func GetUserRankListHandler(c *fiber.Ctx) error {
	// 获取参数
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	nation := c.Query("nation")

	ranks, err := method.GetRankList(page, pageSize, nation)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.RankListResp{
		Code: 200,
		List: ranks,
	})
}
