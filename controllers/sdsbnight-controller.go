package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/helpers"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/models"
	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
)

func SdsbnightAll(c *fiber.Ctx) error {
	field_redis := "SDSB4D_LISTSDSBNIGHT_API"

	var obj entities.Model_sdsbnight
	var arraobj []entities.Model_sdsbnight
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		sdsbnight_date, _ := jsonparser.GetString(value, "sdsbnight_date")
		sdsbnight_prize1, _ := jsonparser.GetString(value, "sdsbnight_prize1")
		sdsbnight_prize2, _ := jsonparser.GetString(value, "sdsbnight_prize2")
		sdsbnight_prize3, _ := jsonparser.GetString(value, "sdsbnight_prize3")

		obj.Sdsbnight_date = sdsbnight_date
		obj.Sdsbnight_prize1 = sdsbnight_prize1
		obj.Sdsbnight_prize2 = sdsbnight_prize2
		obj.Sdsbnight_prize3 = sdsbnight_prize3
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_sdsbnightHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 0)
		log.Println("SDSBDAY MYSQL")
		return c.JSON(result)
	} else {
		log.Println("SDSBDAY CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
