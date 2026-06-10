package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ApiKey     string
	BaseUrl    string
	ApiTimeout int
	Transport  string
	Host       string
	Port       int
	Name       string
	Version    string
}

func Load() *Config {
	cfg := Config{
		ApiKey:     getEnv("VMRAY_API_KEY", ""),
		BaseUrl:    getEnv("VMRAY_BASE_URL", ""),
		ApiTimeout: getEnvInt("VMRAY_API_TIMEOUT", 60),
		Transport:  getEnv("VMRAY_TRANSPORT", "http"),
		Host:       getEnv("VMRAY_HOST", "localhost"),
		Port:       getEnvInt("VMRAY_PORT", 8080),
		Name:       getEnv("VMRAY_MCP_NAME", "vmray-mcp-server"),
		Version:    getEnv("VMRAY_MCP_VERSION", "0.0.1"),
	}

	if cfg.ApiKey == "" || cfg.BaseUrl == "" {
		fmt.Println("API Key or Base URL is not set. Please set the VMRAY_API_KEY and VMRAY_BASE_URL environment variables.")
		os.Exit(1)
	}

	return &cfg
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Invalid value for %s: %s, using default %d", key, valueStr, fallback)
		return fallback
	}
	return value
}
