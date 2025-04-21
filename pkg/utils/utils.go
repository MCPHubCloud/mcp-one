package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 判断文件类型
func GetFileType(filePath string) string {
	lowerCasePath := strings.ToLower(filePath)
	if strings.HasSuffix(lowerCasePath, ".json") {
		return "json"
	} else if strings.HasSuffix(lowerCasePath, ".yaml") || strings.HasSuffix(lowerCasePath, ".yml") {
		return "yaml"
	}
	return ""
}

// 读取文件并解析，使用泛型
func ReadAndParseFile[T any](filePath string) (*T, error) {
	// 检查是否为绝对路径，如果不是则拼接当前工作目录
	if !filepath.IsAbs(filePath) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		filePath = filepath.Join(wd, filePath)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileType := GetFileType(filePath)
	var result T
	switch fileType {
	case "json":
		err = json.Unmarshal(data, &result)
	case "yaml":
		err = yaml.Unmarshal(data, &result)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}
