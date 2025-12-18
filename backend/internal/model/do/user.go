// Package do 定义数据操作对象
// 用于数据库操作的输入输出
package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User 用户数据操作对象
type User struct {
	g.Meta      `orm:"table:sys_user, do:true"`
	UserId      interface{} // 用户ID
	DeptId      interface{} // 部门ID
	UserName    interface{} // 用户账号
	NickName    interface{} // 用户昵称
	UserType    interface{} // 用户类型
	Email       interface{} // 用户邮箱
	Phonenumber interface{} // 手机号码
	Sex         interface{} // 用户性别
	Avatar      interface{} // 头像地址
	Password    interface{} // 密码
	Status      interface{} // 帐号状态
	DelFlag     interface{} // 删除标志
	LoginIp     interface{} // 最后登录IP
	LoginDate   *gtime.Time // 最后登录时间
	CreateBy    interface{} // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    interface{} // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      interface{} // 备注
}
