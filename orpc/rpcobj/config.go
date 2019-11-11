package rpcobj

import (
    "github.com/fluofoxxo/outrun/config"
    "github.com/fluofoxxo/outrun/config/eventconf"
    "github.com/fluofoxxo/outrun/config/infoconf"
)

type Config struct {
}

func (c *Config) ReloadAllConfigs(nothing bool, reply *ConfigReply) error {
    err := config.Parse("config.json")
    if err != nil {
        reply.ConfigSuccess = false
        reply.ConfigError = err.Error()
    } else {
        reply.ConfigSuccess = true
    }

    err = eventconf.Parse(config.CFile.EventConfigFilename)
    if err != nil {
        reply.EventConfigSuccess = false
        reply.EventConfigError = err.Error()
    } else {
        reply.EventConfigSuccess = true
    }

    err = infoconf.Parse(config.CFile.InfoConfigFilename)
    if err != nil {
        reply.InfoConfigSuccess = false
        reply.InfoConfigError = err.Error()
    } else {
        reply.InfoConfigSuccess = true
    }
    return nil
}

func (c *Config) ReloadConfig(nothing bool, reply *ConfigReply) error {
    err := config.Parse("config.json")
    if err != nil {
        reply.ConfigSuccess = false
        reply.ConfigError = err.Error()
    } else {
        reply.ConfigSuccess = true
    }
    return nil
}

func (c *Config) ReloadEventConfig(nothing bool, reply *ConfigReply) error {
    err := eventconf.Parse(config.CFile.EventConfigFilename)
    if err != nil {
        reply.EventConfigSuccess = false
        reply.EventConfigError = err.Error()
    } else {
        reply.EventConfigSuccess = true
    }
    return nil
}

func (c *Config) ReloadInfoConfig(nothing bool, reply *ConfigReply) error {
    err := infoconf.Parse(config.CFile.InfoConfigFilename)
    if err != nil {
        reply.InfoConfigSuccess = false
        reply.InfoConfigError = err.Error()
    } else {
        reply.InfoConfigSuccess = true
    }
    return nil
}

type ConfigReply struct {
    ConfigSuccess      bool
    ConfigError        string
    EventConfigSuccess bool
    EventConfigError   string
    InfoConfigSuccess  bool
    InfoConfigError    string
}
