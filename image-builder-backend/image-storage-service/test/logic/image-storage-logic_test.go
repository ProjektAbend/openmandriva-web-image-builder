package logic

import (
	"github.com/image-storage-service/cmd/logic"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDoesFileExistShouldReturnTrue(t *testing.T) {
	imageStorageLogic := &logic.ImageStorageLogic{}

	doesExist := imageStorageLogic.DoesFileExist("../resources/test.txt")

	require.True(t, doesExist)
}

func TestDoesFileExistShouldReturnFalse(t *testing.T) {
	imageStorageLogic := &logic.ImageStorageLogic{}

	doesExist := imageStorageLogic.DoesFileExist("../resources/test2.txt")

	require.False(t, doesExist)
}
