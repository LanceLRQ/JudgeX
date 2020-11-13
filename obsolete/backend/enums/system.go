package enums

const (
	JudgerUnitTypeStatic 	= 0			// 静态节点  (docker)
	JudgerUnitTypeDynamic	= 1			// 动态节点  (k8s pod)
	JudgerUnitTypeSpecial  	= 2			// 专用节点  (比如提供给某场比赛专用的)
	JudgerUnitTypeVIP  		= 3			// VIP节点  (提供给VIP的)
)

const (
	ApplicationPermissionVisit 	= 1 << iota				// 访问信息
	ApplicationPermissionJudge							// 执行评测
	ApplicationPermissionResult							// 获取评测结果
	ApplicationPermissionResultData						// 获取运行数据
	ApplicationPermissionProblemData					// 获取题目数据
)