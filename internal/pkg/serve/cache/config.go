package cache

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
	"github.com/deeptest-com/deeptest-next/internal/pkg/serve/viper_server"
	"github.com/spf13/viper"
	"strconv"
)

var CONFIG = Redis{
	DB:       0,
	Addr:     "127.0.0.1:6379",
	Password: "",
	PoolSize: 0,
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	PoolSize int    `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"`
}

// IsExist config file is exist
func IsExist() bool {
	return getViperConfig().IsFileExist()
}

// Remove remove config file
func Remove() error {
	return getViperConfig().Remove()
}

// Recover
func Recover() error {
	b, err := json.Marshal(CONFIG)
	if err != nil {
		return err
	}
	return getViperConfig().Recover(b)
}

// getViperConfig get viper config
func getViperConfig() viper_server.ViperConfig {
	configName := "redis"
	return viper_server.ViperConfig{
		Debug:     true,
		Directory: consts.ConfDir,
		Name:      configName,
		Type:      consts.ConfigType,
		Watch: func(vi *viper.Viper) error {
			if err := vi.Unmarshal(&CONFIG); err != nil {
				return fmt.Errorf("get Unarshal error: %v", err)
			}
			// watch config file change
			vi.SetConfigName(configName)
			return nil
		},
		Default: []byte(`
{
	"db": ` + strconv.FormatInt(int64(CONFIG.DB), 10) + `,
	"addr": "` + CONFIG.Addr + `",
	"password": "` + CONFIG.Password + `",
	"pool-size": ` + strconv.FormatInt(int64(CONFIG.PoolSize), 10) + `
}`),
	}
}
