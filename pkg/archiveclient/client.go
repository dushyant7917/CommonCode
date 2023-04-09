package archiveclient

type ArchiveClient interface {
	GetBucketName() string
	GetLatestFile(string) string
}
