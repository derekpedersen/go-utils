package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Write(key string, data [][]string) {
	csvFile, err := os.Create("./bin/output/" + key + "_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	writer.WriteAll(data)
	writer.Flush()
}
