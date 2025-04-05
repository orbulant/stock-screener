package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orbulant/stock-screener/internal/repositories"
)

func GetStockData(c *fiber.Ctx) error {
	result, err := repositories.FetchStockDataFromAPI()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stock data.",
		})
	}

	return c.JSON(fiber.Map{
		"message": result,
	})
}
