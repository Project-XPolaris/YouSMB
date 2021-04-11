package smb

import (
	"bufio"
	"os"
	"strings"
)

type Section struct {
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields"`
}

type Config struct {
	Sections []*Section `json:"sections"`
}

func ReadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	sections := make([]*Section, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := strings.Replace(line, "[", "", 1)
			sectionName = strings.Replace(sectionName, "]", "", 1)
			sections = append(sections, &Section{
				Name:   sectionName,
				Fields: map[string]string{},
			})
			continue
		}
		group := strings.Split(line, "=")
		if len(group) == 2 {
			sections[len(sections)-1].Fields[strings.TrimSpace(group[0])] = strings.TrimSpace(group[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Config{
		Sections: sections,
	}, nil
}
