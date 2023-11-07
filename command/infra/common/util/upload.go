package util

/**
*@Description:
*@Author: BZ
*@date: 2023/11/4 16:34
*@Version: V1.0
 */

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"strconv"
	"strings"
)

const (
	accessKey = "_gFxiIbGDdlYI17zl_nCU2SCeO5neumLlGE0Nvgk"
	secretKey = "Xd-T3IPOYfcqGcBWkpzF9LmbfVwNq0y9whwX32-z"
)

var bucket = "motion-video"
var url = "s3jbnm16b.hd-bkt.clouddn.com"

const (
	Image  = "images"
	Avatar = "avatars"
	Video  = "videos"
)

type UploadRet struct {
	Url      string
	Uuid     int64
	FileName string
	CoverURL string
}

var videoImagExt = []string{"mp4", "flv", "jpg", "png"}

func IsVail(suffix string) bool {
	for _, fileExt := range videoImagExt {
		if suffix == fileExt {
			return true
		}
	}
	return false
}

func UploadFile(f *multipart.FileHeader, class string) (UploadRet, error) {
	filename := f.Filename
	index := strings.LastIndex(filename, ".")

	if !IsVail(filename[index+1:]) {
		return UploadRet{}, errors.New("文件格式不正确")
	}
	id, err := NewID()
	uid := strconv.FormatInt(id, 10)
	key := class + "/" + uid + filename[index:]
	coverKey := bucket + ":" + Image + "/" + uid + "." + "jpg"
	encodeURI := base64.StdEncoding.EncodeToString([]byte(coverKey))

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//视频生成封面
	if class == Video {
		putPolicy.PersistentOps = "vframe/jpg/offset/1|saveas/" + encodeURI
	}

	file, _ := f.Open()
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err = formUploader.Put(context.Background(), &ret, upToken, key, file, f.Size, &putExtra)
	url2 := url + "/" + ret.Key

	return UploadRet{
		url2,
		id,
		filename,
		url + "/" + Image + "/" + uid + "." + "jpg",
	}, err
}
