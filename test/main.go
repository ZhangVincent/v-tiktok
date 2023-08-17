// @description 测试代码
// @author zkp15
// @date 2023/8/9 11:42
// @version 1.0

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/user"
	"v-tiktok/pkg/strs"
)

func init() {
	defer fmt.Println("init defer success")
}

func main() {
	//s := "user_:%d"
	//sprintf := fmt.Sprintf(s, 111)
	//fmt.Println(sprintf)

	//testSimpleString()
	testZSet()
}

func testZSet() {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "centos:6379",
	//	Password: "287177741",
	//	DB:       0,
	//})

	opt, err := redis.ParseURL("redis://:28717774@centos:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	ctx := context.Background()

	marshal, _ := json.Marshal(user.User{ID: 1, Name: "useroftiktok1"})
	userString1 := string(marshal)
	marshal, _ = json.Marshal(user.User{ID: 2, Name: "useroftiktok2"})
	userString2 := string(marshal)
	marshal, _ = json.Marshal(user.User{ID: 3, Name: "useroftiktok3"})
	userString3 := string(marshal)
	zAddResult := client.ZAdd(ctx, "users", redis.Z{Score: 1001, Member: userString1}, redis.Z{Score: 2011, Member: userString2}, redis.Z{Score: 2011, Member: model.User{
		Name:     "vincent",
		Password: "pawwss",
	}})
	if zAddResult.Err() != nil {
		fmt.Println(zAddResult.Err())
	}

	rangeWithScores := client.ZRangeWithScores(ctx, "users", 0, 1)
	result, err := rangeWithScores.Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	client.ZAdd(ctx, "users", redis.Z{Score: 2101, Member: userString3})

	rank := client.ZRank(ctx, "users", "2000")
	i, _ := rank.Result()

	zRange := client.ZRange(ctx, "users", i, i+30)
	i2, err := zRange.Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i2)
}

func testSimpleString() {
	client := redis.NewClient(&redis.Options{
		Addr:     "centos:6379",
		Password: "28717774",
		DB:       0,
	})

	ctx := context.Background()

	userInfo := user.User{
		ID:   1,
		Name: "useroftiktok",
	}
	marshal, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(ctx, "vtiktok:user_2").Result()
	if err != nil {
		fmt.Println(err)
	}
	userString := string(marshal)
	err = client.Set(ctx, "vtiktok:user_2", userString, 30*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err = client.Get(ctx, "vtiktok:user_3").Result()
	if err != nil {
		fmt.Println(err)
	}
	var userResult user.User
	err = json.Unmarshal([]byte(val), &userResult)
	fmt.Println(userResult)
}

func testSaveVideo() {
	file, err := os.Open("static/video/defbf88f8874449b906682237447eff5.mp4")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(SaveObject("testMinio.mp4", content, "video/mp4"))
}

func SaveObject(objectName string, content []byte, contentType string) error {
	ctx := context.Background()
	bucketName := strings.Trim("/mediafiles", "/")

	var (
		endpoint        = "centos:9000"
		accessKeyID     = "vincent"
		secretAccessKey = "28717774"
		useSSL          = false
	)

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if !exists {
		return err
	}

	reader := bytes.NewReader(content)

	info, err := minioClient.PutObject(ctx, bucketName, objectName, reader, int64(len(content)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return nil
}

func getFileSize(file io.Reader) int64 {
	var size int64 = 0
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		size += int64(n)
		if err == io.EOF {
			break
		}
	}
	return size
}

func testTime() {
	//s := "abscdfa"
	//fmt.Println(s)
	//testString(&s)
	//fmt.Println(s)

	var t int64 = 1691026536331
	fmt.Println(time.Now().UnixMilli() / t)
	milli := time.UnixMilli(t)
	fmt.Println(milli.Format("2006-01-02 15:04:05"))
}

func testString(s *string) {
	*s = strings.ReplaceAll(*s, "a", "A")
	fmt.Println(*s)
}

func testGetImageFromVideo() {
	_, err := GetSnapshot("https://www.w3schools.com/html/movie.mp4", "static/pic/test", 1)
	if err != nil {
		return
	}
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	names := strings.Split(snapshotPath, `\`)
	snapshotName = names[len(names)-1] + ".png"
	return
}

func testUUID() {
	s := strs.UUID()
	fmt.Println(s)

	id := uuid.NewV4()
	fmt.Println(id.String())
}
