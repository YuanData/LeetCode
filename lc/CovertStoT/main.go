package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/liuzl/gocc"
)

func convertS2T(text string) (string, error) {
	s2tw, err := gocc.New("s2tw")
	if err != nil {
		return "", err
	}
	return s2tw.Convert(text)
}

func replaceWords(text string, wordDict map[string]string) string {
	for key, pattern := range wordDict {
		re := regexp.MustCompile("(?i)" + pattern)
		text = re.ReplaceAllString(text, key)
	}
	return text
}

func removeImage(text string) string {
	// re := regexp.MustCompile(`<Image alt=".*?" src=".*?"\/>`)
	re := regexp.MustCompile(`(?m)(\n\s*\n)?\s*<Image alt=".*?" src=".*?"\/>\s*\n`)
	return re.ReplaceAllString(text, "")
}

func processFiles(dirpath string) {
	alternatives := map[string]string{
		"宣告":        "聲明",
		"程式":        "代碼",
		"記憶體":       "內存",
		"指標":        "指針",
		"布林":        "布爾",
		"函式":        "函數",
		"套件":        "包",
		"陣列":        "數組",
		"slice":     "切片",
		"介面":        "接口",
		"Goroutine": "協程",
		"channel":   "信道|管道",
		"啟動":        "激活",
		"當機":        "死機",
		"迭代":        "遍歷",
		"Go":        "Golang|Go語言",
		"網路資訊":      "互聯網",
		"遞迴":        "遞歸",
		"窮舉":        "暴力",
		"硬體":        "硬件",
		"軟體":        "軟件",
	}

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".md") {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(data)
			converted, err := convertS2T(content)
			if err != nil {
				return err
			}
			content = replaceWords(converted, alternatives)
			content = removeImage(content)

			// 處理文件名
			newName, err := convertS2T(info.Name())
			if err != nil {
				return err
			}
			newName = replaceWords(newName, alternatives)
			newPath := filepath.Join("done", newName)

			err = ioutil.WriteFile(newPath, []byte(content), 0644)
			if err != nil {
				return err
			}

			err = os.Remove(path)
			if err != nil {
				return err
			}

			fmt.Printf("Processed and moved %s to %s\n", path, newPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing files: %s\n", err)
	}
}

func main() {
	dirpath := "doing" // 替换为你的目录路径
	processFiles(dirpath)
}
