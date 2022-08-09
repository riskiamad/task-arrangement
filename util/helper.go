package util

import (
	"fmt"
	"strconv"
	"strings"
	"task-scheduler/config"
	model "task-scheduler/datamodel"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func GenerateCode(format string, tableName string) (code string, e error) {
	var initialCode, template string
	var lenAbbr, lenInitialCode int

	codeLen := 4
	template = format + "#" + fmt.Sprintf("%0"+strconv.Itoa(codeLen)+"d", 1)
	lenAbbr = len(format)
	if e == nil {
		if template != "" {
			if e = config.DbConn.Raw("SELECT cg.value FROM code_generators cg WHERE cg.table = ? AND cg.value LIKE ?", tableName, format+"%").Scan(&initialCode).Error; e == nil {
				lenInitialCode = len(initialCode)
				tempIncrement := initialCode[lenAbbr:lenInitialCode]
				increment, _ := strconv.Atoi(tempIncrement)
				increments := fmt.Sprintf("%0"+strconv.Itoa(codeLen)+"d", increment+1)
				code = fmt.Sprintf("%s%s", format, increments)
			} else {
				code = fmt.Sprintf("%s%s", format, fmt.Sprintf("%0"+strconv.Itoa(codeLen)+"d", 1))
			}
		}
	}

	// simpan code dokumen ke table code_generator untuk cek duplikat atau tidak
	e = config.DbConn.Create(&model.CodeGenerator{Table: tableName, Value: code}).Error
	if e != nil {
		var isDuplicate = strings.Contains(e.Error(), "Duplicate entry")
		if isDuplicate {
			code, e = GenerateCode(format, tableName)
		}

	}

	return code, e
}
