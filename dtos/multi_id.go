package dtos

type MultiID struct {
	Ids []uint `json:"ids" binding:"required"`
}
