package services

// CreatePassword 创建密码
func CreatePassword() string {
	return "123132"
}

// SendCode 发送验证码
func SendCode(tel string) bool {
	return true
}

// VerifyCode 校验验证码
func VerifyCode(sendCode string, newCode string) bool {
	if sendCode == newCode {
		return true
	}
	return false
}
