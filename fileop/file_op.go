package fileop

import (
	"bufio"
	"encoding/json"
	"os"
)

type IFileOP interface {
	Read(filepath string) (IFileOP, error)
	GetLines() (lines []string)
	ParsedLotto649CSV(split_space string) (results []Lotto649CSV, err error)
	ParsedSuperlotto638(split_space string) (results []Superlotto638CSV, err error)
}

type FileOP struct {
	lines []string
}

func NewFileOP() IFileOP {
	return &FileOP{}
}

// 讀取
func (fop *FileOP) Read(file_path string) (IFileOP, error) {
	data, err := os.Open(file_path)

	if err != nil {
		return fop, err
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	data.Close()
	fop.lines = fileTextLines

	return fop, nil
}

func (fop *FileOP) GetLines() (lines []string) {
	return fop.lines
}

/*
- 目的: 取得結構資料的 keys
*/
func (fop *FileOP) get_struct_keys(tmp_struct interface{}) (keys []string, err error) {
	b, err := json.Marshal(tmp_struct)
	if err != nil {
		return
	}
	tmp_map := map[string]interface{}{}
	json.Unmarshal(b, &tmp_map)
	keys = make([]string, 0, len(tmp_map))
	for k := range tmp_map {
		keys = append(keys, k)
	}
	return
}
