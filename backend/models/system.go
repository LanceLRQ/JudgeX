package models

/**
	系统级别的设置
*/

// 客户端表
type ApplicationClient struct {
	AppKey string 	`gorm:"type:varchar(64);unique_index"`	// 应用标识 (基于snowflake的Base64模式输出)
	AppSecret string `gorm:"type:varchar(128)"`				// 应用密钥

	AppName string 											// 应用名称
	AppDescription string									// 应用描述
	AppUrl string 											// 应用官网
	AppLogo string 											// 应用Logo

	AccessCallbackUrl	string 								// 数据回调地址

	ExpireTime int											// 过期时间（过期后不能再使用RPC相关功能）
}