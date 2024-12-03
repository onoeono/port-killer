package dto

type PortUsageDTO struct {
	Protocol  string `json:"protocol"`  // 协议（如 TCP、UDP）
	LocalAddr string `json:"localAddr"` // 本地地址
	Port      string `json:"port"`      // 端口号
	State     string `json:"state"`     // 连接状态
	PID       string `json:"pid"`       // 进程ID
	Program   string `json:"program"`   //程序名称
}
