package config

import "github.com/spf13/viper"


type Config struct {
    Host string
    Port string
    Prefork bool
    Domain string
    HtmlReload bool
}


func InitConfigure() *Config {
    var config Config

    v := viper.New()
    v.SetConfigName("config")
    v.SetConfigType("toml")
    v.AddConfigPath(".")

    v.SetDefault("Host", "127.0.0.1")
    v.SetDefault("Port", "3750")
    v.SetDefault("Domain", "127.0.0.1:3750")
    v.SetDefault("Prefork", true)
    v.SetDefault("HtmlReload", false)

    err := v.ReadInConfig()
    if err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            panic("Config file not found; ignore error if desired")
        } else {
            panic("Config file was found but another error was produced")
        }
    }
    v.Unmarshal(&config)

    config.Domain = config.Host + ":" + config.Port

    return &config
}
