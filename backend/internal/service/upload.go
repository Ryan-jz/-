package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// IUpload 文件上传服务接口
type IUpload interface {
	UploadImage(ctx context.Context, file *ghttp.UploadFile) (string, error)
	DeleteFile(ctx context.Context, filePath string) error
}

type uploadImpl struct{}

var localUpload IUpload

func init() {
	RegisterUpload(&uploadImpl{})
}

// Upload 获取上传服务实例
func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

// RegisterUpload 注册上传服务实现
func RegisterUpload(i IUpload) {
	localUpload = i
}

// UploadImage 上传图片文件
func (s *uploadImpl) UploadImage(ctx context.Context, file *ghttp.UploadFile) (string, error) {
	// 验证文件类型
	if !s.isImageFile(file.Filename) {
		return "", gerror.New("只支持上传图片文件（jpg、jpeg、png、gif、webp）")
	}

	// 验证文件大小（限制 5MB）
	if file.Size > 5*1024*1024 {
		return "", gerror.New("图片大小不能超过 5MB")
	}

	// 生成文件保存路径
	savePath, err := s.generateSavePath(file.Filename)
	if err != nil {
		g.Log().Errorf(ctx, "生成保存路径失败: %v", err)
		return "", gerror.New("生成保存路径失败")
	}

	// 确保目录存在
	dir := filepath.Dir(savePath)
	if !gfile.Exists(dir) {
		if err := gfile.Mkdir(dir); err != nil {
			g.Log().Errorf(ctx, "创建目录失败: %v", err)
			return "", gerror.New("创建目录失败")
		}
	}

	// 保存文件
	if _, err := file.Save(savePath, true); err != nil {
		g.Log().Errorf(ctx, "保存文件失败: %v", err)
		return "", gerror.New("保存文件失败")
	}

	// 返回访问路径（相对路径）
	relativePath := strings.Replace(savePath, "public/", "/", 1)
	relativePath = strings.ReplaceAll(relativePath, "\\", "/")

	g.Log().Infof(ctx, "文件上传成功: %s", relativePath)
	return relativePath, nil
}

// DeleteFile 删除文件
func (s *uploadImpl) DeleteFile(ctx context.Context, filePath string) error {
	// 转换为实际文件路径
	actualPath := "public" + filePath
	actualPath = strings.ReplaceAll(actualPath, "/", string(filepath.Separator))

	if !gfile.Exists(actualPath) {
		return gerror.New("文件不存在")
	}

	if err := gfile.Remove(actualPath); err != nil {
		g.Log().Errorf(ctx, "删除文件失败: %v", err)
		return gerror.New("删除文件失败")
	}

	g.Log().Infof(ctx, "文件删除成功: %s", filePath)
	return nil
}

// isImageFile 判断是否为图片文件
func (s *uploadImpl) isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// generateSavePath 生成文件保存路径
func (s *uploadImpl) generateSavePath(filename string) (string, error) {
	// 获取文件扩展名
	ext := filepath.Ext(filename)

	// 生成唯一文件名（使用时间戳 + MD5）
	timestamp := time.Now().Format("20060102150405")
	hash, err := gmd5.Encrypt(filename + timestamp)
	if err != nil {
		return "", err
	}

	// 按日期分目录存储
	dateDir := time.Now().Format("2006/01/02")
	newFilename := fmt.Sprintf("%s_%s%s", timestamp, gstr.SubStr(hash, 0, 8), ext)

	// 完整路径：public/uploads/images/2024/01/02/20240102150405_abc12345.jpg
	savePath := filepath.Join("public", "uploads", "images", dateDir, newFilename)

	return savePath, nil
}
