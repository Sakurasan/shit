package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "log"
	// "net/http"
	"strconv"
)

type One struct {
	Id    int64
	Vol   int64
	Time  string
	Day   string
	Month string
	Year  string
	Img   string
	Word  string

	Atitle  string
	Aurl    string
	Note    string
	Author  string
	Article string

	Qtitle string
	Qurl   string
	Qdesc  string
	Answer string
}

var (
	root_url = "http://wufazhuce.com"
	vol      string
	time     string
	img      string
	word     string
	aTitle   string
	aUrl     string
)

func One2three() One {
	doc, err := goquery.NewDocument(root_url)
	if err != nil {
		fmt.Println(err)
	}
	a := One{}

	date := doc.Find("div[class=\"item active\"]")
	a.Img, _ = date.Find("img.fp-one-imagen").Attr("src")

	str := date.Find("p.titulo").Text()
	vol := SubString(str, 4, 4)
	a.Vol, _ = strconv.ParseInt(vol, 10, 64)
	a.Day = date.Find("p.dom").Text()
	a.Time = date.Find("p.may").Text()
	// fmt.Println(reflect.TypeOf(a.Vol))
	a.Word = date.Find("div.fp-one-cita").Find("a").Text()

	Two := doc.Find("div.corriente")
	a.Aurl, _ = Two.Find("p.one-articulo-titulo").Find("a").Attr("href")
	// a.Atitle = Two.Find("p.one-articulo-titulo").Find("a").Text()

	Three := doc.Find("div.fp-one-cuestion").Find("div.corriente")
	a.Qtitle = Three.Find("p.one-cuestion-titulo").Text()
	a.Qurl, _ = Three.Find("p.one-cuestion-titulo").Find("a").Attr("href")

	data2, err := goquery.NewDocument(a.Aurl)
	if err != nil {
		fmt.Println(err)
	}
	art := data2.Find("div.one-articulo")
	a.Atitle = art.Find("h2.articulo-titulo").Text()
	a.Note = art.Find("div.articulo-principal").Text()
	a.Author = art.Find("p.articulo-autor").Text()
	a.Article, _ = art.Find("div.articulo-contenido").Html()

	data3, err := goquery.NewDocument(a.Qurl)
	if err != nil {
		fmt.Println(err)
	}
	QnA := data3.Find("div.one-cuestion")
	a.Qdesc = QnA.Find("div.cuestion-contenido").Eq(0).Text()
	a.Answer, _ = QnA.Find("div.cuestion-contenido").Eq(1).Html()

	// fmt.Println(a)
	return a

}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
