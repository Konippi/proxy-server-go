package serverconfig

import (
	"os"

	"github.com/Konippi/proxy-server-go/config"
)

func Init() (*config.Config, error) {
	cfg := &config.Config{
		ProxyHost: os.Getenv("PROXY_HOST"),
		ProxyPort: os.Getenv("PROXY_PORT"),
	}
	return cfg, nil
}
