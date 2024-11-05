package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"

	"github.com/gofiber/fiber/v2"
)

// GetUserNationalityHandler godoc
//
//	@Summary		获取开发者国籍
//	@Description	获取开发者国籍
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Github用户名"
//	@Success		200			{object}	model.DevNationalityResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/github/user/nation [get]
func GetUserNationalityHandler(c *fiber.Ctx) error {
	// 获取参数
	username := c.Query("username")

	// 调用方法
	nationality, err := method.GetUserNationality(username)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	return c.JSON(model.DevNationalityResp{
		Code:   200,
		Nation: nationality,
	})
}
