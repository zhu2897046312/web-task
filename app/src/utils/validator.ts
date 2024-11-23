// 表单验证规则
export const validators = {
  // 手机号验证
  phone: (value: string) => {
    return /^1[3-9]\d{9}$/.test(value)
  },

  // 邮箱验证
  email: (value: string) => {
    return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(value)
  },

  // 密码验证（至少8位，包含数字和字母）
  password: (value: string) => {
    return /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(value)
  }
} 