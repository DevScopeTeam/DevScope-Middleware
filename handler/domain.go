package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"

	"github.com/gofiber/fiber/v2"
)

// GetUsernameListByTagHandler godoc
//
//	@Summary		获取领域下的用户列表
//	@Description	获取领域下的用户列表
//	@Tags			Domain
//	@Accept			json
//	@Produce		json
//	@Param			uuid		query		string	true	"TagnUUID"
//	@Param			page		query		int		false	"页码（默认：1）"
//	@Param			pageSize	query		int		false	"每页数量（默认：10）"
//	@Success		200			{object}	model.UsernameListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/domain/list [get]
func GetUsernameListByTagHandler(c *fiber.Ctx) error {

	tag_uuid := c.Query("uuid")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)

	if tag_uuid == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "uuid is empty",
		})
	}

	uname_list, err := method.GetUsernameListByTagUUID(tag_uuid, page, pageSize)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.UsernameListResp{
		Code: 200,
		List: uname_list,
	})
}