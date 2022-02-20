package tools

import "gopkg.in/ini.v1"

type Conn struct {
	Host     string
	Port     int
	UserName string
	Password string
	Schema   string
}

func ParseIni(path string) (con *Conn, err error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	con = &Conn{
		Host:     cfg.Section("connections").Key("host").String(),
		Port:     cfg.Section("connections").Key("port").MustInt(),
		UserName: cfg.Section("connections").Key("username").String(),
		Password: cfg.Section("connections").Key("password").String(),
		Schema:   cfg.Section("connections").Key("schema").String(),
	}
	return
}
