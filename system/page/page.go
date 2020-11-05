package page

import (
	"html/template"
	"math"
	"strconv"
)

type page struct {
	Pre      string //上一页样式
	Next     string //下一页样式
	Count    int    //总数
	NowPage  int    //当前页
	PageSize int    //一页条数
	Current  int    //当前页
	PrePage  int    //上一页
	NextPage int    //下一页
	MaxNum   int    //页数按钮显示多少个
}

func NewPage() *page {
	return &page{
		Pre:    "«",
		Next:   "»",
		MaxNum: 10,
	}
}

func (p *page) Pagination(nowPage, pageSize, count int, param string) template.HTML {
	p.NowPage = nowPage
	p.Count = count
	if nowPage <= 1 {
		nowPage = 1
	}
	//总页数
	num := float64(count) / float64(pageSize)
	pageCount := int(math.Ceil(num))

	if count <= 1 || pageCount <= 1 {
		return template.HTML(`<ul class="pagination"><li class="disabled"><span>` + p.Pre + `</span></li><li class="active"><span>1</span></li><li class="disabled"><span>` + p.Next + `</span></li></ul>`)
	}

	p.Current = nowPage //当前页
	p.Count = pageCount //总页数

	if nowPage <= 1 {
		p.PrePage = 1 //上一页
	} else {
		p.PrePage = nowPage - 1 //上一页
	}

	if nowPage >= pageCount {
		p.NextPage = pageCount //下一页
	} else {
		p.NextPage = nowPage + 1 //下一页
	}

	var htmls string
	//拼接上面的
	htmls = `<ul class="pagination">`
	if p.Current <= 1 {
		htmls += `<li class="disabled"><span>` + p.Pre + `</span></li>`
	} else {
		htmls += `<li><a href="?p=` + strconv.Itoa(p.PrePage) + `">` + p.Pre + `</a></li>`
	}

	htmls += middleStr(p.Current, pageCount, p.MaxNum, param)

	//拼接底页
	if p.Current >= pageCount {
		htmls += `<li class="disabled"><span>` + p.Next + `</span></li>`
	} else {
		htmls += `<li><a href="?p=` + strconv.Itoa(p.NextPage) + `">` + p.Next + `</a></li>`
	}

	htmls += `</ul>`

	return template.HTML(htmls)
}

//中间部分拼接
func middleStr(current, pageCount, maxCount int, param string) string {
	var htmls string
	var i int
	//如果总页数不是很多就全部展示出来
	if pageCount <= maxCount {
		for i = 1; i <= pageCount; i++ {
			if current == i {
				htmls += `<li class="active"><span>` + strconv.Itoa(i) + `</span></li>`
			} else {
				if param != "" {
					htmls += `<li><a href="?p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				} else {
					htmls += `<li><a href="?` + param + `&p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
				}

			}
		}
		return htmls
	}
	//如果分页比较多就展示部分按钮，其他部分用...代替
	if current > 1 {
		htmls += `<li><a href="?p=1">...</a></li>`
	}
	//获取最大值
	maxCount -= 1
	now := maxCount + current
	if now >= pageCount {
		now = pageCount
	}
	//里面必须有特定个数的按钮
	cnum := now - maxCount
	for i = cnum; i <= now; i++ {
		if current == i {
			htmls += `<li class="active"><span>` + strconv.Itoa(i) + `</span></li>`
		} else {
			if param != "" {
				htmls += `<li><a href="?p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
			} else {
				htmls += `<li><a href="?` + param + `&p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
			}
		}
	}
	//...结尾
	if now < pageCount && current >= 1 {
		htmls += `<li><a href="?p=` + strconv.Itoa(now) + `">...</a></li>`
	}
	return htmls
}