package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetTing()(*Setting,error)  {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	e := vp.ReadInConfig()
	if e != nil{
		return  nil,e
	}
	return  &Setting{vp},nil
}
