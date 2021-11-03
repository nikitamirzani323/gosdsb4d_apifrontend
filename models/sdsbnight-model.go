package models

import (
	"context"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/config"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_sdsbnightHome() (helpers.Response, error) {
	var obj entities.Model_sdsbnight
	var arraobj []entities.Model_sdsbnight
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			date_sdsb4dnight, 
			prize1_sdsb4dnight , prize2_sdsb4dnight, prize3_sdsb4dnight
			FROM ` + config.DB_tbl_trx_sdsb4d_night + ` 
			ORDER BY date_sdsb4dnight DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			date_sdsb4dnight_db, prize1_sdsb4dnight_db, prize2_sdsb4dnight_db, prize3_sdsb4dnight_db string
		)

		err = row.Scan(&date_sdsb4dnight_db, &prize1_sdsb4dnight_db, &prize2_sdsb4dnight_db, &prize3_sdsb4dnight_db)

		helpers.ErrorCheck(err)
		obj.Sdsbnight_date = date_sdsb4dnight_db
		obj.Sdsbnight_prize1 = prize1_sdsb4dnight_db
		obj.Sdsbnight_prize2 = prize2_sdsb4dnight_db
		obj.Sdsbnight_prize3 = prize3_sdsb4dnight_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
