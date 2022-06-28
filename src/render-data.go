package main

type RenderData struct {
	SidebarItems     []MenuItem `json:"sidebar"`
	ProfileDropItems []MenuItem `json:"profile"`
	User             string
}

type DropdownItem struct {
	Link string `json:"link"`
	Text string `json:"text"`
}

type MenuItem struct {
	Type       string         `json:"type"`
	Link       string         `json:"link"`
	Text       string         `json:"text"`
	Label      string         `json:"label"`
	LabelColor string         `json:"labelColor"`
	Items      []DropdownItem `json:"items"`
}
