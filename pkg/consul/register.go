package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"time"
	"xl-common/pkg/conf"
	"xl-common/pkg/logger"
	"xl-common/pkg/util"
)

type ConsulService struct {
	IP   		string
	Port 		int
	Tag  		[]string
	Name 		string
	Timeout		string
	Interval	string
	OverTime	string
}

func registerService(ca string, cs *ConsulService) {

	//register consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = ca
	client, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Printf("NewClient error\n%v", err)
		return
	}
	agent := client.Agent()
	interval := time.Duration(10) * time.Second
	deregister := time.Duration(1) * time.Minute

	check := &api.AgentServiceCheck{ // 健康检查
		Interval:                       interval.String(),                                // 健康检查间隔
		GRPC:                           fmt.Sprintf("%v:%v/%v", cs.IP, cs.Port, cs.Name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
		DeregisterCriticalServiceAfter: deregister.String(),                              // 注销时间，相当于过期时间
	}
	if len(cs.Timeout) > 0{
		check.Timeout = cs.Timeout
	}
	if len(cs.Interval) > 0 {
		check.Interval = cs.Interval
	}
	if len(cs.OverTime) > 0 {
		check.DeregisterCriticalServiceAfter = cs.OverTime
	}
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", cs.Name, cs.IP, cs.Port), // 服务节点的名称
		Name:    cs.Name,                                          // 服务名称
		Tags:    cs.Tag,                                           // tag，可以为空
		Port:    cs.Port,                                          // 服务端口
		Address: cs.IP,                                            // 服务 IP
		Check: 	 check,
	}

	logger.Logger.Debug("registing to ", zap.String("address", ca))
	if err := agent.ServiceRegister(reg); err != nil {
		logger.Logger.Panic("Service Register error ", zap.Error(err))
		return
	}

}

func RegisterToConsul(config conf.ServiceConf) {

	registerService(config.Consul.Addr, &ConsulService{
		Name: config.App.Name,
		Tag:  []string{config.App.Name},
		IP:   util.IpUtil.GetLocalIp(),
		Port: config.App.Port,
		Timeout: config.Consul.Timeout,
		Interval: config.Consul.Interval,
		OverTime: config.Consul.OverTime,
	})
}