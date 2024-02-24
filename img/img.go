package img

import (
	"api-gateway/img/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func getImageFromId(id string) ([]byte, error) {

    logFile, err := os.OpenFile(fmt.Sprintf("../../logs/%s.img.log", time.Now().Format("2006-01-02_15-04-05")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(logFile)

	saveURI, ok := os.LookupEnv("SAVE_IMG_URI")
	if !ok {
		log.Fatal("saveIMG URI not found")
	}
	savePort, ok := os.LookupEnv("SAVE_IMG_PORT")
	if !ok {
		log.Fatal("saveIMG port not found")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%s/image/%s", saveURI, savePort, id), nil)
	if err != nil {
		log.Panic("Error during request: ", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Error during request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error during request: ", err)
		return nil, err
	}

	return body, nil
}

func parseImageData(recognition []byte) (models.ImageData, error) {
	var output models.ImageData
	err := json.Unmarshal(recognition, &output)
	return output, err
}

func GetImageFromId(id string) (models.ImageData, error) {
	imageData, err := getImageFromId(id)
	if err != nil {
		log.Panic("Error during extraction: ", err)
		return models.ImageData{}, err
	}
	return parseImageData(imageData)
}

func SendImageToAPI(file multipart.File, userId string, header *multipart.FileHeader) (string,error) {

	logFile, err := os.OpenFile(fmt.Sprintf("../../logs/%s.img.log", time.Now().Format("2006-01-02_15-04-05")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	saveUri, ok := os.LookupEnv("SAVE_IMG_URI")
	if !ok {
		log.Fatal("saving URI not found")
	}

	apiURL := fmt.Sprintf("%s/upload", saveUri)

	var requestBody bytes.Buffer
	multiPartWriter := multipart.NewWriter(&requestBody)

	part, err := multiPartWriter.CreateFormFile("image", header.Filename)
	if err != nil {
		log.Panic("Error creating form file: %v", err)
		return "",err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Panic("Error copying file to form file: %v", err)
		return "",err
	}
	part2,err := multiPartWriter.CreateFormField("userId")
	if err != nil {
		log.Panic("Error creating form field: %v", err)
		return "",err
	}
	part2.Write([]byte(userId))

	err = multiPartWriter.Close()
	if err != nil {
		log.Panic("Error closing multipart writer: %v", err)
		return "",err
	}

	req, err := http.NewRequest("POST", apiURL, &requestBody)
	if err != nil {
		log.Panic("Error creating HTTP request: %v", err)
		return "",err
	}
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Error performing HTTP request: %v", err)
		return "",err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panic("API returned non-OK status: %s", resp.Status)
		return "",fmt.Errorf("API returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error during request: ", err)
		return "",nil
	}
    responseIMG,_:=parseResponse(body)
    log.Printf("Message : %s",responseIMG.Message)
    log.Printf("Id : %s",responseIMG.Id)

	return responseIMG.Id,nil
}

func parseResponse(response []byte) (models.IResponseIMG, error) {
	var output models.IResponseIMG
	err := json.Unmarshal(response, &output)
	return output, err
}