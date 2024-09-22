package owl

import "github.com/jinzhu/copier"

const (
	MenuTypeMenu = "菜单"
	MenuTypeBtn  = "按钮"
)

type Meta struct {
	Title string `json:"title"` // 菜单标题
	Icon  string `json:"icon"`  // 菜单图标
}

type Menu struct {
	Path       string   `json:"path"`                  // 前端路由地址
	Name       string   `json:"name"`                  // 前端路由名称，组件名称
	ParentName string   `json:"parentName"`            // 父级菜单名称
	Ancestors  string   `gorm:"comment:祖先;" json:"id"` // 祖先菜单，逗号分割，唯一键
	Rank       int      `json:"rank,omitempty"`        // 菜单排序
	Meta       Meta     `json:"meta"`                  // 菜单 meta 信息，用于前端显示
	MenuType   string   `json:"menuType"`              // 菜单类型，菜单，按钮
	Apis       []string `json:"apis"`                  // 此动作需要拥有的api访问权限，如果是按钮，可以设置此字段
	Children   []*Menu  `json:"children,omitempty"`
}

// Clone 复制菜单
func (i *Menu) Clone() *Menu {
	if i == nil {
		return nil
	}

	// 复制当前节点
	var cloneMenu Menu
	_ = copier.Copy(&cloneMenu, i)
	// 递归复制子节点
	if i.Children != nil {
		clonedChildren := make([]*Menu, len(i.Children))
		for i, child := range i.Children {
			clonedChildren[i] = child.Clone()
		}
		cloneMenu.Children = clonedChildren
	}

	return &cloneMenu
}

type MenuManager struct {
	menus []*Menu
}

func (m *MenuManager) AddMenu(menus ...*Menu) {
	m.menus = append(m.menus, menus...)
}

func (m *MenuManager) GetMenus() []*Menu {
	var clonedMenus []*Menu

	for _, menu := range m.menus {
		clonedMenus = append(clonedMenus, menu.Clone())
	}
	return clonedMenus
}
