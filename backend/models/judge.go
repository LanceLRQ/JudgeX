package models

import "backend/enums"

// 评测设置（非数据库表）
// map类型的表示需要针对不同的语言进行设置，key为uint，是代码语言的标识
type JudgeConfig struct {
	Version string							`json:"version"`					// 版本号(兼容数据结构，基本没用)
	TimeLimit map[uint]uint					`json:"time_limit"`					// 时间限制
	MemLimit map[uint]uint					`json:"mem_limit"`					// 内存空间限制
	ProblemType int							`json:"problem_type"`				// 题目类型

	TestCases []JudgeTestCase 				`json:"test_cases"`					// 测试数据
	DemoCases map[uint]JudgeDemoCase 		`json:"demo_cases"`					// 代码填空样例数据
	AnswerCases map[uint]string				`json:"answer_cases"`				// 答案样例代码

	TestLib TestlibOptions 					`json:"TestLib"`					// testlib设置
	SpecialJudge uint 						`json:"special_judge"`				// 特殊评测模式
	SpecialJudgeConfig SpecialJudgeConfig 	`json:"special_judge_config"`		// 特殊评测设置信息

	// testlib程序、特判程序等工作代码存放点，名称不能重复
	Libraries []JudgeFileInfo 				`json:"libraries"`
	// 共享环境文件地址列表（比如可以预设一些允许程序读写访问的文件）
	EnvFiles []JudgeFileInfo 				`json:"env_files"`
}

// 测试样例
type JudgeTestCase struct {
	Handle string 						`json:"handle"`					// handle
	Name string 						`json:"name"`					// 测试样例名称
	Order int 							`json:"order"`					// 运行优先级
	UpdateTime float64 					`json:"update_time"`			// 更新时间
	Visible bool 						`json:"visible"`				// 数据是否公开
	Available bool 						`json:"available"`				// 数据是否启用
	UseGenerator bool 					`json:"use_generator"`			// 是否使用脚本生成（Testlib）
	Generator string 					`json:"generator"`				// 生成器名字
	GeneratorArgs string 				`json:"generator_args"`			// 生成器命令参数
	StdinSize int						`json:"stdin_size"`				// 输入数据大小
	StdoutSize int						`json:"stdout_size"`			// 输出数据大小
	// 如果没有设置检查器的话，gen来的数据默认视作是通过的，这个和CF不一样（兼容旧版的数据）
	ValidatorVerdict bool 				`json:"validator_verdict"`		// 检查结果
	ValidatorComment bool 				`json:"validator_comment"`		// 检查结果备注（输出信息）
}

// 代码填空样例
type JudgeDemoCase struct {
	Handle string 						`json:"handle"`					// handle
	Name string 						`json:"name"`					// 代码区域名称
	Answers string 						`json:"answers"`				// 回答信息
	Demo string 						`json:"demo"`					// 样例代码（预设用）
	Line int 							`json:"line"`					// 插入位置
}

// 评测用文件信息
type JudgeFileInfo struct {
	FileName string 					`json:"file_name"`				// 文件名
	FileExt string 						`json:"file_ext"`				// 文件扩展名
	FileMIME string 					`json:"file_mime"`				// 文件MIME
	FileSize string 					`json:"file_size"`				// 文件大小
	IsCode bool 						`json:"is_code"`				// 是否是代码文件
	Language uint 						`json:"language"`				// 代码文件的语言
	TestlibType uint 					`json:"testlib_type"`			// testlib功能类型
	TestlibVersion	string 				`json:"testlib_version"`		// testlib版本
	CompileTarget	string 				`json:"compile_target"`			// 编译后的目标文件名称
	CreateTime uint						`json:"create_time"`			// 创建时间
	UpdateTime uint						`json:"update_time"`			// 修改时间
}

