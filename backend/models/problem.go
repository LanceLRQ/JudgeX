package models

import "github.com/jinzhu/gorm"

// 题目数据
type Problem struct {
	gorm.Model
	Title string									// 题目标题
	Author Account `gorm:"foreignkey:ID;association_foreignkey:AuthorId"`		// 题目作者
	AuthorId uint 									// 题目作者Id
	Type uint										// 题目类型

	Legend string `gorm:"type:longtext"`			// 题目正文描述
	InputFormat string 								// 题目输入要求
	OutputFormat string 							// 题目输出要求
	Samples string									// 样例数据JSON
	JudgeOptions string `gorm:"type:longtext"`		// 评测设置信息JSON

	Hint string	`gorm:"type:longtext"`				// 解题提示
	Source string `gorm:"type:text"`				// 题目来源

	Available bool									// 评测是否可用
}
