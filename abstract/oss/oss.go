package oss

type Object struct {
	Name string
	Size int64
}

type OSS interface {
	ListObject(string) ([]*Object, error)
	Get(native string, bucket string, objectFile string) error
	Put(native string, bucket string, objectFile string) error // native, bucket, destination
	Delete(bucket, objectFile string) error
}
