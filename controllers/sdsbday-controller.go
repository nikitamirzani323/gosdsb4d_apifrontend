package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gosdsb4d_apifrontend/entities"
	"github.com/nikitamirzani323/gosdsb4d_apifrontend/helpers"
	"github.com/nikitamirzani323/gosdsb4d_apifrontend/models"
)

func SdsbdayAll(c *fiber.Ctx) error {
	field_redis := "SDSB4D_LISTSDSBDAY_API"

	var obj entities.Model_sdsbday
	var arraobj []entities.Model_sdsbday
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		sdsbday_date, _ := jsonparser.GetString(value, "sdsbday_date")
		sdsbday_prize1, _ := jsonparser.GetString(value, "sdsbday_prize1")
		sdsbday_prize2, _ := jsonparser.GetString(value, "sdsbday_prize2")
		sdsbday_prize3, _ := jsonparser.GetString(value, "sdsbday_prize3")

		obj.Sdsbday_date = sdsbday_date
		obj.Sdsbday_prize1 = sdsbday_prize1
		obj.Sdsbday_prize2 = sdsbday_prize2
		obj.Sdsbday_prize3 = sdsbday_prize3
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_sdsbdayHome()
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
