package CSV_Reader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Farm struct {
	Location   string    `json:"location"`
	Datetime   time.Time `json:"datetime"`
	SensorType string    `json:"sensorType"`
	Value      float64   `json:"value"`
}

// CSV parsing and validation
func ReadCsvFile(filePath string) []string {
	var listFarm []string
	f, _ := os.Open(filePath)

	go func(file io.Reader) {
		records, _ := csv.NewReader(file).ReadAll()
		for _, row := range records {
			if row[0] != "location" && row[1] != "datetime" && row[2] != "sensorType" && row[3] != "value" {
				farm := new(Farm)
				checkFarm := Farm{}

				value, _ := strconv.ParseFloat(row[3], 64)
				if row[2] == "pH" {
					true_pH := Check_pH(value)
					if true_pH != -1 && row[2] != "" {
						farm.Location = row[0]
						farm.Datetime, _ = time.Parse("2006-01-02", row[1][:10])
						farm.SensorType = row[2]
						farm.Value = true_pH
					}
				} else if row[2] == "temprature" {
					true_temprature := Check_Temperature(value)
					if true_temprature != -1 && row[2] != "" {
						farm.Location = row[0]
						farm.Datetime, _ = time.Parse("2006-01-02", row[1][:10])
						farm.SensorType = row[2]
						farm.Value = true_temprature
					}
				} else if row[2] == "rainFall" {
					true_rainfall := Check_Rainfall(value)
					if true_rainfall != -1 && row[2] != "" {
						farm.Location = row[0]
						farm.Datetime, _ = time.Parse("2006-01-02", row[1][:10])
						farm.SensorType = row[2]
						farm.Value = true_rainfall
					}
				}
				if checkFarm != (*farm) {
					b, err := json.Marshal(farm)
					if err != nil {
						fmt.Printf("Error: %s", err)
						return
					}
					listFarm = append(listFarm, string(b))
				}
			}
		}
	}(f)
	time.Sleep(240 * time.Millisecond)
	return listFarm
}
