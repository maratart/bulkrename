package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func rename(oldname, newname string) error {
	err := os.Rename(oldname, newname)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	if len(os.Args) != 2 {
		usageStr := "bulkrename: rename files in the current dir." +
			" Filenames are taken from the csv-list.\n" +
			"1 csv-column: old filename\n2 csv-column: new filename\n\nUsage:\n" +
			filepath.Base(os.Args[0]) + "TABLE.CSV\n"
		fmt.Println(usageStr)
		return
	}
	tableCsv := os.Args[1]

	csvfile, err := os.Open(tableCsv)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	allCsv, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range allCsv {
		oldname, newname := item[0], item[1]
		err := rename(item[0], item[1])
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Success rename", oldname, "to", newname)
		}
	}

}
