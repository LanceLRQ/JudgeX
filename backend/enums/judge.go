package enums

const (
	JudgeLangGCC 		= 1			// GNU C
	JudgeLangGPP		= 2			// GNU C++
	JudgeLangJava		= 4
	JudgeLangPython2	= 8
	JudgeLangPython3	= 16
	JudgeLangCSharp		= 32
	JudgeLangGolang		= 64
	JudgeLangNodeJS		= 128
	JudgeLangPHP		= 256
	JudgeLangRuby		= 512
)

const (
	TestLibTypeNone 		= 0
	TestLibTypeChecker 		= 1
	TestLibTypeGenerator 	= 2
	TestLibTypeInteractor   = 3
	TestLibTypeValidator    = 4
)

const (
	SpecialJudgeDisabled 		= 0		// 停用
	SpecialJudgeCheckerMode 	= 1		// 检查模式(WeJudge版本)
	SpecialJudgeInteractiveMode = 2		// 交互模式(WeJudge版本)
	SpecialJudgeTestlibMode   	= 3		// Testlib模式
)