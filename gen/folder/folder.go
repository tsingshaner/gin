package folder

import "os"

func Mkdir() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for folder, subFolders := range folderTreeMap {
		if len(subFolders) == 0 {
			os.MkdirAll(pwd+"/"+folder, 0755)
			continue
		}
		for _, subFolder := range subFolders {
			os.MkdirAll(pwd+"/"+folder+"/"+subFolder, 0755)
		}
	}
}

type folderTree = map[string][]string

var folderTreeMap = folderTree{
	"constant": []string{"code", "errs"},
	"dto":      []string{},
	"internal": []string{"entity", "handler", "repository", "service"},
	"model":    []string{},
	"test":     []string{"handler", "repository"},
}

var initFiles = []string{
	"./constant/constant.go",
	"./constant/constant.yml",
	"./dto/dto.go",
	"./internal/entity/entity.go",
	"./internal/handler/handler.go",
	"./internal/handler/handler.yml",
}
