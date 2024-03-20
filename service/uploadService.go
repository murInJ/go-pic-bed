package service

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateFileIDName(fileName string) string {
	ext := filepath.Ext(fileName)
	id := uuid.New().String()
	idName := fmt.Sprintf("%s.%s", id, ext)
	return idName
	// idName := fmt.Sprintf("%s/%s.%s", u.Conf.FS.Root, id, ext)
	// return fmt.Sprintf("%s/%s", u.Conf.FS.Path, idName)
}
