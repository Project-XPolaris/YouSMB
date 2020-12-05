package smb

import (
	"fmt"
	"os/exec"
	"strings"
	"yousmb/application"
)

func StartSMBService() error {
	parts := strings.Split(application.Config.StartScript, " ")
	out, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		return err
	}
	fmt.Println("Start SMB Service")
	output := string(out[:])
	fmt.Println(output)
	return nil
}

func StopSMBService() error {
	parts := strings.Split(application.Config.StopScript, " ")
	out, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		return err
	}
	fmt.Println("Stop SMB Service")
	output := string(out[:])
	fmt.Println(output)
	return nil
}

func RestartSMBService() error {
	parts := strings.Split(application.Config.RestartScript, " ")
	out, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		return err
	}
	fmt.Println("Restart SMB Service")
	output := string(out[:])
	fmt.Println(output)
	return nil
}
