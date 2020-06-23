package models

// 评测设置（非数据库表）
type JudgeConfig struct {
	Version string
	TimeLimit map[uint]uint			`json:"time_limit"`					// 时间限制
	MemLimit map[uint]uint			`json:"mem_limit"`					// 内存空间限制
	ProblemType int					`json:"problem_type"`				// 题目类型

	TestCases bool 					`json:"test_cases"`					// 测试数据
	DemoCases bool 					`json:"demo_cases"`					// 代码填空样例数据
	AnswerCases bool 				`json:"answer_cases"`				// 答案样例数据

	TestLib bool 					`json:"TestLib"`					// testlib设置
	SpecialJudge uint 				`json:"special_judge"`				// 特殊评测模式
	// 特殊评测设置
	SpecialJudgeConfig SpecialJudgeConfig `json:"special_judge_config"`

	// testlib程序、特判程序等工作代码存放点，名称不能重复
	Libraries []JudgeFileInfo 		`json:"libraries"`
	// 共享环境文件地址列表（比如可以预设一些允许程序读写访问的文件）
	EnvFiles []JudgeFileInfo 		`json:"env_files"`
}

// 评测用文件信息
type JudgeFileInfo struct {
	FileName string 					`json:"file_name"`				// 文件名
	FileExt string 						`json:"file_ext"`				// 文件扩展名
	FileMIME string 					`json:"file_mime"`				// 文件MIME
	FileSize string 					`json:"file_size"`				// 文件大小
	IsCode bool 						`json:"is_code"`				// 是否是代码文件
	Language uint 						`json:"language"`				// 代码文件的语言
	TestLibType uint 					`json:"testlib_type"`			// testlib功能类型
	TestLibVersion	string 				`json:"testlib_version"`		// testlib版本
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
	CheckerCases SpecialJudgeCheckerCases	`json:"checker_cases"`
}

type SpecialJudgeCheckerCases struct {
	Input string						`json:"input"` 					// 输入数据（限制1k内)
	Output string						`json:"output"` 				// 输出数据（限制1k内)
	Answer string						`json:"answer"` 				// 答案数据（限制1k内)
	Verdict bool 						`json:"verdict"`				// 是否执行过检查
	ExpectedVerdict bool 				`json:"expected_verdict"`		// 期望结果
	CheckerVerdict bool 				`json:"checker_verdict"`		// 检查结果
	CheckerComment bool 				`json:"checker_comment"`		// 检查结果备注（输出信息）
}
