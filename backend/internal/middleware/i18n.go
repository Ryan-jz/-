package middleware

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// I18n 国际化中间件
// 从请求中提取语言信息并存入上下文
func I18n(r *ghttp.Request) {
	ctx := r.Context()
	locale := getLocaleFromRequest(r, ctx)

	// 验证语言是否支持
	if !isSupportedLocale(ctx, locale) {
		locale = getDefaultLocale(ctx)
	}

	// 存入上下文
	r.SetCtx(context.WithValue(ctx, "locale", locale))

	// 设置响应头
	r.Response.Header().Set("Content-Language", locale)

	r.Middleware.Next()
}

// getLocaleFromRequest 从请求中获取语言设置
// 优先级：URL参数 > Header > Cookie > 默认语言
func getLocaleFromRequest(r *ghttp.Request, ctx context.Context) string {
	// 1. 从URL参数获取 ?lang=zh-CN
	locale := r.Get("lang").String()
	if locale != "" {
		return normalizeLocale(locale)
	}

	// 2. 从Header获取 Accept-Language
	locale = r.Header.Get("Accept-Language")
	if locale != "" {
		// Accept-Language 可能是 "zh-CN,zh;q=0.9,en;q=0.8"
		// 取第一个
		if strings.Contains(locale, ",") {
			locale = strings.Split(locale, ",")[0]
		}
		if strings.Contains(locale, ";") {
			locale = strings.Split(locale, ";")[0]
		}
		return normalizeLocale(locale)
	}

	// 4. 返回默认语言
	return getDefaultLocale(ctx)
}

// normalizeLocale 标准化语言代码
// 将 zh, zh-cn, zh_CN 等统一为 zh-CN
func normalizeLocale(locale string) string {
	locale = strings.TrimSpace(locale)
	locale = strings.ToLower(locale)

	// 处理常见格式
	locale = strings.ReplaceAll(locale, "_", "-")

	// 标准化常见语言代码
	switch {
	case strings.HasPrefix(locale, "zh"):
		return "zh-CN"
	case strings.HasPrefix(locale, "en"):
		return "en-US"
	case strings.HasPrefix(locale, "de"):
		return "de-DE"
	default:
		return locale
	}
}

// isSupportedLocale 检查语言是否支持
func isSupportedLocale(ctx context.Context, locale string) bool {
	supported := g.Cfg().MustGet(ctx, "i18n.supported").Strings()
	if len(supported) == 0 {
		// 如果配置为空，使用默认支持列表
		supported = []string{"zh-CN", "en-US", "de-DE"}
	}
	return gstr.InArray(supported, locale)
}

// getDefaultLocale 获取默认语言
func getDefaultLocale(ctx context.Context) string {
	defaultLocale := g.Cfg().MustGet(ctx, "i18n.default").String()
	if defaultLocale == "" {
		return "zh-CN"
	}
	return defaultLocale
}

// GetLocale 从上下文获取当前语言
func GetLocale(ctx context.Context) string {
	locale := ctx.Value("locale")
	if locale == nil {
		return getDefaultLocale(ctx)
	}
	return locale.(string)
}
