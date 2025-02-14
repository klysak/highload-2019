package musiclibrary_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/glumpo/highload-2019/golang/gozam/models"
	"github.com/glumpo/highload-2019/golang/gozam/musiclibrary"
)

func TestOneTrack(t *testing.T) {
	cfg := models.Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBname:   os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
	}

	musicLib, _ := musiclibrary.Open(cfg)
	defer musicLib.Close()

	originName := "kitay brusnika himky les (forest)"
	sampleName := "forest"

	originFileName := originName + ".mp3"
	sampleFileName := sampleName + ".mp3"

	originPath := filepath.Join("testdata/origin", originFileName)
	samplePath := filepath.Join("testdata/sample", sampleFileName)

	_ = musicLib.Index(originPath)
	result, err := musicLib.Recognize(samplePath)
	if err != nil {
		t.Error(err)
	}

	if result != originName {
		t.Error("Wrong")
	}
}
