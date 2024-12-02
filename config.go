package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ConfigurationManager defines the interface for configuration management.
type ConfigurationManager interface {
	Exists(key string) bool
	SetValue(key string, value any) bool
	GetValue(key string) (any, bool)
	Delete(key string) bool
	Clear()
	ToJSON() (string, error)
	FromJSON(value string) error
	SetFileName(name string)
	Save() error
	Load() error
}

// Configuration implements ConfigurationManager interface.
type Configuration struct {
	data     map[string]any
	fileName string // Field to store the associated file name
}

// NewConfiguration creates a new configuration instance.
func NewConfiguration() *Configuration {
	return &Configuration{data: make(map[string]any)}
}

// SetFileName sets the file name for the configuration.
func (conf *Configuration) SetFileName(name string) {
	conf.fileName = name
}

// Exists checks if a key exists in the configuration.
func (conf *Configuration) Exists(key string) bool {
	_, exists := conf.data[key]
	return exists
}

// SetValue assigns a value to a key if not already set.
func (conf *Configuration) SetValue(key string, value any) bool {
	if !conf.Exists(key) {
		conf.data[key] = value
		return true
	}
	return false
}

// GetValue retrieves the value associated with a key.
func (conf *Configuration) GetValue(key string) (any, bool) {
	value, exists := conf.data[key]
	return value, exists
}

// Delete removes a key from the configuration.
func (conf *Configuration) Delete(key string) bool {
	if conf.Exists(key) {
		delete(conf.data, key)
		return true
	}
	return false
}

// Clear removes all keys from the configuration.
func (conf *Configuration) Clear() {
	conf.data = make(map[string]any)
}

// ToJSON converts the configuration to a JSON string.
func (conf *Configuration) ToJSON() (string, error) {
	jsonData, err := json.Marshal(conf.data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// FromJSON loads configuration from a JSON string.
func (conf *Configuration) FromJSON(value string) error {
	newData := make(map[string]any)
	if err := json.Unmarshal([]byte(value), &newData); err != nil {
		return err
	}
	conf.data = newData
	return nil
}

// Save writes the configuration to the file set by SetFileName.
func (conf *Configuration) Save() error {
	if conf.fileName == "" {
		return fmt.Errorf("file name not set")
	}
	jsonData, err := conf.ToJSON()
	if err != nil {
		return err
	}
	return os.WriteFile(conf.fileName, []byte(jsonData), 0644)
}

// Load reads the configuration from the file set by SetFileName.
func (conf *Configuration) Load() error {
	if conf.fileName == "" {
		return fmt.Errorf("file name not set")
	}
	fileData, err := os.ReadFile(conf.fileName)
	if err != nil {
		return err
	}
	return conf.FromJSON(string(fileData))
}
