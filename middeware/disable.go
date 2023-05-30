package middeware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Disable(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Path())
	if strings.Index(ctx.Path(), "list") == -1 {
		return ctx.JSON(fiber.Map{
			"code": 0,
			"data": "主体数据，不能变动",
			"msg":  "操作失败",
		})
	} else {
		return ctx.Next()
	}
}
