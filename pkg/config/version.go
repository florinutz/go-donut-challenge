package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// ExtendVersionTemplate extends the versionTemplate of a cobra command to add commit and build time
func (c *Config) ExtendVersionTemplate(cmd *cobra.Command, commit string, buildTime string) {
	current := strings.TrimSuffix(cmd.VersionTemplate(), "\n")

	extended := fmt.Sprintf(`%s
commit: %s
build time: %s
`, current, commit, buildTime)

	cmd.SetVersionTemplate(extended)
}
