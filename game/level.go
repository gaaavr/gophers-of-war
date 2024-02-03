package game

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// load level csv
func loadLevelData(levelName string) ([][]int, error) {
	file, err := os.Open(fmt.Sprintf("levels/%s.csv", levelName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(data))

	for x, row := range data {
		for _, col := range row {
			tileType, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			result[x] = append(result[x], tileType)
		}
	}

	return result, nil
}
