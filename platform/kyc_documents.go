package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Handles all KYC Document operations
type kycDocumentService service

// KYCDocument data type
type KYCDocument struct {
	UID       string    `json:"uid,omitempty"`
	Type      string    `json:"type,omitempty"`
	Filename  string    `json:"filename,omitempty"`
	Note      string    `json:"note,omitempty"`
	Extension string    `json:"extension,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// KYCDocumentUploadParams are the body params used when uploading a new KYC Document
type KYCDocumentUploadParams struct {
	EvaluationUID string `json:"evaluation_uid"`
	Filename      string `json:"filename"`
	FileContent   string `json:"file_content"`
	Note          string `json:"note"`
	Type          string `json:"type"`
}

// KYCDocumentUploadResponse is the response received when uploading a new KYC Document
type KYCDocumentUploadResponse struct {
	UID       string    `json:"uid,omitempty"`
	Type      string    `json:"type,omitempty"`
	Filename  string    `json:"filename,omitempty"`
	Note      string    `json:"note,omitempty"`
	Extension string    `json:"extension,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// KYCDocumentResponse is an API response containing a list of KYC Documents
type KYCDocumentResponse struct {
	BaseResponse
	Data []*KYCDocument `json:"data"`
}

// List retrieves a list of KYC Documents for a given evaluation
func (k *kycDocumentService) List(evaluationUID string) (*KYCDocumentResponse, error) {
	if evaluationUID == "" {
		return nil, fmt.Errorf("evaluationUID is required")
	}

	v := url.Values{}
	v.Set("evaluation_uid", evaluationUID)

	res, err := k.rizeClient.doRequest(http.MethodGet, "kyc_documents", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &KYCDocumentResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Upload a KYC Document for review
func (k *kycDocumentService) Upload(p *KYCDocumentUploadParams) (*KYCDocumentUploadResponse, error) {
	if p.EvaluationUID == "" ||
		p.Filename == "" ||
		p.FileContent == "" ||
		p.Note == "" ||
		p.Type == "" {
		return nil, fmt.Errorf("All KYCDocumentUploadParams are required")
	}

	bytesMessage, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	res, err := k.rizeClient.doRequest(http.MethodPost, "kyc_documents", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &KYCDocumentUploadResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get is used to retrieve metadata for a KYC Document previously uploaded
func (k *kycDocumentService) Get(uid string) (*KYCDocument, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := k.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("kyc_documents/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &KYCDocument{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// View is used to retrieve a KYC Document (image, PDF, etc) previously uploaded
func (k *kycDocumentService) View(uid string) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	// TODO: Does this require a different Accept header type (image/png)?
	res, err := k.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("kyc_documents/%s/view", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
