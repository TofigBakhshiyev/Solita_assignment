package main

import (
	"backend_solita/CSV_Reader"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func dbConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=test")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")

	}

	return db
}

// Migrating csv data to PostgreSQL Database
func csvMigrateToPostgreSql() *gorm.DB {
	var filePath [4]string
	filePath[0] = "./csv_files/PartialTech.csv"
	filePath[1] = "./csv_files/friman_metsola.csv"
	filePath[2] = "./csv_files/Nooras_farm.csv"
	filePath[3] = "./csv_files/ossi_farm.csv"

	var listFarm [][]CSV_Reader.Farm

	for file := 0; file < len(filePath); file++ {
		listFarm = append(listFarm, CSV_Reader.ReadCsvFile(filePath[file]))
	}

	db := dbConnection()

	db.AutoMigrate(&CSV_Reader.Farm{})

	for i := range listFarm {
		for j := range listFarm[i] {
			db.Create(&listFarm[i][j])
		}
	}

	fmt.Println("Data successfully migrated")
	return db
}

func getAllFarms(db *gorm.DB) []CSV_Reader.Farm {
	var farms []CSV_Reader.Farm
	db.Find(&farms)
	return farms
}

func getMin(db *gorm.DB, input string) []CSV_Reader.Farm {
	var farms []CSV_Reader.Farm
	var where string = "sensor_type = " + "'" + string(input) + "'"
	db.Select([]string{"value", "MIN(sensor_type)"}).Where(where).Group("value").First(&farms)
	return farms
}

func getMax(db *gorm.DB, input string) []CSV_Reader.Farm {
	var farms []CSV_Reader.Farm
	var where string = "sensor_type = " + "'" + string(input) + "'"
	db.Select([]string{"value", "MAX(sensor_type)"}).Where(where).Group("value").First(&farms)
	return farms
}

func main() {
	csvMigrateToPostgreSql()
	db := dbConnection()
	app := fiber.New()

	app.Get("/farms", func(c *fiber.Ctx) error {
		farms := getAllFarms(db)
		json, _ := json.Marshal(farms)

		return c.SendString(string(json))
	})

	app.Get("farms/min/:sensortype", func(c *fiber.Ctx) error {
		farms := getMin(db, c.Params("sensortype"))
		json, _ := json.Marshal(farms)
		return c.SendString(string(json))
	})

	app.Get("farms/max/:sensortype", func(c *fiber.Ctx) error {
		farms := getMax(db, c.Params("sensortype"))
		json, _ := json.Marshal(farms)
		return c.SendString(string(json))
	})

	log.Fatal(app.Listen("localhost:3000"))

	defer db.Close()
}
