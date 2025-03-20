package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orbulant/stock-screener/internal/repositories"
)

func GetStockData(c *fiber.Ctx) error {
	repositories.FetchStockDataFromAPI()
	return c.JSON(fiber.Map{
		"message": "Stock data fetched",
	})
}
