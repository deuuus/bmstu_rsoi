package server

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Age     int32  `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
	Work    string `json:"work" binding:"required"`
}

type PersonUpdate struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}
