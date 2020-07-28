package v1

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/itrabbit/bunker/config"

	"github.com/itrabbit/bunker/models"

	"github.com/itrabbit/bunker/db"

	"github.com/gofiber/fiber"
)

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func getAbsoluteFilePath(file *models.File, variant string) (string, string) {
	if file == nil || file.NameSpace == nil {
		return "", ""
	}
	path := fmt.Sprint(config.GetStoragePath(), string(os.PathSeparator), strings.Join(file.NameSpace.Parts, string(os.PathSeparator)))
	ext, _ := mime.ExtensionsByType(file.MimeType)
	if len(ext) < 1 {
		ext = []string{"raw"}
	}
	fullFileName := fmt.Sprint(file.Alias, ".", ext[0])
	if len(variant) > 0 && variant != "origin" {
		filePath := fmt.Sprint(path, string(os.PathSeparator), variant, "_", fullFileName)
		if _, err := os.Stat(path); os.IsExist(err) {
			return filePath, ""
		}
	}
	return fmt.Sprint(path, string(os.PathSeparator), fullFileName), variant
}

func getFileContent(c *fiber.Ctx) {
	ns, err := db.NameSpaceMapper.FindOne(strings.TrimSpace(strings.ToLower(c.Params("namespace", ""))))
	if err != nil {
		c.Next(fiber.NewError(400, fmt.Sprint("Error get name space info, ", err.Error())))
		return
	}
	alias := strings.TrimSpace(strings.ToLower(fileNameWithoutExtension(c.Params("alias", ""))))
	if len(alias) < 1 {
		c.Next(fiber.NewError(400, "Empty input file alias"))
		return
	}
	file, err := db.FilesMapper.FindOne(ns.ID, alias)
	if err != nil {
		c.Next(fiber.NewError(400, fmt.Sprint("Error get file info, ", err.Error())))
		return
	}
	file.NameSpace = ns
	path, _ := getAbsoluteFilePath(file, c.Query("variant", ""))
	if len(path) < 1 {
		c.Next(fiber.NewError(500, fmt.Sprint("File data not found")))
		return
	}
	// c.Set("Content-Type", file.MimeType)
	if err := c.SendFile(path, c.Query("compress", "") == "1"); err != nil {
		c.Next(err)
		return
	}
}

func uploadFile(c *fiber.Ctx) {

}
