package config

import (
	"encoding/json"
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
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}

// Configuration implements ConfigurationManager interface.
type Configuration struct {
	data map[string]any
}

// NewConfiguration creates a new configuration instance.
func NewConfiguration() *Configuration {
	return &Configuration{data: make(map[string]any)}
}

// Exists checks if a key exists in the configuration.
func (c *Configuration) Exists(key string) bool {
	_, exists := c.data[key]
	return exists
}

// SetValue assigns a value to a key if not already set.
func (c *Configuration) SetValue(key string, value any) bool {
	if !c.Exists(key) {
		c.data[key] = value
		return true
	}
	return false
}

// GetValue retrieves the value associated with a key.
func (c *Configuration) GetValue(key string) (any, bool) {
	value, exists := c.data[key]
	return value, exists
}

// Delete removes a key from the configuration.
func (c *Configuration) Delete(key string) bool {
	if c.Exists(key) {
		delete(c.data, key)
		return true
	}
	return false
}

// Clear removes all keys from the configuration.
func (c *Configuration) Clear() {
	c.data = make(map[string]any)
}

// ToJSON converts the configuration to a JSON string.
func (c *Configuration) ToJSON() (string, error) {
	jsonData, err := json.Marshal(c.data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// FromJSON loads configuration from a JSON string.
func (c *Configuration) FromJSON(value string) error {
	newData := make(map[string]any)
	if err := json.Unmarshal([]byte(value), &newData); err != nil {
		return err
	}
	c.data = newData
	return nil
}

// SaveToFile saves the configuration to a file.
func (c *Configuration) SaveToFile(filename string) error {
	jsonData, err := c.ToJSON()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(jsonData), 0644)
}

// LoadFromFile loads the configuration from a file.
func (c *Configuration) LoadFromFile(filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return c.FromJSON(string(fileData))
}
