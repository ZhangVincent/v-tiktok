// @description 保存文件到本地
// @author zkp15
// @date 2023/8/15 15:22
// @version 1.0

package upload

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func SaveVideoOnLocal(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func SaveImageOnLocal(videoPath, snapshotPath string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		logrus.Error("生成缩略图失败：", err)
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	err = imaging.Save(img, snapshotPath)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	return nil
}
