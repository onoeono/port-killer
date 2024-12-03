package strategy

import (
	"os/exec"
	"port-killer/backend/dto"
	"strings"
)

type LinuxPortUsageStrategy struct{}

func (w *LinuxPortUsageStrategy) GetPortUsage() ([]dto.PortUsageDTO, error) {
	cmd := exec.Command("netstat", "-tulnp")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var dtoList []dto.PortUsageDTO
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		ok, protocol := isTcpOrUdp(line)
		if !ok {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) <= 3 {
			continue
		}

		localAddr := fields[3]
		ip, port, found := strings.Cut(localAddr, ":")
		if !found {
			continue
		}
		portUsage := dto.PortUsageDTO{
			Protocol:  protocol,
			LocalAddr: ip,
			Port:      port,
			State:     "N/A",
			PID:       "N/A",
			Program:   "N/A",
		}

		var pidProgram string
		switch protocol {
		case "tcp":
			portUsage.State = fields[5]
			pidProgram = fields[6]
		case "udp":
			pidProgram = fields[5]
		}
		if pidProgram != "-" {
			pid, program, exist := strings.Cut(pidProgram, "/")
			if exist {
				portUsage.PID = pid
				portUsage.Program = program
			}
		}

		dtoList = append(dtoList, portUsage)
	}
	return dtoList, nil
}

func isTcpOrUdp(str string) (bool, string) {
	if strings.HasPrefix(str, "tcp") {
		return true, "tcp"
	}
	if strings.HasPrefix(str, "udp") {
		return true, "udp"
	}
	return false, ""
}
