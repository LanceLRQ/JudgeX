package polygon

import (
    "gorm.io/gorm"
    "time"
)

// 客户端表
type ApplicationClient struct {
    gorm.Model
    // 应用标识 (基于snowflake的Base64模式输出)
    AppKey string `gorm:"type:varchar(64);unique_index"`
    // 应用密钥
    AppSecret string `gorm:"type:varchar(128)"`
    // 应用名称
    AppName string
    // 应用描述
    AppDescription string
    // 应用官网
    AppUrl string
    // 应用Logo
    AppLogo string
    // 数据回调地址
    AccessCallbackUrl string
    // 过期时间（过期后不能再使用RPC相关功能）
    ExpireTime int
    // 基本数据获取，默认10000次/h
    LimitQueryPerHour uint
    // 执行评测，默认1000 次/h
    LimitJudgePerHour uint
    // 大规模数据，默认1000次/h
    LimitLargeQueryPerHour uint
}

// 系统设置
type GlobalSettings struct {
    // 禁用题目编辑功能
    DisabledProblemEdit bool
    // 暂停全站评测
    PauseJudge bool
}

// 评测节点组
type JudgeMachineNodeGroup struct {
    gorm.Model
    // 分组名称
    Name string `gorm:"type:varchar(256)"`
    // 是否启用这个分组的节点 （优先级高一点）
    Enable bool
}

/*	评测节点
动态节点的IPAddr参数为每个判题机自己的IP和端口号
Online参数表示此节点是否在线(和评测节点是否存活无关，纯开关意义)
节点由Watcher启动时自动注册到数据库中，如果ActiveTime早于当前时间3分钟以上，视作节点已经离线。
*/
type JudgeMachineNode struct {
    gorm.Model
    // 服务名
    Name string `gorm:"type:varchar(256)"`
    // 节点组
    Group JudgeMachineNodeGroup `gorm:"foreignkey:ID;association_foreignkey:GroupId"`
    // 节点组
    GroupId uint `gorm:"index"`
    // IP地址
    IPAddr string `gorm:"type:varchar(128)"`
    // 是否启用节点
    Enable bool
    // 第一次注册时间
    CreateTime time.Time
    // 存活心跳时间
    ActiveTime time.Time
    // 接受语言，不设置默认为all
    Languages string `gorm:"type:text"`
}
