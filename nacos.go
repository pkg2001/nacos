package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func LinkNacos(Namespeaceid, DataId, Group string) (string, error) {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         Namespeaceid, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// Create config client for dynamic configuration
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
		return "", nil
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group})

	if err != nil {
		panic(err)
	}

	return content, err
}
