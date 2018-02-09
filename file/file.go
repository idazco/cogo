package file

import (
	"os"
	"path/filepath"
	"log"
	"github.com/spf13/viper"
	"github.com/kardianos/osext"
	"crypto/sha256"
	"io"
	"fmt"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}

func DirOfApp() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return dir
}

func PathOfApp() string {
	exePath, err := osext.Executable()
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return exePath
}

func ParseYaml(path string, targetStruct interface{}) error {
	viper.SetConfigType("yaml")

	ymlFile, err := os.Open(path)
	if err != nil {
		return err
	}

	if err := viper.ReadConfig(ymlFile); err != nil {
		return err
	}

	if err := viper.Unmarshal(targetStruct); err != nil {
		return err
	}

	return nil
}

func Sha256(path string) (string, error) {

	getHash := func() ([]byte, error) {
		var result []byte

		targetFile, err := os.Open(path)
		if err != nil {
			return result, err
		}
		defer targetFile.Close()

		hash := sha256.New()
		if _, err := io.Copy(hash, targetFile); err != nil {
			return result, err
		}
		return hash.Sum(result), nil
	}

	v, e := getHash()
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", v), nil
}
