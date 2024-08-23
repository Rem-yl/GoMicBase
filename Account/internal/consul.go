package internal

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func getConsulClient() (client *api.Client, err error) {
	defaultConf := api.DefaultConfig()
	consulConf := AccountConf.ConsulConf
	defaultConf.Address = fmt.Sprintf("%s:%d", consulConf.Host, consulConf.Port)

	client, err = api.NewClient(defaultConf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConsulReg(host string, port int, name, id string) error {
	client, err := getConsulClient()
	if err != nil {
		return nil
	}

	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: host,
		Port:    port,
		Tags:    []string{"test"},
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

func GetConsulServiceList() (serviceList map[string]*api.AgentService, err error) {
	client, err := getConsulClient()
	if err != nil {
		return nil, err
	}

	serviceList, err = client.Agent().Services()
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}

func GetFilterConsulService(name string) (service map[string]*api.AgentService, err error) {
	client, err := getConsulClient()
	if err != nil {
		return nil, err
	}

	service, err = client.Agent().ServicesWithFilter(name)
	if err != nil {
		return nil, err
	}

	return service, nil
}
