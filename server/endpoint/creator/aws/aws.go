package aws

type Roles struct {
	Admin       string `json:"admin"`
	AwsOperator string `json:"awsoperator"`
}

type AWS struct {
	Roles Roles `json:"roles"`
}
