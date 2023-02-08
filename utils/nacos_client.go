package utils

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos(NamespaceId, DataId, Group string) string {
	sc := []constant.ServerConfig{
		{
			IpAddr: "nacos.gemini.com",
			Port:   58848,
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
		log.Fatal(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		log.Fatal(err)
	}
	return content

}

