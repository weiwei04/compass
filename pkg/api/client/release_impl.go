package client

import (
	"bytes"
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"encoding/json"
	//json "github.com/json-iterator/go"
)

// all the ugly things

// NewCompassClient will return a default client impl
func NewReleaseClient(addr string, logger Logger) Release {
	return compassClient{
		host:    addr,
		httpCli: http.DefaultClient,
		logger:  logger,
	}
}

type compassClient struct {
	host    string
	httpCli *http.Client
	logger  Logger
}

func (c compassClient) doRequestWithoutBody(ctx context.Context, method, url string, out interface{}) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errorFromResponse(resp)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c compassClient) doRequest(ctx context.Context, method, url string, in interface{}, out interface{}) error {
	data, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("json.Marshal for method[%s] url[%s] failed, err[%s]", method, url, err))
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errorFromResponse(resp)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c compassClient) CreateRelease(ctx context.Context, req *CreateReleaseRequest) (*CreateReleaseResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s", c.host, req.Namespace, req.Name)
	var resp CreateReleaseResponse
	return &resp, c.doRequest(ctx, "POST", path, req, &resp)
}

func (c compassClient) DeleteRelease(ctx context.Context, req *DeleteReleaseRequest) (*DeleteReleaseResponse, error) {
	// TODO: without body
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s", c.host, req.Namespace, req.Name)
	var resp DeleteReleaseResponse
	return &resp, c.doRequestWithoutBody(ctx, "DELETE", path, &resp)
}

func (c compassClient) UpdateRelease(ctx context.Context, req *UpdateReleaseRequest) (*UpdateReleaseResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/values", c.host, req.Namespace, req.Name)
	var resp UpdateReleaseResponse
	return &resp, c.doRequest(ctx, "PUT", path, req, &resp)
}

func (c compassClient) UpgradeRelease(ctx context.Context, req *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/chart", c.host, req.Namespace, req.Name)
	var resp UpgradeReleaseResponse
	return &resp, c.doRequest(ctx, "PUT", path, req, &resp)
}

func (c compassClient) ListReleases(ctx context.Context, req *ListReleasesRequest) (*ListReleasesResponse, error) {
	// todo: check limit & offset
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases?limit=%d",
		c.host, req.Namespace, req.Limit)
	if req.Offset != "" {
		path = path + "&offset=" + req.Offset
	}
	var resp ListReleasesResponse
	return &resp, c.doRequestWithoutBody(ctx, "GET", path, &resp)
}

func (c compassClient) GetReleaseStatus(ctx context.Context, req *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/status?version=%d", c.host, req.Namespace, req.Name, req.Version)
	var resp GetReleaseStatusResponse
	return &resp, c.doRequest(ctx, "GET", path, req, &resp)
}

func (c compassClient) GetReleaseContent(ctx context.Context, req *GetReleaseContentRequest) (*GetReleaseContentResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/content?version=%d", c.host, req.Namespace, req.Name, req.Version)
	var resp GetReleaseContentResponse
	return &resp, c.doRequest(ctx, "GET", path, req, &resp)
}

func (c compassClient) GetReleaseHistory(ctx context.Context, req *GetReleaseHistoryRequest) (*GetReleaseHistoryResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/history?max=%d", c.host, req.Namespace, req.Name, req.Max)
	var resp GetReleaseHistoryResponse
	return &resp, c.doRequest(ctx, "GET", path, req, &resp)
}

func (c compassClient) RollbackRelease(ctx context.Context, req *RollbackReleaseRequest) (*RollbackReleaseResponse, error) {
	path := fmt.Sprintf("%s/v1/namespaces/%s/releases/%s/version/%d",
		c.host, req.Namespace, req.Name, req.Version)
	var resp RollbackReleaseResponse
	return &resp, c.doRequest(ctx, "POST", path, req, &resp)
}

//func (c compassClient) RunReleaseTest(ctx context.Context, req *TestReleaseRequest) (*TestReleaseResponse, error) {
//	path := fmt.Sprintf("/v1/namespaces/%s/releases/%s/tests")
//	return nil, nil
//}
