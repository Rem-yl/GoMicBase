package internal

import (
	conf "Account/Conf"
	"fmt"

	"github.com/hashicorp/consul/api"
)

func GetConsulClient() (client *api.Client, err error) {
	defaultConf := api.DefaultConfig()
	consulConf := conf.AccountConf.ConsulConf
	defaultConf.Address = fmt.Sprintf("%s:%d", consulConf.Host, consulConf.Port)

	client, err = api.NewClient(defaultConf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConsulRegWeb(host string, port int, name, id string, tags []string) error {
	client, err := GetConsulClient()
	if err != nil {
		return err
	}

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

	err = client.Agent().ServiceRegister(registration)
	return err
}

func ConsulRegGrpc(host string, port int, name, id string, tags []string) error {
	client, err := GetConsulClient()
	if err != nil {
		return err
	}

	// register accountServ to consul
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
	err = client.Agent().ServiceRegister(registration)
	return err
}

func GetConsulServiceList() (serviceList map[string]*api.AgentService, err error) {
	client, err := GetConsulClient()
	if err != nil {
		return nil, err
	}

	serviceList, err = client.Agent().Services()
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}

func GetFilterConsulService(filter string) (serviceList map[string]*api.AgentService, err error) {
	client, err := GetConsulClient()
	if err != nil {
		return nil, err
	}

	serviceList, err = client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}
