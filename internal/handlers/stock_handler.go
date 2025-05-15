package handlers

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/orbulant/stock-screener/internal/repositories"
	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
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

func GetPDFData(c *fiber.Ctx) error {

	pdfPath := "assets/financialstatements.pdf"

	// Open the PDF file
	file, err := os.Open(pdfPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to open PDF file: %v", err),
		})
	}
	defer file.Close()

	// Read the PDF file into memory
	data, err := io.ReadAll(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to read PDF file: %v", err),
		})
	}

	// Extract text from the PDF
	var buf bytes.Buffer
	err = pdfapi.ExtractContent(bytes.NewReader(data), &buf, nil, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to extract text from PDF: %v", err),
		})
	}

	// Return the extracted text
	return c.JSON(fiber.Map{
		"message": "PDF data fetched successfully.",
		"data":    buf.String(),
	})
}
