package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		// 开启调试模式
		Prefork:               false,
		DisableStartupMessage: true,
	})
	file, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 日志记录中间件
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:       "${time} ${method} ${path} - ${status} - ${body} - ${latency}\n",
		TimeFormat:   "2006-01-02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       file,
	}))
	// 注册一个中间件
	app.Use(func(c *fiber.Ctx) error {
		// 输出请求信息
		fmt.Printf("[%s] %s\n", c.Method(), c.Path())

		// 继续执行后续的操作
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	log.Fatal(app.Listen(":3000"))
}
