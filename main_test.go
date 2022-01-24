package main

import (
	"backend_solita/CSV_Reader"
	"testing"
)

func Test_Checl_pH(t *testing.T) {
	pH := CSV_Reader.Check_pH(5.5)
	if pH != 5.5 {
		t.Error("ph value is wrong")
	}

	wrong_pH := CSV_Reader.Check_pH(20)
	if wrong_pH != -1 {
		t.Error("ph value is true")
	}
}

func Test_Temprature(t *testing.T) {
	temperature := CSV_Reader.Check_Temperature(-50)
	if temperature != -50 {
		t.Error("temperature value is wrong")
	}

	wrong_temperature := CSV_Reader.Check_Temperature(140)
	if wrong_temperature != -1 {
		t.Error("temperature value is true")
	}
}

func Test_Rainfall(t *testing.T) {
	rainfall := CSV_Reader.Check_Rainfall(0)
	if rainfall != 0 {
		t.Error("rainfall value is wrong")
	}

	wrong_rainfall := CSV_Reader.Check_Rainfall(600)
	if wrong_rainfall != -1 {
		t.Error("rainfall value is true")
	}
}

func Test_CSV_Reader(t *testing.T) {
	var filePath string = "../dev-academy-2022-exercise/PartialTech.csv"
	listOfFarm := CSV_Reader.ReadCsvFile(filePath)
	if len(listOfFarm) == 0 && listOfFarm == nil {
		t.Error("CSV Reader returns empty data")
	}
}
