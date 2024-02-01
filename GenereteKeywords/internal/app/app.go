package app

import (
	"awesomeProject3/internal/config"
	"awesomeProject3/internal/model"
	"awesomeProject3/pkg"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func Run() {
	cfg := config.GetConfig()
	if _, err := os.Stat(cfg.NameOutFileResult); os.IsNotExist(err) {

	} else {
		err := os.Remove(cfg.NameOutFileResult)
		if err != nil {
			fmt.Println("Error del file", err)
			return
		}
	}
	file, err := os.Open(cfg.NameInputDataFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var mainLinkSlice []model.ModelMainLink
	var secondSlice []model.SecondModel
	// Iterate through the records and print them
	for _, record := range records {
		if record[0] != "" {
			mainLinkSlice = append(mainLinkSlice, model.ModelMainLink{MainLink: record[0]})
		}
		if record[2] != "" {
			secondSlice = append(secondSlice, model.SecondModel{
				KeyWord:    record[1],
				ThirstLink: record[2],
			})
		}
	}
	rand.Seed(time.Now().UnixNano())
	for _, mainLink := range mainLinkSlice {
		result := pkg.Filter(mainLink, secondSlice)

		// Выбор 24 случайных элементов
		selectedElements := pkg.GenericKeyLinks(result, cfg.CountKeywords)

		// Вывод выбранных элементов
		err := writeToFile(mainLink, selectedElements, cfg.NameOutFileResult)
		if err != nil {
			log.Panic(err)
		}

	}
}
func writeToFile(mainLink model.ModelMainLink, data []model.SecondModel, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Запись данных в файл
	file.WriteString(fmt.Sprintf("--------------------------------------------------- %s---------------------------------------------------\n", mainLink.MainLink))
	for counter, element := range data {
		line := fmt.Sprintf("%d -  %s, %s\n", counter, element.KeyWord, element.ThirstLink)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
