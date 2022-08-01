package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

const (
	baseUrl string = "https://github.com/trending"
)

type GithubTrending struct {
	c *colly.Collector
}

func (that *GithubTrending) NewCollector() *GithubTrending {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  baseUrl,
		RandomDelay: 500 * time.Microsecond,
		Parallelism: 12,
	})
	that.c = c
	return that
}

func (that *GithubTrending) SpiderGithub() {
	c := that.c
	c.OnRequest(func(r *colly.Request) {
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response", len(r.Body))
	})
	that.getLanguage()
	that.getSince()
	that.getContent()
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("response", response.Body)
		fmt.Println(err.Error())
	})
	c.Visit(baseUrl)
}

func (that *GithubTrending) getLanguage() {
	c := that.c
	c.OnHTML("#select-menu-language .select-menu-list", func(e *colly.HTMLElement) {
		e.DOM.Find(".select-menu-item-text").Each(func(i int, selection *goquery.Selection) {
			fmt.Println("================", strings.TrimSpace(selection.Text()))
		})
	})
}

func (that *GithubTrending) getSince() {
	c := that.c
	c.OnHTML("#select-menu-date .select-menu-list", func(e *colly.HTMLElement) {
		e.DOM.Find(".select-menu-item-text").Each(func(i int, selection *goquery.Selection) {
			fmt.Println("================", strings.TrimSpace(selection.Text()))
		})
	})
}

type ProjectDetailStruct struct {
	Url        string `json:"url",des:"项目链接"`
	Author     string `json:"author",des:"作者"`
	Desc       string `json:"desc",des:"简介"`
	Repo       string `json:"repo",des:"项目仓库"'`
	Starts     int    `json:"starts",des:"目前start数"`
	Forks      int    `json:"forks",des:"目前forks数"`
	Language   int    `json:"language",des:"语言"`
	AddedStars string `json:"added_stars",des:"今天或者这周或者这个月的starts数"`
	Avatars    string `json:"avatars",des:"项目贡献者的头像地址集合"`
}

func (that *GithubTrending) getContent() {
	c := that.c
	c.OnHTML(".Box", func(e *colly.HTMLElement) {
		e.DOM.Find(".Box-row").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				//url, _ := s.Find(".lh-condensed > a").Attr("href")
				//fmt.Println("url=================,", url)
				//repo := strings.TrimLeft(url, "/")
				//fmt.Println("repo=================,", repo)
				//_s := strings.Split(repo, "/")
				//author := _s[0]
				//fmt.Println("author=================,", author)
				//desc := s.Find("p[class='col-9 color-fg-muted my-1 pr-4']").Text()
				//fmt.Println("desc===================,", strings.TrimSpace(desc))
				//language := s.Find(".repo-language-color").Siblings().Text()
				//fmt.Println("language===================,", language)
				//starts := s.Find("a[class='Link--muted d-inline-block mr-3']").First().Text()
				//fmt.Println("starts===================,", strings.TrimSpace(starts))
				//forks := s.Find("a[class='Link--muted d-inline-block mr-3']").Last().Text()
				//fmt.Println("forks===================,", strings.TrimSpace(forks))
				//added_stars := s.Find("span[class='d-inline-block float-sm-right']").Text()
				//fmt.Println("added_stars===================,", strings.TrimSpace(added_stars))
				_avatarsList := make([]string, 0)
				s.Find("img[class='avatar mb-1 avatar-user']").Each(func(i int, s *goquery.Selection) {
					avatar, _b := s.Attr("src")
					if _b {
						_avatarsList = append(_avatarsList, avatar)
					}
				})
				avatars := strings.Join(_avatarsList, ",")
				fmt.Println("avatars===================", avatars)
			}
		})
	})
}
