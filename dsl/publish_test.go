package dsl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

func TestPublish_validate(t *testing.T) {
	dir, _ := os.Getwd()
	testFile := fmt.Sprintf(filepath.Join(dir, "publish_test.go"))

	p := &Publisher{
		request: &types.PublishRequest{},
	}

	err := p.validate()
	if err.Error() != "PactURLs is mandatory" {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactURLs: []string{testFile},
		},
	}

	err = p.validate()
	if err.Error() != "PactBroker is mandatory" {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactBroker: "http://foo.com",
			PactURLs:   []string{testFile},
		},
	}

	err = p.validate()
	if err.Error() != "ConsumerVersion is mandatory" {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactBroker: "http://foo.com",
			PactURLs: []string{
				testFile,
			},
			ConsumerVersion:    "1.0.0",
			PactBrokerUsername: "userwithoutpass",
		},
	}

	err = p.validate()
	if err.Error() != "Must provide both or none of PactBrokerUsername and PactBrokerPassword" {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactBroker: "http://foo.com",
			PactURLs: []string{
				testFile,
			},
			ConsumerVersion:    "1.0.0",
			PactBrokerPassword: "passwithoutuser",
		},
	}

	err = p.validate()
	if err.Error() != "Must provide both or none of PactBrokerUsername and PactBrokerPassword" {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactURLs: []string{
				"aoeuaoeu",
			},
		},
	}

	err = p.validate()
	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Fatalf("Expected a different error but got '%s'", err.Error())
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactBroker: "http://foo.com",
			PactURLs: []string{
				testFile,
			},
			ConsumerVersion: "1.0.0",
		},
	}

	err = p.validate()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	p = &Publisher{
		request: &types.PublishRequest{
			PactBroker: "http://foo.com",
			PactURLs: []string{
				testFile,
			},
			ConsumerVersion: "1.0.0",
		},
	}

	err = p.validate()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestPublish_readLocalPactFile(t *testing.T) {
	file := createSimplePact(true)
	defer os.Remove(file.Name())
	p := &Publisher{request: &types.PublishRequest{}}

	f, _, err := p.readLocalPactFile(file.Name())

	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	if f.Consumer.Name != "Some Consumer" {
		t.Fatalf("Expected Consumer name to be 'Some Consumer'")
	}

	if f.Provider.Name != "Some Provider" {
		t.Fatalf("Expected Provider name to be 'Some Provider'")
	}
}

func TestPublish_readLocalPactFileFail(t *testing.T) {
	p := &Publisher{request: &types.PublishRequest{}}
	_, _, err := p.readLocalPactFile("thisfileprobablydoesntexist")

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	brokenFile := createSimplePact(false)
	defer os.Remove(brokenFile.Name())

	_, _, err = p.readLocalPactFile(brokenFile.Name())

	if err == nil {
		t.Fatalf("Expected error but got none")
	}
}

func TestPublish_readRemotePactFile(t *testing.T) {
	p := &Publisher{request: &types.PublishRequest{}}
	url := createMockRemoteServer(true)

	f, _, err := p.readRemotePactFile(url)

	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	if f.Consumer.Name != "Some Consumer" {
		t.Fatalf("Expected Consumer name to be 'Some Consumer'")
	}

	if f.Provider.Name != "Some Provider" {
		t.Fatalf("Expected Provider name to be 'Some Provider'")
	}
}

func TestPublish_readRemotePactFileFail(t *testing.T) {
	p := &Publisher{request: &types.PublishRequest{}}
	url := createMockRemoteServer(false)

	_, _, err := p.readRemotePactFile(url)

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	_, _, err = p.readRemotePactFile(fmt.Sprintf("%s/iknowthisfiledoesntexist", url))
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
}

func createMockRemoteServer(valid bool) string {
	file := createSimplePact(valid)
	dir := filepath.Dir(file.Name())
	path := filepath.Base(file.Name())
	port, _ := utils.GetFreePort()
	go http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))

	return fmt.Sprintf("http://localhost:%d/%s", port, path)
}

