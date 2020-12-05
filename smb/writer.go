package smb

import (
	"bufio"
	"fmt"
	"os"
)

func WriteConfig(config *Config, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, section := range config.Sections {
		fmt.Fprintln(w, fmt.Sprintf("[%s]", section.Name))
		for key, value := range section.Fields {
			fmt.Fprintln(w, fmt.Sprintf("    %s = %s", key, value))
		}
		fmt.Fprint(w, "\n\n")
	}
	return w.Flush()
}
