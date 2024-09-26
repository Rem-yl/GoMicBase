package registry

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulRegistery struct {
	Host   string
	Port   int32
	client *api.Client
}

func NewConsulRegistery(host string, port int32) (*ConsulRegistery, error) {
	c := &ConsulRegistery{
		Host: host,
		Port: port,
	}

	err := c.NewClient()

	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ConsulRegistery) NewClient() error {
	defualtConf := api.DefaultConfig()
	defualtConf.Address = fmt.Sprintf("%s:%d", c.Host, c.Port)

	client, err := api.NewClient(defualtConf)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *ConsulRegistery) RegisterGrpcServ(host string, port int, name, id string) error {
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", "host.docker.internal", port), // consul使用docker部署, 在docker内 host.docker.internal 映射到宿主机的localhost, 避免网络问题
			Interval:                       "1s",
			Timeout:                        "3s",
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err := c.client.Agent().ServiceRegister(registration)
	return err
}

func (c *ConsulRegistery) RegisterWeb(host string, port int, name, id string) error {
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", "host.docker.internal", port), // consul使用docker部署, 在docker内 host.docker.internal 映射到宿主机的localhost, 避免网络问题
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
