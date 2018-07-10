package aws

const (
	// AwsAdminArnKey is the key in the Secret under which the ARN for the admin role is held.
	AdminArnKey = "aws.admin.arn"
	// AwsAwsoperatorArnKey is the key in the Secret under which the ARN for the aws-operator role is held.
	AwsoperatorArnKey = "aws.awsoperator.arn"
)

type Request struct {
	AdminARN       string
	AwsOperatorARN string
}
