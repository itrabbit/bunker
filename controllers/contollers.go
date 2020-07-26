package controllers

import (
	"fmt"

	v1 "github.com/itrabbit/bunker/controllers/v1"

	"github.com/gofiber/fiber"
)

var versions = make([]string, 0)

func GetVersions() []string {
	return versions[:]
}

func Init(router fiber.Router, version string) fiber.Router {
	var v fiber.Router
	defer func() {
		if v != nil {
			versions = append(versions, version)
		}
	}()
	switch version {
	case "v1":
		v = router.Group(fmt.Sprint("/", version))
		v1.Init(v)
		break
	default:
		fmt.Println("[WARN] Controllers for version", version, "not found!")
		break
	}
	return v
}
