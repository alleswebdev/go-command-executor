package main

import (
	"github.com/alleswebdev/go-command-executor/internal/command"
	"github.com/alleswebdev/go-command-executor/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

func main() {
	cfg := config.GetAppConfig()
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Static("/", "./web/commander-front/dist")
	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.JSON(cfg.Commands)
	})
	app.Post("/api/exec/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if len(name) == 0 {
			return c.SendStatus(500)
		}

		commandsMap := command.GetCommandsMapFromConfig(cfg)
		if cmd, ok := commandsMap[name]; ok {
			result, err := cmd.Start()
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "cmd.Start").Error())
			}

			return c.JSON(result)
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.Port)))

}
