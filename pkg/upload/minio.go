// @description minio
// @author zkp15
// @date 2023/8/15 11:26
// @version 1.0

package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
	"v-tiktok/model/config"
)

func getMinioClient() *minio.Client {
	m := config.Instance.Uploader.Minio
	var (
		endpoint        = m.Endpoint
		accessKeyID     = m.AccessKeyID
		secretAccessKey = m.SecretAccessKey
		useSSL          = m.UseSSL
	)

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, "mediafiles")
	fmt.Println(exists, err)

	return minioClient
}

func SaveImageOnMinio(videoPath, imageName string, frameNum int) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		logrus.Error("生成缩略图失败：", err)
		return "", err
	}
	content, err := ioutil.ReadAll(buf)
	fileName := time.Now().Format("2006/01/02/") + imageName
	return SaveObject(fileName, content, "image/png")
}

func SaveVideoOnMinio(videoName string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	content, err := ioutil.ReadAll(src)
	fileName := time.Now().Format("2006/01/02/") + videoName
	return SaveObject(fileName, content, "video/mp4")
}

func SaveObject(objectName string, content []byte, contentType string) (string, error) {
	c := config.Instance.Uploader.Minio
	ctx := context.Background()
	bucketName := strings.Trim(c.Path, "/")
	minioClient := getMinioClient()

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if !exists {
		return "", err
	}

	reader := bytes.NewReader(content)
	info, err := minioClient.PutObject(ctx, bucketName, objectName, reader, int64(len(content)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return UrlJoin(c.Host, c.Path, objectName), nil
}
