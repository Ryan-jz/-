package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	targetURL := "https://www.bad-reichenhaller.de/de/nachhaltigkeit.html#recycling-tipps"
	outputDir := "./downloaded_images"

	// 创建输出目录
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}

	fmt.Printf("开始爬取: %s\n", targetURL)

	// 获取网页内容
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("HTTP 状态码错误: %d\n", resp.StatusCode)
		return
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("解析 HTML 失败: %v\n", err)
		return
	}

	baseURL, _ := url.Parse(targetURL)
	imageCount := 0

	// 查找所有图片标签
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgSrc, exists := s.Attr("src")
		if !exists {
			return
		}

		// 处理相对路径
		imgURL, err := url.Parse(imgSrc)
		if err != nil {
			fmt.Printf("解析图片 URL 失败: %s\n", imgSrc)
			return
		}

		// 转换为绝对路径
		absoluteURL := baseURL.ResolveReference(imgURL).String()

		// 下载图片
		if err := downloadImage(absoluteURL, outputDir); err != nil {
			fmt.Printf("下载失败 [%s]: %v\n", absoluteURL, err)
		} else {
			imageCount++
			fmt.Printf("已下载 [%d]: %s\n", imageCount, absoluteURL)
		}
	})

	fmt.Printf("\n完成！共下载 %d 张图片到 %s\n", imageCount, outputDir)
}

// downloadImage 下载图片到指定目录
func downloadImage(imageURL, outputDir string) error {
	// 发送请求
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP 状态码: %d", resp.StatusCode)
	}

	// 生成文件名
	parsedURL, _ := url.Parse(imageURL)
	filename := path.Base(parsedURL.Path)

	// 如果文件名为空或只是 "/"，使用默认名称
	if filename == "" || filename == "/" {
		filename = fmt.Sprintf("image_%d.jpg", len(imageURL))
	}

	// 清理文件名中的特殊字符
	filename = strings.ReplaceAll(filename, "?", "_")
	filename = strings.ReplaceAll(filename, "&", "_")

	filepath := filepath.Join(outputDir, filename)

	// 创建文件
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件
	_, err = io.Copy(file, resp.Body)
	return err
}
