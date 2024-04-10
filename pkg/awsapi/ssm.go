package awsapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

// SSMIface : -
type SSMIface interface {
	GetSSMParameter(key string) (string, error)
	GetSSMParameters(keys []string) (map[string]string, error)
}

// SSMInstance : ssm instance
type SSMInstance struct {
	client ssmiface.SSMAPI
}

// NewSSMClient : ssm client 生成
func NewSSMClient(client ssmiface.SSMAPI) SSMIface {
	return &SSMInstance{
		client: client,
	}
}

// GetSSMParameters : SSM パラメータキーを指定し復号し値を返す
func (d *SSMInstance) GetSSMParameter(key string) (string, error) {
	r, err := d.client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}
	return *r.Parameter.Value, nil
}

func (d *SSMInstance) GetSSMParameters(keys []string) (map[string]string, error) {
	names := []*string{}
	for _, k := range keys {
		names = append(names, aws.String(k))
	}

	ssmParameters := &ssm.GetParametersInput{
		Names:          names,
		WithDecryption: aws.Bool(true),
	}

	r, err := d.client.GetParameters(ssmParameters)
	if err != nil {
		return nil, err
	}

	s := make(map[string]string)
	for _, p := range r.Parameters {
		s[*p.Name] = *p.Value
	}
	return s, nil
}
