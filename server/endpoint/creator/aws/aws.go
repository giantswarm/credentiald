package aws

type AWSRoles struct {
	Admin       string `json:"admin"`
	AwsOperator string `json:"awsoperator"`
}

type AWS struct {
	Roles AWSRoles `json:"roles"`
}
