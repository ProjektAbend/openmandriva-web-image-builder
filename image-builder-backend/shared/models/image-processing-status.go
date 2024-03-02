package models

type ImageProcessingStatus int

const (
	REQUESTED ImageProcessingStatus = iota
	ACCEPTED
	FETCHED
	BUILD_STARTED
	BUILD_FAILED
	BUILD_FINISHED
	UPLOAD_STARTED
	UPLOAD_FAILED
	AVAILABLE
	EXPIRED
)

type ImageBuildStatus struct {
	ImageId ImageId
	Status  ImageProcessingStatus
}
