package orms

import "github.com/jinzhu/gorm"

// 评测记录
type JudgeRecord struct {
	gorm.Model

	// 提交者
	Author Account `gorm:"foreignkey:ID;association_foreignkey:AuthorId"`
	// 题提交者Id
	AuthorId 			uint
	// 题目关联
	Problem Problem `gorm:"foreignkey:ID;association_foreignkey:ProblemId"`
	// 题目关联Id
	ProblemId 			uint

	// 评测状态
	Flag 				uint
	// 评测语言
	CodeLang 			uint

	// 最大运行时间（毫秒）
	ExeTime 			uint
	// 最大内存占用（KB）
	ExeMem 				uint
	// 代码长度（字节）
	CodeLength 			uint
	// 代码行数
	CodeLines 			uint
	// 代码内容
	CodeContent    		string 		`gorm:"type:longtext"`

	// 运行时信息
	RuntimeInfo			string		`gorm:"type:text"`
	// 测试数据正确行数
	SameLines			uint
	// 测试数据总行数
	TotalLines			uint
	// 评测任务ID （用于和判题模块交互）
	JudgeTaskId			uint64
	// 评测结果信息（JSON）
	JudgeResult 		string 		`gorm:"type:longtext"`
	// 优先使用的节点
	TargetNode			uint
	// 查重分组ID （同一分组的提交将被执行查重，不填则不执行查重）
	CrossCheckGroupId	string 		`gorm:"type:varchar(128);index"`
}

