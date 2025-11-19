package models

type MsgListReq struct {
	TargetId uint64 `json:"target_id" binding:"required"`
	AfterId  uint64 `json:"after_id" `
	Limit    int    `json:"limit"    binding:"omitempty,min=1,max=100"`
}
