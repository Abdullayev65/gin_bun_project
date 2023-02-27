package utill

import (
	"fmt"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"io"
	"mime/multipart"
	"os"
	"time"
)

var pathForFiles = "/home/abdullayev/Desktop/"

func UploadFiles(fsMap map[string][]*multipart.FileHeader) []models.Attachment {
	attachs := make([]models.Attachment, 0)
	for _, fs := range fsMap {
		for _, f := range fs {
			file, err := f.Open()
			if err != nil {
				return nil
			}
			pathFile := generatePathAndCheckDir(f.Filename)
			err = copyTo(pathFile, file)
			if err != nil {
				return nil
			}
			{ // append attach
				attachs = append(attachs, models.Attachment{
					FileName: f.Filename, Path: pathFile})
			}
		}
	}
	return attachs
}
func UploadFile(OneFile *multipart.FileHeader) *models.Attachment {
	file, err := OneFile.Open()
	if err != nil {
		return nil
	}
	pathFile := generatePathAndCheckDir(OneFile.Filename)
	err = copyTo(pathFile, file)
	if err != nil {
		return nil
	}

	attachment := models.Attachment{
		FileName: OneFile.Filename, Path: pathFile}

	return &attachment
}

func generatePathAndCheckDir(name string) (pathFile string) {
	pathDir, fileName := generatePath(name)
	mkdirIfNotExists(pathDir)
	pathFile = pathDir + fileName
	return
}
func generatePath(name string) (pathDir, fileName string) {
	now := time.Now()
	pathDir = fmt.Sprintf("%d/%s/%d/%d/", now.Year(),
		now.Month().String(), now.Day(), now.Hour())
	fileName = fmt.Sprintf("m%ds%d__", now.Minute(), now.Second()) + name
	pathDir = pathForFiles + pathDir
	return
}

func mkdirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err1 := os.MkdirAll(path, os.ModePerm)
		if err1 != nil {
			return //my500
		}
	}
}

func copyTo(newFilePath string, src io.Reader) error {
	out, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	return nil
}
