package orms

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
	系统级别的设置
*/

// 客户端表
type ApplicationClient struct {
	gorm.Model
	AppKey 				string 	`gorm:"type:varchar(64);unique_index"`	// 应用标识 (基于snowflake的Base64模式输出)
	AppSecret 			string 	`gorm:"type:varchar(128)"`				// 应用密钥

	AppName 			string 											// 应用名称
	AppDescription 		string											// 应用描述
	AppUrl 				string 											// 应用官网
	AppLogo 			string 											// 应用Logo

	AccessCallbackUrl	string 											// 数据回调地址

	ExpireTime 			int												// 过期时间（过期后不能再使用RPC相关功能）
	Permission			uint											// 访问权限

	LimitQueryPerHour	uint											// 基本数据获取，默认10000次/h
	LimitJudgePerHour	uint											// 执行评测，默认1000 次/h
	LimitLargeQueryPerHour	uint										// 大规模数据，默认1000次/h
}


// 系统设置
type GlobalSettings struct {
	DisabledProblemEdit			bool									// 禁用题目编辑功能
	PauseJudge					bool									// 暂停全站评测
	DisabledDynamicNode			bool									// 禁用动态节点
}


/*	 评测节点

	动态节点的IPAddr可以存放k8s的namespace名称等，handle一般为pod的hostname，Online参数无效
	非动态节点的Handle通过配置文件固定，Online参数表示此节点是否在线(和评测节点是否存活无关，纯开关意义)

	节点由Watcher启动时自动注册到数据库中，如果ActiveTime早于当前时间3分钟以上，视作节点已经离线。
*/
type JudgerNode struct {
	Name 				string		`gorm:"type:varchar(256)"`			// 服务名
	Handle 				string		`gorm:"type:varchar(256)"`			// 标识符
	Type				uint											// 节点类型
	IPAddr				string 		`gorm:"type:varchar(128)"`			// IP地址
	Online 				bool 											// 是否在线
	CreateTime			time.Time										// 第一次注册时间
	ActiveTime			time.Time										// 存活心跳时间
}