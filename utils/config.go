package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
	FS struct {
		Path string `json:"path"`
		Root string `json:"root"`
	} `json:"fs"`
	Registry struct {
		Able        bool   `json:"able"`
		Host        string `json:"host"`
		Port        int    `json:"port"`
		User        string `json:"user"`
		Passwd      string `json:"passwd"`
		ServiceName string `json:"service_name"`
	} `json:"registry"`
	// DB struct {
	// 	User     string `json:"user"`
	// 	Password string `json:"password"`
	// 	Name     string `json:"name"`
	// } `json:"db"`
}

var Conf *Config

// NewDefaultConfig 创建一个默认配置实例。
//
// 返回值:
//
//	*Config: 返回一个指向Config结构体的指针，该结构体包含了服务器的默认配置。
func NewDefaultConfig() *Config {
	return &Config{
		Server: struct {
			Port int    `json:"port"` // 服务器监听端口
			Host string `json:"host"` // 服务器绑定地址
		}{
			Port: 80,        // 默认监听端口为8080
			Host: "0.0.0.0", // 默认绑定所有网络接口
		},
		FS: struct {
			Path string `json:"path"` // 文件服务器的根目录
			Root string `json:"root"` // 文件服务器的根域名
		}{
			Path: "/pic",
			Root: "./data",
		},
		Registry: struct {
			Able        bool   `json:"able"`
			Host        string `json:"host"`
			Port        int    `json:"port"`
			User        string `json:"user"`
			Passwd      string `json:"passwd"`
			ServiceName string `json:"service_name"`
		}{
			Able:        false,
			Host:        "127.0.0.1",
			Port:        6379,
			User:        "root",
			Passwd:      "123456",
			ServiceName: "pic-bed/pic-bed",
		},
	}
}

// loadConfigJSON 从指定文件名加载配置文件，如果文件不存在，则创建并使用默认配置
// 参数：
//
//	filename: 配置文件的路径
//
// 返回值：
//
//	*Config: 加载或生成的配置结构体指针
//	error: 如果过程中发生错误，则返回错误信息；否则返回nil
func LoadConfigJSON(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		// 如果文件不存在，创建并使用默认配置
		defaultConfig := NewDefaultConfig()
		Conf = defaultConfig
		SaveConfigJSON(filename)
	} else if err != nil {
		// 读取文件失败，返回错误
		return nil, fmt.Errorf("failed to read file: %w", err)
	} else {
		var config *Config
		// 成功读取文件，解析JSON内容
		err = json.Unmarshal(content, &config)
		if err != nil {
			// 解析JSON失败，返回错误
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
		Conf = config
	}
	return Conf, nil
}

// saveConfigJSON 将给定的配置信息保存到指定的文件中，以JSON格式持久化存储。
// 参数：
//
//	filename: 配置文件的路径
//	config: 需要保存的配置结构体实例
//
// 返回值：
//
//	error: 如果过程中发生错误，则返回错误信息；否则返回nil
func SaveConfigJSON(filename string) error {
	// 序列化配置结构体为JSON字节流
	jsonBytes, jsonErr := json.Marshal(Conf)
	if jsonErr != nil {
		return fmt.Errorf("failed to marshal config: %w", jsonErr)
	}

	// 将JSON字节流写入文件
	err := ioutil.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
