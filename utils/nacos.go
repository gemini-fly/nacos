package utils

import (
        "dpf/nacos/models"
        "encoding/json"
        "github.com/nacos-group/nacos-sdk-go/clients"
        "github.com/nacos-group/nacos-sdk-go/common/constant"
        "github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos() *models.AKConfig {
        sc := []constant.ServerConfig{
                {
                        IpAddr: "nacos.gemini.com",
                        Port:   8848,
                },
        }
        cc := constant.ClientConfig{
                NamespaceId:         "cb93ebba-4410-4828-a95a-9e5d90927c05",
                TimeoutMs:           5000,
                LogLevel:            "error",
        }
        configClient, err := clients.CreateConfigClient(map[string]interface{}{
                "serverConfigs": sc,
                "clientConfig":  cc,
        })
        if err != nil {
                panic(err)
        }
        content, err := configClient.GetConfig(vo.ConfigParam{
                DataId: "aks",
                Group:  "ops",
        })
        if err != nil {
                panic(err)
        }
        Config := &models.AKConfig{}
        err = json.Unmarshal([]byte(content), &Config)
        if err != nil {
                panic(err)
        }
        return Config
}