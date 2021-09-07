package smb

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type StatusMonitor struct {
	Process      []*StatusProcess
	StatusShares []*StatusShares
}
type StatusProcess struct {
	PID      string            `json:"pid"`
	Username string            `json:"username"`
	Group    string            `json:"group"`
	Machine  string            `json:"machine"`
	Raw      map[string]string `json:"raw"`
}
type StatusShares struct {
	Service   string            `json:"service"`
	PID       string            `json:"pid"`
	Machine   string            `json:"machine"`
	ConnectAt *time.Time        `json:"connectAt"`
	Raw       map[string]string `json:"raw"`
}

func (s *StatusMonitor) getData(arg string) ([]map[string]string, error) {
	cmd := exec.Command("smbstatus", arg)
	rawOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	output := string(rawOutput)
	lines := strings.Split(output, "\n")
	dividerIdx := -1
	for idx, line := range lines {
		fmt.Println(line)
		if strings.HasPrefix(line, "-") {
			dividerIdx = idx
		}
	}
	if dividerIdx == -1 {
		return nil, errors.New("parse error")
	}
	// parse table header
	headerLine := lines[dividerIdx-1]
	headerItem := strings.Split(headerLine, "  ")
	headers := make([]string, 0)
	for _, rawHeader := range headerItem {
		header := strings.TrimSpace(rawHeader)
		if len(header) == 0 {
			continue
		}
		headers = append(headers, header)
	}
	headerPosition := make(map[string]int, 0)
	for _, header := range headers {
		headerIdx := strings.Index(headerLine, header)
		if headerIdx != -1 {
			headerPosition[header] = headerIdx
		}
	}
	// parse row
	values := make([]map[string]string, 0)
	for _, valueLine := range lines[dividerIdx+1:] {
		if len(strings.TrimSpace(valueLine)) == 0 {
			continue
		}
		value := make(map[string]string, 0)
		for hIdx, header := range headers {
			if hIdx == len(headers)-1 {
				value[header] = strings.TrimSpace(valueLine[headerPosition[header]:])
				continue
			}
			value[header] = strings.TrimSpace(valueLine[headerPosition[header]:headerPosition[headers[hIdx+1]]])
		}
		values = append(values, value)
	}
	return values, nil
}
func (s *StatusMonitor) GetProcess() ([]map[string]string, error) {
	return s.getData("-p")
}
func (s *StatusMonitor) GetShares() ([]map[string]string, error) {
	return s.getData("--shares")
}
func (s *StatusMonitor) Run() {
	go func() {
		for true {
			processList := make([]*StatusProcess, 0)
			processValue, err := s.GetProcess()
			if err == nil {
				for _, processRow := range processValue {
					process := &StatusProcess{}
					process.PID = processRow["PID"]
					process.Machine = processRow["Machine"]
					process.Group = processRow["Group"]
					process.Username = processRow["Username"]
					process.Raw = processRow
					processList = append(processList, process)
				}
			}
			s.Process = processList
			sharesList := make([]*StatusShares, 0)
			sharesValues, err := s.GetShares()
			if err == nil {
				for _, sharesRow := range sharesValues {
					shares := &StatusShares{
						Service: sharesRow["Service"],
						PID:     sharesRow["pid"],
						Machine: sharesRow["Machine"],
						Raw:     sharesRow,
					}
					timeLayout := "Mon Jan  2 15:04:05 PM 2006 MST"
					connectAt, err := time.Parse(timeLayout, sharesRow["Connected at"])
					if err == nil {
						shares.ConnectAt = &connectAt
					}
					sharesList = append(sharesList, shares)
				}
			}
			s.StatusShares = sharesList
			<-time.After(3 * time.Second)
		}
	}()
}
