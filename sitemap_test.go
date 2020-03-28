package goSitemap

import (
	"fmt"
	"testing"
)

func TestGenerateSiteMap(t *testing.T) {
	urlSet := UrlSet{Urls:make([]Url, 2, 2)}

	for i := 0; i < len(urlSet.Urls); i++ {
		urlSet.Urls[i].Priority = "1.0"
		urlSet.Urls[i].Loc = "http://www.baidu.com"
		urlSet.Urls[i].ChangeFreq = "Weekly"
		urlSet.Urls[i].LastMod = "2020-03-29"
	}

	err := GenerateSiteMap("sitemap.xml", &urlSet)
	if err != nil {
		fmt.Println(err)
	}
}
