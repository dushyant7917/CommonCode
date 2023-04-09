// +build ali

package ali

type OSSClient struct {
}

func NewOSSClient() *OSSClient {
	return &OSSClient{}
}

func (c *OSSClient) GetBucketName() string {
	return "foo"
}
