package config

import (
	"fmt"
	"os"
	"training/package/log"
)



// Returns the path of report pdf templates
// Copy all pdf templates from doc/pdf-tempates path to .storage/reports/pdf-templates
func PDFReportTemplates() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting working directory %v:", err)
		return ""
	}

	path = fmt.Sprintf("%s/.storage/reports/pdf-templates/", path)
	err = CreateFolderIfDoesntExist(path)
	if err != nil {
		log.Errorf("error getting pdf templates directory: %v", err)
		return ""
	}

	return path
}

func ExcelReportTemplates() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting working directory %v:", err)
		return ""
	}

	path = fmt.Sprintf("%s/.storage/reports/excel-templates/", path)
	err = CreateFolderIfDoesntExist(path)
	if err != nil {
		log.Errorf("error getting excel templates directory: %v", err)
		return ""
	}

	return path
}

func ExcelReportPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting working directory %v:", err)
		return ""
	}

	path = fmt.Sprintf("%s/.storage/reports/excel/", path)
	err = CreateFolderIfDoesntExist(path)
	if err != nil {
		log.Error("error getting excel directory: %v", err)
		return ""
	}

	return path
}

func CreateFolderIfDoesntExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
