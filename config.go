package commander

import (
	"path"

	"github.com/clonercl/common/pkg/log"
)

func LoadViperConfig(confPath string) (err error) {
	x := path.Base(confPath)
	log.Info(x)

	// file, err := ioutil.ReadFile(confPath)
	// if err != nil {
	// 	return err
	// }
	// switch path.Ext(confPath) {
	// case ".json":
	// case ".yaml":
	// default:
	// 	err = fmt.Errorf("config type not supported")
	// 	log.Error(err)
	// 	return err
	// }

	// return nil
	return nil
}
