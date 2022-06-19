package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	DouBanMovie()
}

func DouBanMovie() {
	startUrl := "https://movie.douban.com/top250"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "",
		RandomDelay: 500 * time.Microsecond,
		Parallelism: 12,
	})
	//解析列表页
	c.OnHTML(".grid_view", func(e *colly.HTMLElement) {
		e.DOM.Find("li").Each(func(i int, s *goquery.Selection) {
			href, found := s.Find("div.hd > a").Attr("href")
			if found {
				fmt.Printf("FOUND %v-> %s\n", i, href)
				ParseDetail(c, href)
			}
			if len(movieList) == 10 {
				SaveToCsv(movieList)
				movieList = []MovieDetailStruct{}
			}
		})
	})
	// 查找下一页
	c.OnHTML("div.paginator > span.next", func(element *colly.HTMLElement) {
		href, found := element.DOM.Find("a").Attr("href")
		if found {
			element.Request.Visit(element.Request.AbsoluteURL(href))
		}
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Host", "movie.douban.com")
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response", len(r.Body))
	})
	c.Visit(startUrl)
}

type MovieDetailStruct struct {
	Title    string `json:"title",des:"电影名称"`
	Time     string `json:"time'",des:"上映时间"`
	Duration string `json:"duration",des:"播放时长"`
	Director string `json:"director",des:"导演"`
	FilmType string `json:"filmType",des:"类型"`
	Address  string `json:"address",des:"制片国家/地区"`
	Language string `json:"language",des:"语言"`
	Des      string `json:"des",des:"描述"`
}

var movieList []MovieDetailStruct

func ParseDetail(collector *colly.Collector, href string) {
	c := collector.Clone()
	//解析详情页面
	c.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.DOM.Find("#content>h1 span:first-child").Text()
		re := regexp.MustCompile("\\((.*?)\\)")
		timeTemp := e.DOM.Find("#content>h1 span:last-child").Text()
		time := re.FindStringSubmatch(timeTemp)
		duration := e.DOM.Find("#content #info span[property='v:runtime']").Text()
		director := e.DOM.Find("#content #info .attrs a[rel='v:directedBy']").Text()
		var _typeList []string
		e.DOM.Find("#content #info span[property='v:genre']").Each(func(i int, s *goquery.Selection) {
			_typeList = append(_typeList, s.Text())
		})
		filmType := strings.Join(_typeList, ",")
		des := e.DOM.Find("#link-report span[property='v:summary']").Text()
		re = regexp.MustCompile(`[\s\p{Zs}]{1,}`)
		des = re.ReplaceAllString(des, "")
		text := e.DOM.Find("#content #info ").Text()
		re = regexp.MustCompile("语言:(.*)")
		_language := re.FindStringSubmatch(text)
		re2 := regexp.MustCompile("制片国家/地区:(.*)")
		_address := re2.FindStringSubmatch(text)
		var language string
		if len(_language) > 1 {
			language = _language[1]
		} else {
			language = ""
		}
		var address string
		if len(_address) > 1 {
			address = _address[1]
		} else {
			address = ""
		}
		movie := MovieDetailStruct{
			Title:    title,
			Time:     time[1],
			Duration: duration,
			Director: director,
			FilmType: filmType,
			Address:  address,
			Language: language,
			Des:      des,
		}
		movieList = append(movieList, movie)
		fmt.Println("数组长度", len(movieList))
	})
	//错误响应
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})
	c.Visit(href)
}

func SaveToCsv(movieLists []MovieDetailStruct) {
	_b, _err := PathExists("./doubanlist.csv")
	if _err != nil {
		fmt.Println("文件检查出错", _err)
	}
	if !_b {
		csvFile, err := os.Create("doubanlist.csv")
		if err != nil {
			panic(err)
		}
		defer csvFile.Close()
		//	文件头
		MovieDetailStructs := MovieDetailStruct{
			Title:    "电影名称",
			Time:     "上映时间",
			Duration: "播放时长",
			Director: "导演",
			FilmType: "类型",
			Address:  "制片国家/地区",
			Language: "语言",
			Des:      "描述",
		}
		writer := csv.NewWriter(csvFile)
		err = writer.Write(toString(MovieDetailStructs))
		if err != nil {
			panic(err)
		}
		writer.Flush()
	}
	if _b {
		file, err := os.OpenFile("doubanlist.csv", os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("文件打开错误", err)
			panic(err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		for _, movie := range movieLists {
			line := toString(movie)
			err := writer.Write(line)
			if err != nil {
				panic(err)
			}
		}
		writer.Flush()
	}
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//将结构体转换成字符串
func toString(c MovieDetailStruct) []string {
	line := []string{
		c.Title,
		c.Time,
		c.Duration,
		c.Director,
		c.FilmType,
		c.Address,
		c.Language,
		c.Des,
	}
	return line
}
