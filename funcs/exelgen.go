package funcs

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"time"
)

func CreateExelFile(s Structures, v reflect.Value, lendata int) interface{} {

	var err error
	sheet := "Report"
	xlsx := excelize.NewFile()
	index, _ := xlsx.NewSheet(sheet)
	err = xlsx.SetRowHeight(sheet, 1, 20)
	if err != nil {
		return err
	}

	headersStyle, _ := xlsx.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      true,
			Italic:    false,
			Underline: "",
			Family:    "Times New Roman",
			Size:      12,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#d7e7ec"},
			Shading: 0,
		},
	})

	firstColStyle, _ := xlsx.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Times New Roman",
			Size:      11,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        false,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#d7e7ec"},
			Shading: 0,
		},
	})

	mainStyle, _ := xlsx.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Times New Roman",
			Size:      11,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			WrapText:        false,
		},
	})

	xlsx.SetCellStyle(sheet, "A2", fmt.Sprintf("%c%d", 'A'+s.NumOfFields()-1, 2), headersStyle)

	for i := 0; i < s.NumOfFields(); i++ {
		fieldName := s.StructFields().Field(i).Name
		xlsx.SetColWidth(sheet, "A", fmt.Sprintf("%c", 'A'+i), 25)
		cellName := fmt.Sprintf("%c%d", 'A'+i, 2) // Преобразуем индекс в символы A, B, C, ...
		xlsx.SetCellValue(sheet, cellName, fieldName)

	}

	for i := 0; i < lendata; i++ {
		row := v.Index(i)
		for j := 0; j < row.NumField(); j++ {
			cellName := fmt.Sprintf("%c%d", 'A'+j, i+3)
			xlsx.SetCellValue(sheet, cellName, row.Field(j).Interface())
			xlsx.SetCellStyle(sheet, cellName, cellName, mainStyle)
		}
		xlsx.SetCellStyle(sheet, "A3", fmt.Sprintf("A%v", i+3), firstColStyle)
	}

	xlsx.SetActiveSheet(index)
	fileName := fmt.Sprint("C", time.Now().UnixNano()) + ".xlsx"

	err = xlsx.SaveAs("files/" + fileName)
	if err != nil {
		fmt.Println("Ошибка при сохранении файла Excel:", err)
		return err
	}

	fmt.Println("Файл Excel успешно сгенерирован")
	return nil
}
