package python

import (
	"fmt"
	"os/exec"
	"time"
)

func AutomateCrawler() {
	for {
		// 执行 Python 文件
		cmd := exec.Command("python", "s.py")
		err := cmd.Run()
		if err != nil {
			fmt.Println("执行 Python 文件出错:", err)
		}

		// 等待 24 小时
		time.Sleep(24 * time.Hour)
	}
}
