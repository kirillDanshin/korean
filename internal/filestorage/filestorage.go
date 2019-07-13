package filestorage

import (
	"fmt"
	"github.com/ZergsLaw/korean/internal/lib/hash"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	upload           = "upload"
	productAvatarDir = "product"
)

type (
	// Сtx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

func SaveAvatar(ctx Ctx, productID int, avatarFile io.ReadCloser) (fileName string, err error) {
	dirPath := filepath.Join(upload, productAvatarDir, GetPathForId(productID))
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", errors.Wrapf(err, "failed to create dir")
	}

	fileName = hash.RandToken()

	file, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create file")
	}

	bytes, err := ioutil.ReadAll(avatarFile)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read file")
	}
	defer avatarFile.Close()

	if _, err = file.Write(bytes); err != nil {
		return "", errors.Wrapf(err, "failed to write byte in file")
	}

	return fileName, nil
}

func GetProductAvatar(userID int, fileName string) (pathToFile string) {
	return createPath(productAvatarDir, userID, fileName)
}

func createPath(dir string, ID int, fileName string) (pathToFile string) {
	return filepath.Join(upload, dir, GetPathForId(ID), fileName)
}

func GetPathForId(ID int) string {
	defaultDate := fmt.Sprintf("%09d", ID)

	first := defaultDate[0:3]
	second := defaultDate[3:6]
	third := defaultDate[6:9]

	return fmt.Sprintf(upload+"/"+productAvatarDir+"/d%s/d%s/d%s", first, second, third)
}
