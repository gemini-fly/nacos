package utils

import (
	"io"
	"log"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos(NacosHost, NamespaceId, DataId, Group string, NacosPort uint64) io.Reader {
	sc := []constant.ServerConfig{
		{
			IpAddr: NacosHost,
			Port:   NacosPort,
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
	return strings.NewReader(string(content))
}
