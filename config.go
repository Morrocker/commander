package commander

import (
	"path"
	"strings"

	"github.com/clonercl/common/pkg/log"
	"github.com/spf13/viper"
)

func LoadViperConfig(confPath string, confRoot string) (err error) {
	x := path.Base(confPath)
	fileArr := strings.Split(x, ".")
	log.Info(x)
	viper.SetConfigType(fileArr[1])
	viper.SetConfigName(fileArr[0])
	viper.AddConfigPath(".")
	if confRoot != "" {
		viper.AddConfigPath(confRoot)
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
	}
	return nil
}
