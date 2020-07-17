package orms

import "github.com/jinzhu/gorm"

// 评测记录
type JudgeRecord struct {
	gorm.Model
	Author   		Account 		`gorm:"foreignkey:ID;association_foreignkey:AuthorId"` 		// 提交者
	AuthorId 		uint    																	// 题提交者Id
	Problem   		Problem 		`gorm:"foreignkey:ID;association_foreignkey:ProblemId"` 	// 题目关联
	ProblemId 		uint    																	// 题目关联Id

	Flag 			uint 											// 评测状态
	CodeLang 		uint											// 评测语言

	ExeTime 		uint 											// 最大运行时间（毫秒）
	ExeMem 			uint 											// 最大内存占用（KB）
	CodeLength 		uint 											// 代码长度（字节）
	CodeLines 		uint 											// 代码行数
	CodeContent    	string 			`gorm:"type:longtext"`			// 代码内容

	RuntimeInfo		string			`gorm:"type:text"`				// 运行时信息
	SameLines		uint											// 测试数据正确行数
	TotalLines		uint 											// 测试数据总行数

	JudgeTaskId		uint64 											// 评测任务ID （用于和判题模块交互）
	JudgeResult 	string 			`gorm:"type:longtext"`			// 评测结果信息（JSON）

	TargetNode		uint											// 优先使用的节点
	CrossCheckGroupId	string 		`gorm:"type:varchar(128);index"`		// 查重分组ID （同一分组的提交将被执行查重，不填则不执行查重）
}

