package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"

	model "github.com/zhyq0826/go-tutorial/devops/models/devops"
	"github.com/zhyq0826/go-tutorial/devops/resources"
)

var domainClient = RequestClient{
	HOST:        "http://127.0.0.1:8000",
	Prefix:      "/domain",
	ContentType: "application/json",
}

func TestDomainList(t *testing.T) {
	url := "/list"
	resp := domainClient.Get(url)
	if resp == nil {
		t.Fail()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []resources.DomainForm
	json.Unmarshal(data, &v)
	for _, d := range v {
		if d.Host == "" || d.Name == "" {
			t.Fail()
		}
	}
}

func TestDomainCreate(t *testing.T) {
	form := resources.DomainForm{
		Domain: model.Domain{
			Name:    "i am name",
			Host:    "i am url",
			Private: 1,
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := domainClient.Post("/create", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestDomainUpdate(t *testing.T) {
	form := resources.DomainForm{
		Domain: model.Domain{
			Name:    "i am name",
			Host:    "i am url",
			Private: 1,
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := domainClient.Put("/1", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestDomainDelete(t *testing.T) {
	resp := domainClient.Delete("/10", io.Reader(nil))
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}
