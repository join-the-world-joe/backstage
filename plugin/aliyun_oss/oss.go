package aliyun_oss

import (
	"backstage/abstract/oss"
	oss2 "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type _oss struct {
	opts     *Options
	id       string
	secret   string
	endpoint string
	client   *oss2.Client
}

func NewOSS(opts ...Option) (oss.OSS, error) {
	var id, secret, endpoint string
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if len(options.accessKeyId) > 0 {
		id = options.accessKeyId
	}

	if len(options.accessKeySecret) > 0 {
		secret = options.accessKeySecret
	}

	if len(options.endpoint) > 0 {
		endpoint = options.endpoint
	}

	client, err := oss2.New(endpoint, id, secret)
	if err != nil {
		return nil, err
	}

	return &_oss{
		id:       id,
		client:   client,
		secret:   secret,
		opts:     &options,
		endpoint: endpoint,
	}, nil
}

func (p *_oss) ListObject(bucket string) ([]*oss.Object, error) {
	var output []*oss.Object
	bkt, err := p.client.Bucket(bucket)
	if err != nil {
		return nil, err
	}
	lsRes, err := bkt.ListObjects()
	if err != nil {
		return nil, err
	}
	for _, object := range lsRes.Objects {
		output = append(output, &oss.Object{
			Name: object.Key,
			Size: object.Size,
		})
	}
	return output, nil
}

func (p *_oss) Put(native, bkt, objectFile string) error {
	bucket, err := p.client.Bucket(bkt)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectFile, native)
	if err != nil {
		return err
	}
	return nil
}

func (p *_oss) Get(native, bkt, objectFile string) error {
	bucket, err := p.client.Bucket(bkt)
	if err != nil {
		return err
	}

	err = bucket.GetObjectToFile(objectFile, native)
	if err != nil {
		return err
	}
	return nil
}

func (p *_oss) Delete(bkt, objectFile string) error {
	bucket, err := p.client.Bucket(bkt)
	if err != nil {
		return err
	}

	err = bucket.DeleteObject(objectFile)
	if err != nil {
		return err
	}
	return nil
}

func (p *_oss) IsObjectExist(bkt, objectFile string) (bool, error) {
	bucket, err := p.client.Bucket(bkt)
	if err != nil {
		return false, err
	}
	b, err := bucket.IsObjectExist(objectFile)
	if err != nil {
		return false, err
	}
	return b, nil
}
