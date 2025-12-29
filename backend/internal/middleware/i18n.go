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

	if !isSupportedLocale(ctx, locale) {
		locale = getDefaultLocale(ctx)
	}

	r.SetCtx(context.WithValue(ctx, "locale", locale))
	r.Response.Header().Set("Content-Language", locale)
	r.Middleware.Next()
}

func getLocaleFromRequest(r *ghttp.Request, ctx context.Context) string {
	locale := r.Get("lang").String()
	if locale != "" {
		return normalizeLocale(locale)
	}

	locale = r.Header.Get("Accept-Language")
	if locale != "" {
		if strings.Contains(locale, ",") {
			locale = strings.Split(locale, ",")[0]
		}
		if strings.Contains(locale, ";") {
			locale = strings.Split(locale, ";")[0]
		}
		return normalizeLocale(locale)
	}

	return getDefaultLocale(ctx)
}

func normalizeLocale(locale string) string {
	locale = strings.TrimSpace(locale)
	locale = strings.ToLower(locale)
	locale = strings.ReplaceAll(locale, "_", "-")

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

func isSupportedLocale(ctx context.Context, locale string) bool {
	supported := g.Cfg().MustGet(ctx, "i18n.supported").Strings()
	if len(supported) == 0 {
		supported = []string{"zh-CN", "en-US", "de-DE"}
	}
	return gstr.InArray(supported, locale)
}

func getDefaultLocale(ctx context.Context) string {
	defaultLocale := g.Cfg().MustGet(ctx, "i18n.default").String()
	if defaultLocale == "" {
		return "zh-CN"
	}
	return defaultLocale
}

func GetLocale(ctx context.Context) string {
	locale := ctx.Value("locale")
	if locale == nil {
		return getDefaultLocale(ctx)
	}
	return locale.(string)
}
