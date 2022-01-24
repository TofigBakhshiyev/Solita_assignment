package main

import (
	"backend_solita/CSV_Reader"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var filePath [4]string
	filePath[0] = "./csv_files/PartialTech.csv"
	filePath[1] = "./csv_files/friman_metsola.csv"
	filePath[2] = "./csv_files/Nooras_farm.csv"
	filePath[3] = "./csv_files/ossi_farm.csv"

	var listFarm [][]string

	for file := 0; file < len(filePath); file++ {
		listFarm = append(listFarm, CSV_Reader.ReadCsvFile(filePath[file]))
	}

	b, err := json.Marshal(listFarm)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	app := fiber.New()

	app.Get("/farms", func(c *fiber.Ctx) error {
		return c.SendString(string(b))
	})

	log.Fatal(app.Listen("localhost:3000"))
}
