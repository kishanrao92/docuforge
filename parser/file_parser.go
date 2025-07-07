package parser

import (
	"os"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	Path string
	Type string
}

type InfraDocument struct {
	Path    string
	Type    string
	Content string
}

var SupportedFileTypes = map[string]string{

	".yaml": "yaml",
	".yml":  "yml",
	".hcl":  "hcl",
	".tf":   "tf",
}

func GetConfigFiles(dirPath string) ([]ConfigFile, error) {

	var configFiles []ConfigFile

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// skip directory
		if info.IsDir() {
			return nil
		}

		// Check if file has the correct extension
		for ext, fileType := range SupportedFileTypes {
			if strings.HasSuffix(info.Name(), ext) {
				configFiles = append(configFiles, ConfigFile{
					Path: path,
					Type: fileType,
				})
				break
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return configFiles, nil

}

func LoadInfraDocuments(configFiles []ConfigFile) ([]InfraDocument, error) {
	var infraDocs []InfraDocument

	for _, file := range configFiles {
		content, err := os.ReadFile(file.Path)
		if err != nil {
			return nil, err
		}

		infraDocs = append(infraDocs, InfraDocument{
			Path:    file.Path,
			Type:    file.Type,
			Content: string(content),
		})
	}
	return infraDocs, nil
}
