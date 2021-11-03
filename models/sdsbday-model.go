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

func Fetch_sdsbdayHome() (helpers.Response, error) {
	var obj entities.Model_sdsbday
	var arraobj []entities.Model_sdsbday
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			datesdsb4dday, 
			prize1_sdsb4dday , prize2_sdsb4dday, prize3_sdsb4dday
			FROM ` + config.DB_tbl_trx_sdsb4d_day + ` 
			ORDER BY datesdsb4dday DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			datesdsb4dday_db, prize1_sdsb4dday_db, prize2_sdsb4dday_db, prize3_sdsb4dday_db string
		)

		err = row.Scan(&datesdsb4dday_db, &prize1_sdsb4dday_db, &prize2_sdsb4dday_db, &prize3_sdsb4dday_db)

		helpers.ErrorCheck(err)
		obj.Sdsbday_date = datesdsb4dday_db
		obj.Sdsbday_prize1 = prize1_sdsb4dday_db
		obj.Sdsbday_prize2 = prize2_sdsb4dday_db
		obj.Sdsbday_prize3 = prize3_sdsb4dday_db
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
