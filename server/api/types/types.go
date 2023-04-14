package types

type Result struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type IDUri struct {
	ID int32 `uri:"id"`
}
