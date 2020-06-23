package models

import (
	"github.com/jinzhu/gorm"
)

// 用户账户表表
type Account struct {
	gorm.Model
	UserName 			string	`gorm:"unique;not null"`	// 登录名称
	PassWord 			string								// 登录密码
	Email 				string								// 登录邮箱
	EmailValidated 		bool								// 邮箱是否验证
	Locked 				bool								// 账号锁定

	IsAdmin 			bool								// 系统管理员身份
	PermCreateProblem 	bool								// 创建题目
	PermManageProblems 	bool								// 管理非自己创建的题目
	PermProblemDatas 	bool 								// 查看非自己创建题目的数据
}

// 登录日志表
type LoginLog struct {
	gorm.Model
	IpAddress 		string									// IP地址
	LoginAccount 	Account									// 登录用户
	UserAgent 		string 									// UA信息
}

// WeJudge SSO服务绑定信息表
// WeJudge SSO service binding records.
type WeJudgeSSOBinding struct {
	gorm.Model
	LoginAccount 		Account								// 登录用户
	WeJudgeAccountId	string								// WeJudge 账号Id
}