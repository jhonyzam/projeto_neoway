package models

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func GetBase() []DataStoreInsert {
	file, err := os.Open("base/base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dataStores := make([]DataStoreInsert, 0)
	i := 0
	for scanner.Scan() {
		i++
		if i == 1 {
			continue
		}
		dataStores = append(dataStores, converLineToDataStore(scanner.Text()))
	}

	return dataStores
}

func converLineToDataStore(scannerText string) DataStoreInsert {
	r := regexp.MustCompile("[^\\s]+")
	splitLine := r.FindAllString(scannerText, -1)

	return DataStoreInsert{
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
