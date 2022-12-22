package utils

import (
	"encoding/json"
	"nacos/tools"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos(NamespaceId string, DataId string, Group string) *tools.Config {
	sc := []constant.ServerConfig{
		{
			IpAddr: "nacos.gemini.com",
			Port:   8848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId: NamespaceId,
		TimeoutMs:   5000,
		LogLevel:    "error",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		panic(err)
	}
	Config := &tools.Config{}
	err = json.Unmarshal([]byte(content), &Config)
	if err != nil {
		panic(err)
	}
	return Config
}