func createSimplePact(valid bool) *os.File {
	var data []byte
	if valid {
		data = []byte(`
    {
      "consumer": {
        "name": "Some Consumer"
      },
      "provider": {
        "name": "Some Provider"
      }
    }
  `)
	} else {
		data = []byte(`
    {
      "consumer": {
        "name": "Some Consumer"
      }
    }
  `)
	}

	tmpfile, err := ioutil.TempFile("/tmp", "pactgo")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func TestPublish_readPactFile(t *testing.T) {
	p := &Publisher{request: &types.PublishRequest{}}
	url := createMockRemoteServer(true)

	f, _, err := p.readPactFile(url)

	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	if f.Consumer.Name != "Some Consumer" {
		t.Fatalf("Expected Consumer name to be 'Some Consumer'")
	}

	if f.Provider.Name != "Some Provider" {
		t.Fatalf("Expected Provider name to be 'Some Provider'")
	}

	localFile := createSimplePact(true)
	f, _, err = p.readPactFile(localFile.Name())

	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	if f.Consumer.Name != "Some Consumer" {
		t.Fatalf("Expected Consumer name to be 'Some Consumer'")
	}

	if f.Provider.Name != "Some Provider" {
		t.Fatalf("Expected Provider name to be 'Some Provider'")
	}
}

func TestPublish_readPactFileFail(t *testing.T) {
	p := &Publisher{request: &types.PublishRequest{}}
	url := createMockRemoteServer(false)

	_, _, err := p.readPactFile(url)

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	_, _, err = p.readPactFile(fmt.Sprintf("%s/iknowthisfiledoesntexist", url))
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
}

func TestPublish_Publish(t *testing.T) {
	p := &Publisher{}

	f := createSimplePact(true)
	broker := setupMockServer(true, t)
	defer broker.Close()

	err := p.Publish(&types.PublishRequest{
		PactURLs:        []string{f.Name()},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestPublish_PublishFail(t *testing.T) {
	p := &Publisher{}

	broker := setupMockServer(true, t)
	defer broker.Close()

	err := p.Publish(&types.PublishRequest{
		PactURLs:        []string{"aoeuaoeuaoeu"},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
	})

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	err = p.Publish(&types.PublishRequest{
		PactURLs:        []string{"http://localhost:1234/foo/bar"},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
	})

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	// real file but broken broker
	f := createSimplePact(true)
	brokenBroker := setupMockServer(false, t)
	defer broker.Close()
	err = p.Publish(&types.PublishRequest{
		PactURLs:        []string{f.Name()},
		PactBroker:      brokenBroker.URL,
		ConsumerVersion: "1.0.0",
	})

	if err == nil {
		t.Fatalf("Expected error but got none")
	}

	if strings.TrimSpace(err.Error()) != strings.TrimSpace("something went wrong") {
		t.Fatalf("Expected error to be 'something went wrong' but got: %s", err.Error())
	}
}

func TestPublish_PublishWithTags(t *testing.T) {
	p := &Publisher{}

	f := createSimplePact(true)
	broker := setupMockServer(true, t)
	defer broker.Close()

	err := p.Publish(&types.PublishRequest{
		PactURLs:        []string{f.Name()},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
		Tags:            []string{"latest"},
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestPublish_tagRequest(t *testing.T) {
	p := &Publisher{}
	f := createSimplePact(true)

	broker := setupMockServer(true, t)
	defer broker.Close()

	err := p.tagRequest("Some Consumer", &types.PublishRequest{
		PactURLs:        []string{f.Name()},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
		Tags:            []string{"latest"},
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestPublish_tagRequestFail(t *testing.T) {
	p := &Publisher{}
	f := createSimplePact(true)

	broker := setupMockServer(false, t)
	defer broker.Close()

	err := p.tagRequest("Some Consumer", &types.PublishRequest{
		PactURLs:        []string{f.Name()},
		PactBroker:      broker.URL,
		ConsumerVersion: "1.0.0",
		Tags:            []string{"latest"},
	})

	if err == nil {
		t.Fatalf("Expected error but got none")
	}
}

func TestPublish_EndToEnd(t *testing.T) {
	t.Skip()
	p := &Publisher{}
	err := p.Publish(&types.PublishRequest{
		PactURLs:        []string{"../pacts/billy-bobby.json"},
		PactBroker:      "http://localhost:8080",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"latest", "foobar", "sit4"},
	})

	fmt.Println(err)

}
