package nacos

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
)

// NacosConfig nacos 通用解析struct
type NacosConfig struct {
	// nacos
	Group string `yaml:"Group"`
	// 后台管理服务对应的dataId
	SystemDataId string `yaml:"SystemDataId"`
	// clent
	Endpoint    string `yaml:"Endpoint"`
	NamespaceId string `yaml:"NamespaceId"`
	// server
	NacosServer struct {
		IpAddr      string `yaml:"IpAddr"`
		ContextPath string `yaml:"ContextPath"`
		Port        int64  `yaml:"Port"`
		Scheme      string `yaml:"Scheme"`
	} `yaml:"NacosServer"`
	//
}

// GetConfigInfoFromNacos 从nacos 获取配置文件
func GetConfigInfoFromNacos(nacosConfigPaht string, dataId string, destConfigPath string) (nacosCfg NacosConfig, err error) {
	// 获取本地配置文件，获取 nacos 地址
	err = YamlToStruct(nacosConfigPaht, &nacosCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 根据本地nacos配置文件信息从 nacos 获取目标配置信息
	err = GetNacosConfig(nacosCfg, dataId, destConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// GetConfigBytesInfoFromNacos 从nacos 获取配置文件信息
func GetConfigBytesInfoFromNacos(nacosConfigPaht string, dataId string) (nacosBytes []byte, err error) {
	var nacosCfg NacosConfig
	// 获取 nacos 配置文件
	err = YamlToStruct(nacosConfigPaht, &nacosCfg)
	if err != nil {
		return
	}
	// client
	clientConfig := constant.ClientConfig{
		Endpoint:    nacosCfg.Endpoint,
		NamespaceId: nacosCfg.NamespaceId,
	}
	// server
	serverConfigs := []constant.ServerConfig{{
		IpAddr:      nacosCfg.NacosServer.IpAddr,
		ContextPath: nacosCfg.NacosServer.ContextPath,
		Port:        uint64(nacosCfg.NacosServer.Port),
		Scheme:      nacosCfg.NacosServer.Scheme,
	}}
	// link
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// get configFile
	content, err := configClient.GetConfig(
		vo.ConfigParam{
			DataId: dataId,         // 配置名称
			Group:  nacosCfg.Group, // 配置分组
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return []byte(content), nil
}

// YamlToStruct   Yaml文件转struct
func YamlToStruct(file string, nacos *NacosConfig) error {
	// 读取文件
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print(err)
		return err
	}
	// 转换成Struct
	err = yaml.Unmarshal(b, nacos)
	if err != nil {
		log.Printf("%v\n", err.Error())
		return err
	}
	return nil
}

// GetNacosConfig 获取 nacos 配置文件信息
func GetNacosConfig(conf NacosConfig, dataId string, configPath string) error {
	// client
	clientConfig := constant.ClientConfig{
		Endpoint:    conf.Endpoint,
		NamespaceId: conf.NamespaceId,
	}
	// server
	serverConfigs := []constant.ServerConfig{{
		IpAddr:      conf.NacosServer.IpAddr,
		ContextPath: conf.NacosServer.ContextPath,
		Port:        uint64(conf.NacosServer.Port),
		Scheme:      conf.NacosServer.Scheme,
	}}
	// link
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		fmt.Println(err)
	}
	// get configFile
	content, err := configClient.GetConfig(
		vo.ConfigParam{
			DataId: dataId,     // 配置名称
			Group:  conf.Group, // 配置分组
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
	// 创建 nacos 配置文件，写入
	file, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()             // 及时关闭file句柄
	write := bufio.NewWriter(file) // 写入文件时，使用带缓存的 *Writer
	write.WriteString(content)     //
	write.Flush()                  // Flush将缓存的文件真正写入到文件中
	//
	return err
}
