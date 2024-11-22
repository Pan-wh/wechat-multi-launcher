package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

// 获取微信安装路径
func getWeChatInstallPath() (string, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Tencent\WeChat`, registry.READ)
	if err != nil {
		return "", fmt.Errorf("无法打开注册表键: %v", err)
	}
	defer key.Close()

	installPath, _, err := key.GetStringValue("InstallPath")
	if err != nil {
		return "", fmt.Errorf("无法读取注册表值: %v", err)
	}

	return installPath, nil
}

// 启动微信
func launchWeChat(installPath string) error {
	wechatExe := fmt.Sprintf(`%s\WeChat.exe`, installPath)
	err := windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(wechatExe), nil, nil, windows.SW_NORMAL)
	if err != nil {
		return fmt.Errorf("无法启动微信: %v", err)
	}
	return nil
}

func main() {
	installPath, err := getWeChatInstallPath()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	// 获取用户输入
	fmt.Print("请输入要启动的微信实例数量：")
	var input string
	fmt.Scanln(&input)

	// 处理输入
	numInstances, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || numInstances <= 0 {
		fmt.Printf("无效的输入: %v\n", input)
		os.Exit(1)
	}

	for i := 0; i < numInstances; i++ {
		if err := launchWeChat(installPath); err != nil {
			fmt.Printf("启动失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("成功启动微信实例 #%d\n", i+1)
	}
}
