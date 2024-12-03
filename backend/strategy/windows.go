package strategy

import (
	"os/exec"
	"port-killer/backend/dto"
	"strings"
)

type WindowsPortUsageStrategy struct{}

func (s *WindowsPortUsageStrategy) GetPortUsage() ([]dto.PortUsageDTO, error) {
	cmd := exec.Command("netstat", "-ano")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// 解析 netstat 输出
	var portUsages []dto.PortUsageDTO
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "TCP") || strings.HasPrefix(line, "UDP") {
			fields := strings.Fields(line)
			if len(fields) > 4 {
				localAddr := fields[1]
				port := localAddr[strings.LastIndex(localAddr, ":")+1:]

				portUsage := dto.PortUsageDTO{
					Protocol:  fields[0],
					LocalAddr: localAddr,
					Port:      port,
					PID:       fields[4],
					State:     fields[3], // 连接状态
				}
				portUsages = append(portUsages, portUsage)
			}
		}
	}
	return portUsages, nil
}
