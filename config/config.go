package config

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Listen string `json:"listen"`
	Token  string `json:"token"`
}
type WechatConfig struct {
	CorpId  string `json:"corpid"`
	Secret  string `json:"secret"`
	AgentId int64  `json:"agentid"`
}
type GlobalConfig struct {
	Debug  bool          `json:"debug"`
	Http   *HttpConfig   `json:"http"`
	Wechat *WechatConfig `json:"wechat"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("load configuration file", cfg, "successfully")
	return nil
}
