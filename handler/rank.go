package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"

	"github.com/gofiber/fiber/v2"
)

// GetUserScoreHandler godoc
//
//	@Summary		获取开发者评分
//	@Description	获取开发者评分
//	@Tags			Score
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	model.DevScoreResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/score/get [get]
func GetUserScoreHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	score, err := method.GetRank(username)
	if err != nil {
		// 调用方法
		score, err = method.CalculateDeveloperScore(username)
		if err != nil {
			return c.JSON(model.OperationResp{
				Code: 400,
				Msg:  err.Error(),
			})
		}
		// 存储至数据库
		method.AddRank(score)
	}

	// 返回结果
	return c.JSON(model.DevScoreResp{
		Code:  200,
		Score: score,
	})
}
