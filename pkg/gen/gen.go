package gen

import (
	"os"
	"sync"
	"text/template"

	"github.com/tsingshaner/gin-starter/pkg/logger"
)

type CodeMap = map[string]map[int]string

// 生成自定义 code markdown 文档
func GenerateCodeDocs(docsDir string, modules []CodeMap) {
	if err := createDir(docsDir); err != nil {
		logger.Error("create dir error:", docsDir)
		panic(err)
	}

	md, err := docTemplate()
	if err != nil {
		panic(err)
	}

	logger.Info("generate code docs start")

	wg := &sync.WaitGroup{}
	for _, mod := range modules {
		for name, codes := range mod {
			wg.Add(1)
			go func(name string, codes map[int]string) {
				defer wg.Done()
				logger.Info("generate code docs for module:", name)
				generateDocs(docsDir+"/"+name+".md", codes, md)
			}(name, codes)
		}
	}
	wg.Wait()
}

func generateDocs(
	filePath string,
	codes map[int]string,
	template *template.Template,
) {
	writer := getFileWriter(filePath)
	defer writer.Close()

	if err := template.Execute(writer, codes); err != nil {
		logger.Warn(filePath, "execute template error:", err)
	}
	logger.Info("generate code docs success:", filePath)
}

func docTemplate() (*template.Template, error) {
	return template.New("doc").Parse(
		`## 自定义 code 说明

| code | 说明 |
| ---- | ---- |
{{range $key, $value := .}}| {{$key}} | {{$value}} |
{{end}}`)
}

func getFileWriter(filePath string) *os.File {
	writer, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("open file error:", err)
		panic(err)
	}
	return writer
}

func createDir(dir string) error {
	if !isExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
