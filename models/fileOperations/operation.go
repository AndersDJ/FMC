package fileOperations

import (
	"encoding/base64"
	"fmt"

	"github.com/AndersDJ/FMC/pkg/utils"

	log "github.com/cihub/seelog"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type FileInfo struct {
	Key     int    `json:"key" validate:"required"`
	Name    string `json:"name" validate:"required"`
	IsDir   bool   `json:"is_dir" validate:"required"`
	Size    string `json:"size" validate:"required"`
	ModTime string `json:"mod_time" validate:"required"`
}
type FileInfoList struct {
	BasePath string     `json:"base_path" validate:"required"`
	List     []FileInfo `json:"list" validate:"required"`
}

var BP afero.Fs

func Init() {
	defer log.Flush()
	basePath := viper.GetString("control_path")
	BP = afero.NewBasePathFs(afero.NewOsFs(), basePath)
	// BP = afero.NewOsFs()
	s, _ := BP.Stat(".")
	log.Infof("Name: %s", s.Name())
	log.Infof("isDir: %s", s.IsDir())
	log.Infof("Mode: %s", s.Mode().String())
	log.Infof("ModTime: %s", s.ModTime().String())
	log.Infof("Size: %d", s.Size())
	log.Infof("Sys: %v", s.Sys())

	f, err := BP.Open("./../")
	if err != nil {
		log.Error(err.Error())
		return
	}
	stat, _ := f.Stat()
	log.Info(stat.Size())

	dirList, _ := f.Readdir(2000)
	log.Info(len(dirList))
	for _, d := range dirList {
		log.Infof(d.Name())
	}
	// l, _ := f.Readdirnames(1)
	// log.Info(l)
}

func ListDir(d afero.File, basePath string) (*FileInfoList, error) {
	s, err := d.Stat()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if !s.IsDir() {
		err := fmt.Errorf("not dir")
		log.Error("%s is not a dir", d.Name())
		return nil, err
	}
	files, _ := d.Readdir(10000)
	log.Info(len(files))
	var l []FileInfo
	for k, v := range files {
		var f FileInfo
		f.Key = k
		f.Name = v.Name()
		f.IsDir = v.IsDir()
		if v.IsDir() {
			f.Size = "-"
		} else {
			f.Size = utils.CountSize(int(v.Size()))
		}
		f.ModTime = v.ModTime().Format("2006-01-02 15:04:05")
		l = append(l, f)
	}
	encodeBP := base64.StdEncoding.EncodeToString([]byte(basePath))

	return &FileInfoList{
		BasePath: encodeBP,
		List:     l,
	}, nil
}
