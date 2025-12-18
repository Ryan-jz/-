// Package consts 定义系统常量
// 包含系统中使用的各种常量定义
package consts

const (
	// ContextKey 上下文键名
	ContextKey = "ContextKey"
	
	// DefaultPageSize 默认分页大小
	DefaultPageSize = 10
	
	// MaxPageSize 最大分页大小
	MaxPageSize = 100
)

// 用户状态
const (
	UserStatusNormal  = 1 // 正常
	UserStatusDisable = 0 // 禁用
)

// 菜单类型
const (
	MenuTypeDir    = "M" // 目录
	MenuTypeMenu   = "C" // 菜单
	MenuTypeButton = "F" // 按钮
)

// 数据权限范围
const (
	DataScopeAll             = "1" // 全部数据权限
	DataScopeCustom          = "2" // 自定义数据权限
	DataScopeDept            = "3" // 本部门数据权限
	DataScopeDeptAndChild    = "4" // 本部门及以下数据权限
	DataScopeSelf            = "5" // 仅本人数据权限
)

// 是否标识
const (
	YesNo_Yes = 1 // 是
	YesNo_No  = 0 // 否
)

// 操作日志类型
const (
	LogTypeLogin  = 1 // 登录日志
	LogTypeOperate = 2 // 操作日志
)
