package config

type Config struct {
	Port    string `json:"port"`
	Env     string `json:"env"`
	AppName string `json:"app_name"`
}
