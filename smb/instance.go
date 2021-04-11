package smb

import (
	"errors"
	"yousmb/application"
)

var Current *Config
var (
	FolderAlreadyExist = errors.New("exist item")
)

func LoadConfig(configPath string) error {
	config, err := ReadConfig(configPath)
	if err != nil {
		return err
	}
	Current = config
	return nil
}

func (c *Config) SaveFileAndRestart() error {
	err := WriteConfig(c, application.Config.SmbConfigPath)
	if err != nil {
		return err
	}
	err = RestartSMBService()
	if err != nil {
		return err
	}
	return nil
}
func (c *Config) AddFolder(name string, properties map[string]string) error {
	for _, section := range c.Sections {
		if section.Name == name {
			return FolderAlreadyExist
		}
	}
	err := InitSMBFolder(properties["path"])
	if err != nil {
		return err
	}
	c.Sections = append(c.Sections, &Section{
		Name:   name,
		Fields: properties,
	})

	err = c.SaveFileAndRestart()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) RemoveFolder(name string) error {
	targetIndex := -1
	for index, section := range c.Sections {
		if section.Name == name {
			targetIndex = index
		}
	}
	if targetIndex == -1 {
		return nil
	}
	c.Sections[targetIndex] = c.Sections[len(c.Sections)-1]
	c.Sections[len(c.Sections)-1] = nil
	c.Sections = c.Sections[:len(c.Sections)-1]
	err := c.SaveFileAndRestart()
	if err != nil {
		return err
	}
	return nil
}
func (c *Config) UpdateFolder(name string, properties map[string]string) error {
	for _, section := range c.Sections {
		if section.Name == name {
			for key, value := range properties {
				if value == "-" {
					delete(section.Fields, key)
					continue
				}
				// path updated
				if key == "path" && section.Fields["path"] != value {
					err := InitSMBFolder(value)
					if err != nil {
						return err
					}
				}
				section.Fields[key] = value
			}
		}
	}
	err := c.SaveFileAndRestart()
	if err != nil {
		return err
	}
	return nil
}
