package main

import (
	"bytes"
	"example.com/m/v2/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	baseFolder = "./images/shehui/"
	baseUrl    = "http://www.***.com"
	startUrl   = "http://www.***.com/shehui/"
)

func main() {
	fmt.Println("图片爬虫")
	ImageSpider()
}

func ImageSpider() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "",
		RandomDelay: 500 * time.Microsecond,
		Parallelism: 12,
	})
	c.DetectCharset = false
	//解析详情页
	c.OnHTML("#waterfall .pic", func(e *colly.HTMLElement) {
		e.DOM.Find("li").Each(func(i int, s *goquery.Selection) {
			name, found := s.Find(".picbox img").Attr("alt")
			var str string
			if found {
				var data []byte = []byte(name)
				str = utils.ConvertByte2String(data, "GB18030")
				CreateFolder(str)

			}
			href, found := s.Find(".description a").Attr("href")
			filepath := fmt.Sprintf("%s%s", baseFolder, str)
			if found {
				href = fmt.Sprintf("%s%s", baseUrl, href)
				fmt.Printf("FOUND %v-> %s\n", i, href)
				ViewImageDetail(c, href, filepath)
			}

		})
	})

	//寻找下一页
	c.OnHTML(".itempages>ul", func(e *colly.HTMLElement) {
		href, found := e.DOM.Find("li:nth-last-child(3)>a").Attr("href")
		if found {
			href = fmt.Sprintf("%s%s", startUrl, href)
			e.Request.Visit(e.Request.AbsoluteURL(href))
			fmt.Printf("当前页码 -> %s\n", href)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Host", "www.je678.com")
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response", len(r.Body))
	})
	c.Visit(startUrl)
}

// ViewImageDetail 同一文件夹下图片处理
func ViewImageDetail(c *colly.Collector, src string, filepath string) {
	detailC := c.Clone()
	detailC.OnHTML("div.picture_txt ", func(e *colly.HTMLElement) {
		src, found := e.DOM.Find("img").Attr("src")
		if found {
			src = fmt.Sprintf("%s%s", baseUrl, src)
			fmt.Println("下载图片地址 -> ", src)
			DownImage(c, src, filepath)
		}
	})
	//寻找下一页
	detailC.OnHTML(".itempages2>ul li:last-child>a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "#" {
			href := fmt.Sprintf("%s%s", startUrl, href)
			e.Request.Visit(href)
		}

	})
	detailC.Visit(src)
}

// DownImage 下载图片
func DownImage(c *colly.Collector, src string, filepath string) {
	imgC := c.Clone()
	imgC.OnResponse(func(r *colly.Response) {
		url := fmt.Sprintf("%s", r.Request.URL)
		str := strings.Split(url, "/")
		fileName := str[len(str)-1]
		reader := bytes.NewReader(r.Body)
		body, _ := ioutil.ReadAll(reader)
		filepath = StrReplace(filepath)
		fileTruePath := fmt.Sprintf("%s/%s", filepath, fileName)
		err := ioutil.WriteFile(fileTruePath, body, 0755)
		if err != nil {
			fmt.Println("文件保存失败")
		}
		fmt.Println("文件保存成功:", fileTruePath)
	})
	imgC.Visit(src)
}

// CreateFolder 创建文件夹
func CreateFolder(string2 string) {
	string2 = StrReplace(string2)
	path := fmt.Sprintf("%s%s", baseFolder, string2)
	_b, err := utils.PathExists(path)
	if err != nil {
		fmt.Println("文件夹判断：", err)
		return
	}
	if !_b {
		os.MkdirAll(path, 0777)
	}
}

func StrReplace(str string) string {
	re3, _ := regexp.Compile("<b>")
	re4, _ := regexp.Compile("</b>")
	rep2 := re3.ReplaceAllString(str, "")
	rep2 = re4.ReplaceAllString(rep2, "")
	return rep2
}
