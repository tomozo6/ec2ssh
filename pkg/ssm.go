package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMDescribeInstancesAPI interface {
	DescribeInstanceInformation(ctx context.Context, params *ssm.DescribeInstanceInformationInput, optFns ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error)
}

func getSSMInstances(c context.Context, api SSMDescribeInstancesAPI, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
	return api.DescribeInstanceInformation(c, input)
}
