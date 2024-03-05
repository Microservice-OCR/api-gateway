package ocr

import (
	"api-gateway/ocr/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "time"
)

func getExtract(imageName string) ([]byte, error) {

    // logFile, err := os.OpenFile(fmt.Sprintf("../../logs/%s.ocr.log", time.Now().Format("2006-01-02_15-04-05")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.SetOutput(logFile)

	ocrURI, ok := os.LookupEnv("OCR_ENGINE_URI")
	if !ok {
		log.Fatal("OCR engine URI not found")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/recognize/%s", ocrURI, imageName), nil)
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

func getExtractFromId(id string) ([]byte, error) {

    // logFile, err := os.OpenFile(fmt.Sprintf("../../logs/%s.ocr.log", time.Now().Format("2006-01-02_15-04-05")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.SetOutput(logFile)

	ocrURI, ok := os.LookupEnv("OCR_ENGINE_URI")
	if !ok {
		log.Fatal("OCR engine URI not found")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/recognizeFromId/%s", ocrURI, id), nil)
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

func postExtractFromId(id string, bodyData models.IInput) ([]byte, error) {

    // logFile, err := os.OpenFile(fmt.Sprintf("../../logs/%s.ocr.log", time.Now().Format("2006-01-02_15-04-05")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.SetOutput(logFile)

	ocrURI, ok := os.LookupEnv("OCR_ENGINE_URI")
	if !ok {
		log.Fatal("OCR engine URI not found")
	}
	// Convertir les données du corps en JSON
	jsonData, err := json.Marshal(bodyData)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/recognizeFromId/%s", ocrURI, id), bytes.NewBuffer((jsonData)))
	if err != nil {
		log.Panic("Error during request: ", err)
		return nil, err
	}
	// Définir l'en-tête Content-Type sur application/json
	req.Header.Set("Content-Type", "application/json")

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

func parseRecognition(recognition []byte) (models.Recognition, error) {
	var output models.Recognition
	err := json.Unmarshal(recognition, &output)
	return output, err
}

func parseComplexRecognition(recognition []byte) ([]models.IComplexRecognition, error) {
	var output []models.IComplexRecognition
	err := json.Unmarshal(recognition, &output)
	return output, err
}

func GetOCR(imageName string) (models.Recognition, error) {
	recognition, err := getExtract(imageName)
	if err != nil {
		log.Panic("Error during extraction: ", err)
		return models.Recognition{}, err
	}
	return parseRecognition(recognition)
}

func GetOCRFromId(id string) (models.Recognition, error) {
	recognition, err := getExtractFromId(id)
	if err != nil {
		log.Panic("Error during extraction: ", err)
		return models.Recognition{}, err
	}
	return parseRecognition(recognition)
}

func PostOCRFromId(id string, bodyData models.IInput) ([]models.IComplexRecognition, error) {
	recognition, err := postExtractFromId(id, bodyData)
	if err != nil {
		log.Panic("Error during extraction: ", err)
		return []models.IComplexRecognition{}, err
	}
	return parseComplexRecognition(recognition)
}