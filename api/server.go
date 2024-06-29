package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	app := fiber.New()

}

/*

GET:
	read one file
	download one file
	fetch query of files in bucket

POST:
	upload one file

PUT:

DELETE:
	delete one file
	clear bucket
	delete bucket (if empty)

*/
