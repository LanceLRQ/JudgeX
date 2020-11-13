package enums

const (
	JudgeLangGCC 		= 1			// GNU C
	JudgeLangGCPP		= 2			// GNU C++
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

const (
	JudgeFlagQueuing				= -2 		// -2 Queuing
	JudgeFlagRunning 				= -1   		// -1 Running
	JudgeFlagAC 					= 0   		// 0 Accepted
	JudgeFlagPE  					= 1	    	// 1 Presentation Error
	JudgeFlagTLE  					= 2			// 2 Time Limit Exceeded
	JudgeFlagMLE  					= 3			// 3 Memory Limit Exceeded
	JudgeFlagWA  					= 4	    	// 4 Wrong Answer
	JudgeFlagRE 					= 5	    	// 5 Runtime Error
	JudgeFlagOLE  					= 6			// 6 Output Limit Exceeded
	JudgeFlagCE 					= 7	    	// 7 Compile Error
	JudgeFlagSE  					= 8     	// 8 System Error
	JudgeFlagRejudge				= 9			// 9 Rejudge Queuing

	JudgeFlagSpecialJudgerTimeout  	= 10    	// 10 Special Judger Timeout
	JudgeFlagSpecialJudgerError 	= 11    	// 11 Special Judger ERROR
	JudgeFlagSpecialJudgerFinish    = 12		// 12 Special Judger Finish, Need Standard Checkup
)