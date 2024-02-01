package pkg

import (
	"awesomeProject3/internal/model"
	"sync"
)

func Filter(key model.ModelMainLink, valueSlice []model.SecondModel) []model.SecondModel {
	var (
		filteredModels []model.SecondModel
		wg             sync.WaitGroup
		mu             sync.Mutex
	)

	for _, m := range valueSlice {
		wg.Add(1)
		go func(m model.SecondModel) {
			defer wg.Done()
			if key.MainLink != m.ThirstLink {
				mu.Lock()
				filteredModels = append(filteredModels, m)
				mu.Unlock()
			}
		}(m)
	}

	wg.Wait()
	return filteredModels
}
