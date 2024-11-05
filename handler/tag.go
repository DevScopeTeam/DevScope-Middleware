package handler

import (
	"DevScope-Middleware/method"
	"DevScope-Middleware/model"

	"github.com/gofiber/fiber/v2"
)

// GetTagHandler godoc
//
//	@Summary		获取Tag
//	@Description	获取Tag
//	@Tags			Tag
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"UUID"
//	@Success		200		{object}	model.TagResp
//	@Failure		400		{object}	model.OperationResp
//	@Router			/tag/get [get]
func GetTagHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")
	if uuid == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "uuid is empty",
		})
	}

	// 调用方法
	tag, err := method.GetTag(uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	return c.JSON(model.TagResp{
		Code: 200,
		Tag:  tag,
	})
}

// GetTagListHandler godoc
//
//	@Summary		获取Tag列表
//	@Description	获取Tag列表
//	@Tags			Tag
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.TagListResp
//	@Failure		400	{object}	model.OperationResp
//	@Router			/tag/list [get]
func GetTagListHandler(c *fiber.Ctx) error {
	tags, err := method.GetTagList()
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.TagListResp{
		Code: 200,
		List: tags,
	})
}
