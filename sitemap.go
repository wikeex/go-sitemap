package goSitemap

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Url struct {
	XmlName xml.Name		`xml:"url"`
	Loc string				`xml:"loc"`
	ChangeFreq string		`xml:"changefreq"`
	Priority string			`xml:"priority"`
	LastMod string			`xml:"lastmod"`
}

type UrlSet struct {
	XmlName xml.Name		`xml:"urlset"`
	Xmlns string			`xml:"xmlns,addr"`
	Urls []Url
}

func GenerateSiteMap(filePath string, urlSet *UrlSet) error {
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	output, err := xml.MarshalIndent(urlSet, "  ", "    ")
	if err != nil {
		return err
	}
	xmlData := append([]byte(xml.Header), output...)
	err = ioutil.WriteFile(filePath, xmlData, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
