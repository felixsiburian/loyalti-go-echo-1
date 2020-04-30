package UploadToBlob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/gofrs/uuid"
	"net/url"
	"time"
)

func GetAccountInfo()(string, string, string, string){
	azrKey := "194MUKH8BxaH4xasKSVedJgS6mLR6FLVbxVYbZLii13ZI5WnN308xIZBRcHsarw8nn+D5+O3p15r7BdtWHHQTw=="
	fmt.Println("azrKey : ", azrKey)
	azrBlobAccountName := "loyaltiexpress"
	fmt.Println("azrAccount : ", azrBlobAccountName)
	azrBlobContainer := "loyalti-images"
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
	fmt.Println("U : ", u  )
	if err != nil {
		fmt.Println("Error parse : ", err.Error())
	}
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)
	if errC != nil {
		return "", errC
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
	fmt.Println("blobURL : ",blockBlobUrl)
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

	//var bloburl = azblob
	return blockBlobUrl.String(), errU
}
