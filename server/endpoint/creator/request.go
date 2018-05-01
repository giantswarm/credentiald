package creator

type Roles struct {
	Admin       string `json:"admin"`
	AwsOperator string `json:"awsoperator"`
}

type AWS struct {
	Roles Roles `json:"roles"`
}

type Request struct {
	Provider string `json:"provider"`
	AWS      AWS    `json:"aws"`

	Organization string `json:"-"`
}
