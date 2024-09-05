package register

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Host string
	Port int32
}

type ConsulRegistery struct {
	Config *ConsulConfig
	client *api.Client
}

func (c *ConsulRegistery) NewClient() error {
	defualtConf := api.DefaultConfig()
	defualtConf.Address = fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port)

	client, err := api.NewClient(defualtConf)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *ConsulRegistery) RegisterGrpcServ(host string, port int, name, id string, tags []string) error {
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

	err := c.client.Agent().ServiceRegister(registration)
	return err
}

func (c *ConsulRegistery) RegisterWeb(host string, port int, name, id string, tags []string) error {
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

	err := c.client.Agent().ServiceRegister(registration)
	return err
}

func (c *ConsulRegistery) Deregister(servId string) error {
	return c.client.Agent().ServiceDeregister(servId)
}
