package main

import (
	"bufio"
	"datastore/models"
	"log"
	"os"
	"regexp"
)

func getBase() []models.DataStoreInsert {
	file, err := os.Open("base/base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dataStores := make([]models.DataStoreInsert, 0)
	i := 0
	for scanner.Scan() {
		i++
		if i == 1 {
			continue
		}
		dataStores = append(dataStores, converLineToDataStore(scanner.Text()))

		if i == 10000 {
			break
		}
	}

	return dataStores
}

func converLineToDataStore(scannerText string) models.DataStoreInsert {
	r := regexp.MustCompile("[^\\s]+")
	splitLine := r.FindAllString(scannerText, -1)

	return models.DataStoreInsert{
		Cpf:           splitLine[0],
		Private:       splitLine[1],
		Incompleto:    splitLine[2],
		LastDate:      splitLine[3],
		AvgTicket:     splitLine[4],
		LastTicket:    splitLine[5],
		StoreFrequent: splitLine[6],
		StoreLast:     splitLine[7],
	}
}
