package cnf

import (
	"fmt"

	"go.uber.org/config"
)

func LoadConfig(configPath string, target interface{}) error {
	provider, err := config.NewYAMLProviderFromFiles(configPath)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	if err := provider.Get("config").Populate(target); err != nil {
		return fmt.Errorf("provider: %w", err)
	}
	return nil
}
