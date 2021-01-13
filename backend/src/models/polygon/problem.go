package polygon

import (
    "gorm.io/gorm"
)

// 题目数据
type Problem struct {
    gorm.Model
    // 题目标题
    Title string
    // 题目作者
    Author Account `gorm:"foreignkey:ID;association_foreignkey:AuthorId"`
    // 题目作者Id
    AuthorId uint `gorm:"index"`
    // 题目类型
    ProblemType int
    // 题目正文描述
    Description string `gorm:"type:longtext"`
    // 题目输入要求
    Input string `gorm:"type:text"`
    // 题目输出要求
    Output string `gorm:"type:text"`
    // 样例数据JSON
    Sample string `gorm:"type:longtext"`
    // 评测设置信息JSON
    JudgeOptions string `gorm:"type:longtext"`
    // 解题提示
    Tips string `gorm:"type:longtext"`
    // 题目来源
    Source string `gorm:"type:text"`
    // 评测是否可用
    Available bool
    // 标签
    Tags []ProblemTag `gorm:"many2many:problem_tags_related"`
}

type ProblemTag struct {
    gorm.Model
    // 反向标签
    Problems []Problem `gorm:"many2many:problem_tags_related"`
}
