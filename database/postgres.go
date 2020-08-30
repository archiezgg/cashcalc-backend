/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	postgresDB  *gorm.DB
	postgresURL = os.Getenv("DATABASE_URL")
)

func StartupPostgres() {
	dbHost, dbUser, dbPort, dbPw, dbName := services.GetPostgresDBSpecsFromURL(postgresURL)
}