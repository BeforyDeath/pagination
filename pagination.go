package pagination

type Pagination struct {
	total        int
	Limit        int
	visibleRange int
	current      int
	numPages     int
	Pages        Pages
}

type Pages struct {
	First  int
	Last   int
	Active int
	Prev   int
	Next   int
	Page   []int
}

func Create(total, limit, visibleRange int) *Pagination {

	if limit <= 0 {
		limit = 1
	}

	if visibleRange < 4 {
		visibleRange = 4
	}

	p := &Pagination{0, limit, visibleRange, 0, 0, Pages{}}

	p.SetTotal(total)

	return p
}
func (p *Pagination) SetTotal(total int) {
	if total <= 0 {
		total = 1
	}
	p.total = total
	p.setNumPages()
	p.current = 1
}

func (p *Pagination) setNumPages() {
	p.numPages = p.total / p.Limit
	if p.total%p.Limit > 0 {
		p.numPages++
	}
}

func (p *Pagination) Get(current int) Pages {

	if current > p.numPages {
		// todo return error, response 404
		current = p.numPages
	}
	p.current = current

	p.Pages = Pages{}

	firstPage := p.current - (p.visibleRange / 2)
	if firstPage < 1 {
		firstPage = 1
	}

	endPage := firstPage + p.visibleRange
	if endPage > p.numPages {
		endPage = p.numPages + 1
	}

	if firstPage > 1 {
		p.Pages.First = 1
	}

	if p.current > 1 {
		p.Pages.Prev = p.current - 1
	}

	p.Pages.Page = make([]int, 0)
	for i := firstPage; i < endPage; i++ {
		if i == p.current {
			p.Pages.Active = i
		}
		p.Pages.Page = append(p.Pages.Page, i)
	}

	if p.current < p.numPages {
		p.Pages.Next = p.current + 1
	}

	if endPage < p.numPages+1 {
		p.Pages.Last = p.numPages
	}
	return p.Pages
}

func (p *Pagination) GetOffset() (offset int) {
	offset = (p.current - 1) * p.Limit
	return
}
