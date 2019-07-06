package main

type sendOut struct {
	TeamID string `json:"team_id"`
	ItemID string `json:"item_id"`
}
type sendCampdocOut struct {
	Title    string `json:"Title"`
	Note     string `json:"Note"`
	LinkName string `json:"linkName"`
	Link     string `json:"Link"`
	ImgLink  string `json:"ImgLink"`
}
