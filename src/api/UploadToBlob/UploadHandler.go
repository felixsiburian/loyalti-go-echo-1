package UploadToBlob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

func GetAccountInfo()(string, string, string, string){
	azrKey := fmt.Sprintf("%s",viper.Get("AZRKEY"))
	fmt.Println("azrKey : ", azrKey)
	azrBlobAccountName := fmt.Sprintf("%s",viper.Get("AZRACCOUNTNAME"))
	fmt.Println("azrAccount : ", azrBlobAccountName)
	azrBlobContainer := fmt.Sprintf("%s",viper.Get("AZRBLOBCONTAINER"))
	fmt.Println("azrContainer : ", azrBlobContainer)
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	fmt.Println(azrPrimaryBlobServiceEndpoint)
	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}

func GetBlobName() string {
	t := time.Now()
	uuid, _ := uuid.NewV4()

	return fmt.Sprintf("image_%s-%v.jpg", t.Format("20060102"), uuid)
}

func UplodBytesToBlob(b []byte)(string, error) {
	fmt.Println("masuk kesini")
	azrKey, accountName, endPoint, container := GetAccountInfo()
	u, err := url.Parse(fmt.Sprint(endPoint, container, "/", GetBlobName()))
	if err != nil {
		fmt.Println("Error parse : ", err.Error())
	}
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)
	if errC != nil {
		return "", errC
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
	fmt.Println(blockBlobUrl)
	ctx := context.Background()
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "image/jpg",
		},
	}
	fmt.Println(o)

	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	if errU != nil {
		fmt.Println("Erorr Upload ", errU.Error())
	}
	return blockBlobUrl.String(), errU
}
