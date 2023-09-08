package content

import (
	"fmt"
	_ "html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

func CreateProjectStructure(config domain.ProjectConfig) error {
	if err := os.Mkdir(config.ProjectName, 0755); err != nil {
		return err
	}

	subDirs := []string{
		"app",
		"config/database",
		"docs",
		"domain",
		fmt.Sprintf("%s", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/delivery", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/repository", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/usecase", strings.ToLower(config.StructName)),
	}

	for _, dir := range subDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.ProjectName, dir), 0755); err != nil {
			return err
		}
	}

	return nil
}
