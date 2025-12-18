// Package admin 提供管理后台接口
// 需要认证后才能访问
package admin

import (
	"context"

	"gf-admin/api/v1/admin"
	"gf-admin/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// UserController 用户管理控制器
// 处理用户的增删改查操作
var User = cUser{}

type cUser struct{}

// List 获取用户列表
// 支持分页和条件查询
func (c *cUser) List(ctx context.Context, req *admin.UserListReq) (res *admin.UserListRes, err error) {
	// 构建查询
	model := g.DB().Model("sys_user").Where("del_flag", 0)

	// 条件查询
	if req.UserName != "" {
		model = model.WhereLike("user_name", "%"+req.UserName+"%")
	}
	if req.Phonenumber != "" {
		model = model.WhereLike("phonenumber", "%"+req.Phonenumber+"%")
	}
	if req.Status >= 0 {
		model = model.Where("status", req.Status)
	}

	// 分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = consts.DefaultPageSize
	}
	if pageSize > consts.MaxPageSize {
		pageSize = consts.MaxPageSize
	}

	// 查询总数
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 查询列表
	var list []admin.UserItem
	err = model.Page(page, pageSize).
		Order("user_id desc").
		Scan(&list)

	if err != nil {
		return nil, err
	}

	return &admin.UserListRes{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// Detail 获取用户详情
func (c *cUser) Detail(ctx context.Context, req *admin.UserDetailReq) (res *admin.UserDetailRes, err error) {
	var user admin.UserItem
	err = g.DB().Model("sys_user").
		Where("user_id", req.UserId).
		Where("del_flag", 0).
		Scan(&user)

	if err != nil {
		return nil, err
	}

	// 检查是否查询到数据
	if user.UserId == 0 {
		return nil, gerror.New("用户不存在")
	}

	return &admin.UserDetailRes{
		User: user,
	}, nil
}

// Create 创建用户
func (c *cUser) Create(ctx context.Context, req *admin.UserCreateReq) (res *admin.UserCreateRes, err error) {
	// 检查用户名是否存在
	count, err := g.DB().Model("sys_user").
		Where("user_name", req.UserName).
		Where("del_flag", 0).
		Count()

	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("用户名已存在")
	}

	// 插入用户
	result, err := g.DB().Model("sys_user").Data(g.Map{
		"user_name":   req.UserName,
		"nick_name":   req.NickName,
		"email":       req.Email,
		"phonenumber": req.Phonenumber,
		"sex":         req.Sex,
		"password":    req.Password, // 实际应用中需要加密
		"status":      req.Status,
		"dept_id":     req.DeptId,
		"remark":      req.Remark,
	}).Insert()

	if err != nil {
		return nil, err
	}

	userId, _ := result.LastInsertId()
	return &admin.UserCreateRes{
		UserId: userId,
	}, nil
}

// Update 更新用户
func (c *cUser) Update(ctx context.Context, req *admin.UserUpdateReq) (res *admin.UserUpdateRes, err error) {
	// 更新用户信息
	_, err = g.DB().Model("sys_user").
		Where("user_id", req.UserId).
		Data(g.Map{
			"nick_name":   req.NickName,
			"email":       req.Email,
			"phonenumber": req.Phonenumber,
			"sex":         req.Sex,
			"status":      req.Status,
			"dept_id":     req.DeptId,
			"remark":      req.Remark,
		}).Update()

	return &admin.UserUpdateRes{}, err
}

// Delete 删除用户（逻辑删除）
func (c *cUser) Delete(ctx context.Context, req *admin.UserDeleteReq) (res *admin.UserDeleteRes, err error) {
	_, err = g.DB().Model("sys_user").
		Where("user_id", req.UserId).
		Data(g.Map{"del_flag": 2}).
		Update()

	return &admin.UserDeleteRes{}, err
}
