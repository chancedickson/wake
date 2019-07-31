package config

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path"

	"github.com/pelletier/go-toml"
)

// Config contains an unmarshalled configuration file's values
type Config struct {
	IPAddress  string `toml:"ip_address"`
	MacAddress string `toml:"mac_address"`
	Port       uint16 `toml:"port"`
}

// PathFor gives you the config file path for a given computer name
func PathFor(name string) (*string, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	path := path.Join(user.HomeDir, ".wake", fmt.Sprintf("%s.toml", name))
	return &path, nil
}

// LoadConfig loads the config for the given computer name
func LoadConfig(name string) (*Config, error) {
	path, err := PathFor(name)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return nil, err
	}
	if !tree.Has("ip_address") {
		return nil, fmt.Errorf("The property 'ip_address' was not found in the config file at %v", *path)
	}
	if !tree.Has("mac_address") {
		return nil, fmt.Errorf("The property 'mac_address' was not found in the config file at %v", *path)
	}
	if !tree.Has("port") {
		tree.Set("port", 9)
	}
	config := Config{}
	tree.Unmarshal(&config)
	return &config, nil
}
