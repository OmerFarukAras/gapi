package models

type RequestLOG struct {
	CID  string `json:"id"`
	Data string `json:"data"`
	Time string `json:"time"`
}

func (c RequestLOG) ID() (jsonField string, value interface{}) {
	value = c.CID
	jsonField = "id"
	return
}
