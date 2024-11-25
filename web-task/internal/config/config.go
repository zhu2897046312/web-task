package config

import (
    "fmt"
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
}

type ServerConfig struct {
    Port int
    Mode string
}

type DatabaseConfig struct {
    Driver   string
    Host     string
    Port     int
    Username string
    Password string
    DBName   string
    Charset  string
}

// DSN 返回数据库连接字符串
func (c *DatabaseConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
        c.Username,
        c.Password,
        c.Host,
        c.Port,
        c.DBName,
        c.Charset,
    )
}

type RedisConfig struct {
    Host     string
    Port     int
    Password string
    DB       int
}

var GlobalConfig Config

func Init() error {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    
    // 只设置相对于 go.mod 的配置路径
    viper.AddConfigPath("configs")
    
    if err := viper.ReadInConfig(); err != nil {
        return fmt.Errorf("failed to read config file: %v\nSearched paths: %v", err, viper.ConfigFileUsed())
    }

    if err := viper.Unmarshal(&GlobalConfig); err != nil {
        return fmt.Errorf("failed to unmarshal config: %v", err)
    }

    return nil
} 