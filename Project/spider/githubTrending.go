package spider

import (
	"example.com/m/v2/model/db"
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
	c            *colly.Collector
	SinceType    int    `json:"since_type"`
	LanguageType int    `json:"language_type"`
	BaseUrl      string `json:"base_url"`
}

type SearchStruct struct {
	Since        string `json:"since"`
	SinceType    int    `json:"since_type"`
	Language     string `json:"language"`
	LanguageType int    `json:"language_type"`
}

var _trendingList = make([]*db.GithubTrending, 0)

func (that *GithubTrending) NewCollector(params *SearchStruct) *GithubTrending {
	_baseUrl := baseUrl + "/" + params.Language + "?since=" + params.Since
	that.LanguageType = params.LanguageType
	that.SinceType = params.SinceType
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  _baseUrl,
		RandomDelay: 500 * time.Microsecond,
		Parallelism: 12,
	})
	that.BaseUrl = _baseUrl
	that.c = c
	return that
}

func (that *GithubTrending) SpiderGithub() (data interface{}, err string) {
	c := that.c
	c.OnRequest(func(r *colly.Request) {
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response", len(r.Body))
	})
	var _err string
	//that.getLanguage()
	//that.getSince()
	that.getContent()
	c.OnScraped(func(response *colly.Response) {
		if len(_trendingList) > 0 {
			_data, _err := db.NewGithubTrending().Add(&_trendingList)
			fmt.Println("=============数据库插入数据===========_data", _data)
			fmt.Println("========================_err", _err)
			_trendingList = make([]*db.GithubTrending, 0)
		}
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("response1233123", response.StatusCode)
		fmt.Println("response1233123", len(response.Body))
		fmt.Println(err.Error())
		_err = err.Error()

	})
	c.Visit(that.BaseUrl)
	if _err != "" {
		return nil, _err
	} else {
		return 1, ""
	}
}

//type LanguageStruct struct {
//	Name string `json:"name"`
//	Type int    `json:"type"`
//}

func (that *GithubTrending) getLanguage() {
	c := that.c
	_languageList := make([]*db.GithubLanguage, 0)
	_type := 1
	c.OnHTML("#select-menu-language .select-menu-list", func(e *colly.HTMLElement) {
		e.DOM.Find(".select-menu-item-text").Each(func(i int, selection *goquery.Selection) {
			_item := new(db.GithubLanguage)
			_item.Name = strings.TrimSpace(selection.Text())
			_item.Type = _type
			_item.IsEnable = 2
			_type += 1
			_languageList = append(_languageList, _item)
		})
		_, _ = db.NewGithubLanguage().Add(&_languageList)
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

func (that *GithubTrending) getContent() {
	c := that.c
	c.OnHTML(".Box", func(e *colly.HTMLElement) {
		e.DOM.Find(".Box-row").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Find(".lh-condensed > a").Attr("href")
			repo := strings.TrimLeft(url, "/")
			_s := strings.Split(repo, "/")
			author := _s[0]
			desc := s.Find("p[class='col-9 color-fg-muted my-1 pr-4']").Text()
			language := s.Find(".repo-language-color").Siblings().Text()
			starts := s.Find("a[class='Link--muted d-inline-block mr-3']").First().Text()
			forks := s.Find("a[class='Link--muted d-inline-block mr-3']").Last().Text()
			added_stars := s.Find("span[class='d-inline-block float-sm-right']").Text()
			_avatarsList := make([]string, 0)
			s.Find("img[class='avatar mb-1 avatar-user']").Each(func(i int, s *goquery.Selection) {
				avatar, _b := s.Attr("src")
				if _b {
					_avatarsList = append(_avatarsList, avatar)
				}
			})
			avatars := strings.Join(_avatarsList, ",")
			_item := new(db.GithubTrending)
			_item.Author = author
			_item.Repo = repo
			_item.Url = url
			_item.Desc = desc
			_item.Starts = strings.TrimSpace(starts)
			_item.Forks = strings.TrimSpace(forks)
			_item.Language = strings.TrimSpace(language)
			_item.AddedStars = strings.TrimSpace(added_stars)
			_item.Avatars = strings.TrimSpace(avatars)
			_item.LanguageType = that.LanguageType
			_item.SinceType = that.SinceType
			_trendingList = append(_trendingList, _item)
		})
	})
}
