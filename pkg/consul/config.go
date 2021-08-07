package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
	"xl-common/pkg/conf"
	"xl-common/pkg/logger"
)

type Context struct {
	DateSource DataSource `yaml:"dataSource"`
	Redis      Redis      `yaml:"redis"`
	Token      TokenCnf   `yaml:"token"`
	Worker     WorkerCnf  `yaml:"worker"`
}

/**
 * 数据源信息
 */
type DataSource struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Schema   string `yaml:"schema"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type TokenCnf struct {
	Secret    string `yaml:"secret"`
	Issuer    string `yaml:"issuer"`
	Expired   int64  `yaml:"expired"`
	Refresh   int64  `yaml:"refresh"`
	TicketExp int64  `yaml:"ticketExp"`
}

type WorkerCnf struct {
	Url string `yaml:"url"`
}

var (
	KVContext Context
)

func GetConsulKV(conf conf.ServiceConf) {
	config := api.DefaultConfig()
	config.Address = conf.Consul.Addr
	client, err := api.NewClient(config)
	if err != nil {
		logger.Logger.Panic("new Consul client fail", zap.Error(err))
	}
	dev := os.Getenv("DEV")
	if len(dev) == 0 {
		dev = "dev"
	}
	key := fmt.Sprintf("%s-%s.yml", conf.App.Name, dev)
	kv, _, err := client.KV().Get(key, nil)
	if err != nil {
		logger.Logger.Panic("get ConsulKV fail", zap.Error(err))
	}
	if kv == nil {
		logger.Logger.Panic("get KV is null", zap.String("key", key))
	}
	err = yaml.Unmarshal(kv.Value, &KVContext)
	if err != nil {
		logger.Logger.Panic("ConsulKV to yaml fail", zap.Error(err))
	}
	logger.Logger.Info("GetConsulKV success")
}
