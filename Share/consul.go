package share

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

// 获取consul服务
// https://www.cnblogs.com/yanweifeng/p/17517634.html
func GetConsulClient(consulConf ConsulConfig) (client *api.Client, err error) {
	defaultConf := api.DefaultConfig()
	defaultConf.Address = fmt.Sprintf("%s:%d", consulConf.Host, consulConf.Port)

	client, err = api.NewClient(defaultConf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// 注册consul grpc服务
func ConsulRegGrpc(client *api.Client, host string, port int, name, id string, tags []string) error {
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", host, port),
			Interval:                       "1s",
			Timeout:                        "3s",
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err := client.Agent().ServiceRegister(registration)
	return err
}

func ConsulRegWeb(client *api.Client, host string, port int, name, id string, tags []string) error {
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: host,
		Port:    port,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", host, port),
			Interval:                       "1s",
			Timeout:                        "3s",
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err := client.Agent().ServiceRegister(registration)
	return err
}

func GetConsulServiceList(client *api.Client) (serviceList map[string]*api.AgentService, err error) {
	serviceList, err = client.Agent().Services()
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}

func GetFilterConsulService(client *api.Client, filter string) (serviceList map[string]*api.AgentService, err error) {
	serviceList, err = client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}
