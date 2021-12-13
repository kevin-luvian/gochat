package util

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func GetRootPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logrus.Fatal("No caller information")
	}

	return filepath.Join(filepath.Dir(filename), "../..")
}
