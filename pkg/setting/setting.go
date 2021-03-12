package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	v := viper.New()
	v.SetConfigName("config")
	for _,config := range configs{
		if config != ""{
			v.AddConfigPath(config)
		}
	}

	v.AddConfigPath("configs/")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp: v}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting)WatchSettingChange()  {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}

