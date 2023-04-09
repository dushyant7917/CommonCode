// +build aws ali

package archiveclient

type ArchiveClient interface {
	GetBucketName() string
	GetLatestFile(string) string
}