type SpecialJudgeConfig struct {
	Program string						`json:"program"` 				// 特殊评测目标程序位置（相对）
	CodeContent string					`json:"code_content"` 			// 特判程序源代码
	Libaray string						`json:"libaray"` 				// 使用已有的文件（设置后上面两个自动失效）
	Language uint 						`json:"language"`				// 特判程序编译语言（启用Testlib后只支持c++）
	RedirectSTD bool 					`json:"redirect_std"`			// 重定向Std流到判题程序（Testlib时无效）
	TimeLimit uint						`json:"time_limit"`				// 时间限制（ms)
	MemLimit uint						`json:"mem_limit"`				// 内存空间限制（kb）
	// 验证样例（基于Testlib）
	CheckerCases []SpecialJudgeCheckerCases	`json:"checker_cases"`
}

type SpecialJudgeCheckerCases struct {
	Input string						`json:"input"` 					// 输入数据（限制1k内)
	Output string						`json:"output"` 				// 输出数据（限制1k内)
	Answer string						`json:"answer"` 				// 答案数据（限制1k内)
	Verdict bool 						`json:"verdict"`				// 是否执行过检查
	ExpectedVerdict int 				`json:"expected_verdict"`		// 期望结果
	CheckerVerdict int 					`json:"checker_verdict"`		// 检查结果
	CheckerComment bool 				`json:"checker_comment"`		// 检查结果备注（输出信息）
}


type TestlibOptions struct {
	Version string						`json:"version"`				// 声明它要用哪个版本的Testlib
	Language uint 						`json:"language"`				// 编译语言（默认是CPP，未来支持Java，其他不支持）
	Validator string 					`json:"validator"`				// 设置Validator的名字
	ValidatorCase []TestlibValidatorCase `json:"validator_case"`		// 设置校验器的测试样例（一般是人工输入）
	// 未来这边可以加入对拍(stress)功能
}

type TestlibValidatorCase struct {
	Input string						`json:"input"`					// 输入数据（建议限制1k内)
	Verdict bool 						`json:"verdict"`				// 是否执行过检查
	ExpectedVerdict bool 				`json:"expected_verdict"`		// 期望结果
	ValidatorVerdict bool 				`json:"validator_verdict"`		// 检查结果
	ValidatorComment bool 				`json:"validator_comment"`		// 检查结果备注（输出信息）
}

/* Object Generators */

func NewJudgeConfig() JudgeConfig {
	ret := JudgeConfig{}
	ret.Version = "1.0"
	ret.ProblemType = enums.ProblemTypeCode
	ret.SpecialJudge = enums.SpecialJudgeDisabled
	ret.TestLib = NewTestlibOptions()
	return ret
}
func NewJudgeTestCase() JudgeTestCase {
	ret := JudgeTestCase{}
	ret.Available = true
	ret.ValidatorVerdict = true
	return ret
}
func NewJudgeDemoCase() JudgeDemoCase {
	ret := JudgeDemoCase{}
	return ret
}
func NewJudgeFileInfo() JudgeFileInfo {
	ret := JudgeFileInfo{}
	ret.Language = enums.JudgeLangGCPP
	ret.TestlibType = enums.TestLibTypeNone
	ret.TestlibVersion = "latest"
	return ret
}
func NewSpecialJudgeConfig() SpecialJudgeConfig {
	ret := SpecialJudgeConfig{}
	ret.Language = enums.JudgeLangGCPP
	ret.RedirectSTD = true
	ret.TimeLimit = 1000
	ret.MemLimit = 32768
	return ret
}
func NewSpecialJudgeCheckerCases() SpecialJudgeCheckerCases {
	ret := SpecialJudgeCheckerCases{}
	ret.CheckerVerdict = enums.JudgeFlagAC
	ret.ExpectedVerdict = enums.JudgeFlagAC
	return ret
}
func NewTestlibOptions() TestlibOptions {
	ret := TestlibOptions{}
	ret.Language = enums.JudgeLangGCPP
	return ret
}
func NewTestlibValidatorCase() TestlibValidatorCase {
	ret := TestlibValidatorCase{}
	return ret
}