package owl

type Menu struct {
	Name  string  `json:"name"`
	Url   string  `json:"url"`
	Sort  int     `json:"sort"`
	Icon  string  `json:"icon"`
	Meta  string  `json:"meta"`
	Path  string  `json:"path"`
	Child []*Menu `json:"child"`
}

type MenuManager struct {
	menus []*Menu
}

func (m *MenuManager) AddMenu(menus ...*Menu) {
	m.menus = append(m.menus, menus...)
}

func (m *MenuManager) GetMenus() []*Menu {
	return m.menus
}
