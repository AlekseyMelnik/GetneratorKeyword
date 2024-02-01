package pkg

import (
	"awesomeProject3/internal/model"
	"math/rand"
)

func GenericKeyLinks(slice []model.SecondModel, count int) []model.SecondModel {
	length := len(slice)
	if count > length {
		count = length
	}

	// Перемешивание слайса
	rand.Shuffle(length, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	// Выбор первых count элементов
	return slice[:count]
}
