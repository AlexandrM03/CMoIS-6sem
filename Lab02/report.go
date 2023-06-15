package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func createReport(freq map[rune]int, name string) error {
	f := excelize.NewFile()
	f.NewSheet(name)
	headers := map[string]string{
		"A1": "Symbol",
		"B1": "Frequency",
	}

	for k, v := range headers {
		f.SetCellValue(name, k, v)
	}

	row := 2
	for k, v := range freq {
		f.SetCellValue(name, "A"+fmt.Sprint(row), string(k))
		f.SetCellValue(name, "B"+fmt.Sprint(row), v)
		row++
	}

	if err := f.AddChart(name, "C1", &excelize.Chart{
		Type: "col",
		Series: []excelize.ChartSeries{
			{
				Name:       name + "!$B$1",
				Categories: name + "!$A$2:$A$" + fmt.Sprint(row),
				Values:     name + "!$B$2:$B$" + fmt.Sprint(row),
			},
		},
		Title: excelize.ChartTitle{
			Name: "Frequency",
		},
		XAxis: excelize.ChartAxis{
			Font: excelize.Font{
				Family: "Calibri",
				Color:  "#000000",
			},
		},
		YAxis: excelize.ChartAxis{
			Font: excelize.Font{
				Family: "Calibri",
				Color:  "#000000",
			},
		},
	}); err != nil {
		return err
	}

	f.SetColWidth(name, "A", "B", 20)

	if err := f.SaveAs("reports.xlsx"); err != nil {
		return err
	}

	return nil
}
