package types

type Interpretation struct {
	ID		int 	`id`
	Text 	string	`text`
}

type Vectors struct {
	S	S `json:"S"`
}

type S struct {
	h string `h`
	s string `s`
}