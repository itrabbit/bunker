package v1

import "github.com/gofiber/fiber"

func Init(router fiber.Router) {
	router.Get("/files/:namespace/:alias", getFileContent)
	router.Post("/files", uploadFile)
}
