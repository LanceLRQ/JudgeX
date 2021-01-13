package polygon

import "gorm.io/gorm"

// 用户账户表表
type Account struct {
    gorm.Model
    // 登录名称
    UserName string `gorm:"type:varchar(64);unique;not null"`
    // 登录密码
    PassWord string
    // 登录邮箱
    Email string
    // 邮箱是否验证
    EmailValidated bool
    // 账号锁定
    Locked bool
    // 系统管理员身份
    IsAdmin bool
    // 创建题目
    PermCreateProblem bool
    // 管理非自己创建的题目
    PermManageProblems bool
    // 查看非自己创建题目的数据
    PermProblemDatas bool
}

// 登录日志表
type LoginLog struct {
    gorm.Model
    // IP地址
    IpAddress string
    // 登录用户
    LoginAccount Account `gorm:"foreignkey:ID;association_foreignkey:LoginAccountId"`
    // 登录用户Id
    LoginAccountId uint `gorm:"index"`
    // UA信息
    UserAgent string
}

// WeJudge SSO服务绑定信息表
// WeJudge SSO service binding records.
type WeJudgeSSOBinding struct {
    gorm.Model
    // 登录用户
    LoginAccount Account `gorm:"foreignkey:ID;association_foreignkey:LoginAccountId"`
    // 登录用户Id
    LoginAccountId uint `gorm:"index"`
    // WeJudge 账号Id
    WeJudgeAccountId string
}
