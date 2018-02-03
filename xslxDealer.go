package main

import (
	//"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"xlsx2csv/config"
	"path/filepath"
)

func getNewCsvHandler(new_path string) (*os.File, error)  {
	var _, err = os.Stat(new_path)
	if os.IsExist(err) {
		os.Remove(new_path)
	}
	var csv_file, create_err = os.Create(new_path)
	return  csv_file, create_err
}

func getNewPath(new_name, old_name, sheet_name, dir string) (string) {
	var csv_name = new_name
	if (len(csv_name) == 0) {
		csv_name = old_name + "_" + sheet_name + ".csv"
	}
	var csv_path = csv_name
	var new_csv_dir, _ = filepath.Split(csv_path)
	if new_csv_dir == "" {
		csv_path = dir + "/" + csv_name
	}

	return csv_path
}

func transXlsx2Csv(file_path, new_name string) int {
	var xlsx, err = excelize.OpenFile(file_path)
	if err != nil {
		fmt.Println(err)
		return config.BOOL_FALSE
	}
	sheets_map := xlsx.GetSheetMap()
	var dir, old_name = getFileNameAndPath(file_path)
	for _, sheet_name := range sheets_map {
		rows := xlsx.GetRows(sheet_name)
		var csv_path = getNewPath(new_name, old_name, sheet_name, dir)
		var csv_file, err = getNewCsvHandler(csv_path)
		if err != nil {
			fmt.Println(err)
			return config.BOOL_FALSE
		}
		/*var csv_writer = csv.NewWriter(csv_file)
		err = csv_writer.WriteAll(rows)*/
		for _, row := range rows {
			var row_str = strings.TrimRight(strings.Join(row, ","), ",") + "\r\n"
			var _, err = csv_file.WriteString(row_str)
			if (err != nil) {
				fmt.Println(err)
				csv_file.Close()
				return config.BOOL_FALSE
			}
		}

		if err != nil {
			fmt.Println(err)
			return config.BOOL_FALSE
		}
		fmt.Println(fmt.Sprintf("file %s, trans over", csv_path))
		//csv_writer.Flush()
		csv_file.Close()
	}

	return config.BOOL_TRUE
}

func getFileNameAndPath(file_path string) (string, string) {
	var name = ""
	var xlsx_path = filepath.Dir(file_path)
	var xlsx_name = filepath.Base(file_path)
	var xlsx_name_arr = strings.Split(xlsx_name, ".")
	name = xlsx_name_arr[0]
	return xlsx_path, name
}

func fileCheck(file_path string) (error) {
	var _, err = os.Stat(file_path)
	if err != nil {
		return  err
	}
	var file_name = filepath.Base(file_path)
	var file_name_arr = strings.Split(file_name, ".")
	if len(file_name_arr) != 2 || file_name_arr[0] == "" {
		return errors.New("file name is error!")
	}
	var ext = file_name_arr[1]
	if (ext != "xlsx" && ext != "xls") {
		return errors.New("file ext is error!")
	}
	return nil
}
