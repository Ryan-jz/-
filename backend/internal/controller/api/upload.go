package api

import (
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadController 文件上传控制器
var Upload = cUpload{}

type cUpload struct{}

// Image 上传图片
func (c *cUpload) Image(r *ghttp.Request) {
	// 获取上传的文件
	file := r.GetUploadFile("file")
	if file == nil {
		r.Response.WriteJson(g.Map{
			"code":    1,
			"message": "请选择要上传的文件",
		})
		return
	}

	// 调用上传服务
	url, err := service.Upload().UploadImage(r.Context(), file)
	if err != nil {
		g.Log().Error(r.Context(), "图片上传失败:", err)
		r.Response.WriteJson(g.Map{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	g.Log().Infof(r.Context(), "图片上传成功: %s", url)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "上传成功",
		"data": g.Map{
			"url": url,
		},
	})
}

// Delete 删除文件
func (c *cUpload) Delete(r *ghttp.Request) {
	var req struct {
		Url string `json:"url" v:"required#文件路径不能为空"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	// 调用删除服务
	if err := service.Upload().DeleteFile(r.Context(), req.Url); err != nil {
		g.Log().Error(r.Context(), "文件删除失败:", err)
		r.Response.WriteJson(g.Map{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	g.Log().Infof(r.Context(), "文件删除成功: %s", req.Url)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "删除成功",
	})
}
