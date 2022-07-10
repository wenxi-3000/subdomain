package libs

import (
	"subdomain/utils"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//生成配置文件config.yaml
func InitConfig(opt *Options) error {
	rootPath := opt.Paths.Root
	v := viper.New()
	v.AddConfigPath(rootPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if !utils.FileExists(opt.ConfigFile) {
		v.SetDefault("fofa", map[string][]string{})
		v.SetDefault("censys", map[string][]string{})
		err := v.WriteConfigAs(opt.ConfigFile)
		return err
	}
	GetKeys(opt)
	return nil
}

func GetKeys(opt *Options) {
	v, _ := LoadConfig(*opt)
	opt.Keys = make(map[string][]string)
	for _, souce := range Resources {
		opt.Keys[souce] = v.GetStringSlice(souce)
	}

}

// LoadConfig load config
func LoadConfig(options Options) (*viper.Viper, error) {
	options.ConfigFile, _ = homedir.Expand(options.ConfigFile)
	rootPath := options.Paths.Root
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(rootPath)
	// InitConfig(&options)
	if err := v.ReadInConfig(); err != nil {
		InitConfig(&options)
		return v, nil
	}
	return v, nil
}
