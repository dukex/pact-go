// +build consumer

// Package main contains a runnable Consumer Pact test example.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"

	v3 "github.com/pact-foundation/pact-go/v3"
)

type s = v3.String

func TestConsumerV2(t *testing.T) {
	v3.SetLogLevel("TRACE")

	mockProvider, err := v3.NewHTTPMockProviderV2(v3.MockHTTPProviderConfigV2{
		Consumer: "V2Consumer",
		Provider: "V2Provider",
		Host:     "127.0.0.1",
		Port:     8080,
		TLS:      true,
	})

	// Override default matching behaviour
	// mockProvider.SetMatchingConfig(v3.PactSerialisationOptionsV2{
	// QueryStringStyle: v3.AlwaysArray,
	// QueryStringStyle: v3.Array,
	// QueryStringStyle: v3.Default,
	// })

	// TODO: probably better than deferring to the execute test phase, but not sure
	if err != nil {
		t.Fatal(err)
	}

	// Set up our expected interactions.
	mockProvider.
		AddInteraction().
		Given("User foo exists").
		UponReceiving("A request to do a foo").
		WithRequest(v3.Request{
			Method:  "POST",
			Path:    v3.Regex("/foobar", `\/foo.*`),
			Headers: v3.MapMatcher{"Content-Type": s("application/json"), "Authorization": s("Bearer 1234")},
			Query: v3.QueryMatcher{
				"baz": []interface{}{
					v3.Regex("bar", "[a-z]+"),
					v3.Regex("bat", "[a-z]+"),
					v3.Regex("baz", "[a-z]+"),
				},
			},
			// Body: v3.MapMatcher{
			// 	"name": s("billy"),
			// },
			Body: v3.MatchV2(&User{}),
		}).
		WillRespondWith(v3.Response{
			Status:  200,
			Headers: v3.MapMatcher{"Content-Type": s("application/json")},
			// Body:    v3.Match(&User{}),
			Body: v3.MapMatcher{
				"dateTime": v3.Regex("2020-01-01", "[0-9\\-]+"),
				"name":     s("FirstName"),
				"lastName": s("LastName"),
				"itemsMin": v3.ArrayMinLike("thereshouldbe3ofthese", 3),
				// Add any of these this to demonstrate adding a v3 matcher failing the build (not at the type system level unfortunately)
				// "id": v3.Integer(1),
				// "superstring": v3.Includes("foo"),
				// "accountBalance": v3.Decimal(123.76),
				// "itemsMinMax": v3.ArrayMinMaxLike(27, 3, 5),
				// "equality": v3.Equality("a thing"),
			},
		})

	// Execute pact test
	if err := mockProvider.ExecuteTest(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}

func TestConsumerV3(t *testing.T) {
	v3.SetLogLevel("TRACE")

	mockProvider, err := v3.NewHTTPMockProviderV3(v3.MockHTTPProviderConfigV2{
		Consumer: "V3Consumer",
		Provider: "V3Provider",
		Host:     "127.0.0.1",
		Port:     8080,
		TLS:      true,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Set up our expected interactions.
	mockProvider.
		AddInteraction().
		Given(v3.ProviderStateV3{
			Name: "User foo exists",
			Parameters: map[string]string{
				"id": "foo",
			},
		}).
		UponReceiving("A request to do a foo").
		WithRequest(v3.Request{
			Method:  "POST",
			Path:    v3.Regex("/foobar", `\/foo.*`),
			Headers: v3.MapMatcher{"Content-Type": s("application/json"), "Authorization": s("Bearer 1234")},
			// Body: v3.MapMatcher{
			// 	"name":     s("billy"),
			// 	"dateTime": v3.DateTimeGenerated("2020-02-02", "YYYY-MM-dd"),
			// },

			// Alternative use MatchV3
			Body: v3.MatchV3(&User{}),
			Query: v3.QueryMatcher{
				"baz": []interface{}{
					v3.Regex("bar", "[a-z]+"),
					v3.Regex("bat", "[a-z]+"),
					v3.Regex("baz", "[a-z]+"),
				},
			},
		}).
		WillRespondWith(v3.Response{
			Status:  200,
			Headers: v3.MapMatcher{"Content-Type": s("application/json")},
			// Body:    v3.MatchV3(&User{}),
			Body: v3.MapMatcher{
				"dateTime":       v3.Regex("2020-01-01", "[0-9\\-]+"),
				"name":           s("FirstName"),
				"lastName":       s("LastName"),
				"superstring":    v3.Includes("foo"),
				"id":             v3.Integer(12),
				"accountBalance": v3.Decimal(123.76),
				"itemsMinMax":    v3.ArrayMinMaxLike(27, 3, 5),
				"itemsMin":       v3.ArrayMinLike("thereshouldbe3ofthese", 3),
				"equality":       v3.Equality("a thing"),
			},
		})

	// Execute pact test
	if err := mockProvider.ExecuteTest(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}

func TestMessagePact(t *testing.T) {
	provider, err := v3.NewMessagePactV3(v3.MessageConfig{
		Consumer:             "V3MessageConsumer",
		Provider:             "V3MessageProvider", // must be different to the HTTP one, can't mix both interaction styles
		SpecificationVersion: v3.V3,
	})

	if err != nil {
		t.Fatal(err)
	}

	message := provider.AddMessage()
	message.
		Given(v3.ProviderStateV3{
			Name: "User with id 127 exists",
			Parameters: map[string]interface{}{
				"id": 27,
			},
		}).
		ExpectsToReceive("a user event").
		WithMetadata(v3.MapMatcher{
			"Content-Type": s("application/json; charset=utf-8"),
		}).
		// WithContent(v3.MatchV3(&User{})).
		WithContent(v3.MapMatcher{
			"datetime": v3.Regex("2020-01-01", "[0-9\\-]+"),
			"name":     s("FirstName"),
			"lastName": s("LastName"),
			"id":       v3.Integer(12),
		}).
		AsType(&User{})

	provider.VerifyMessageConsumer(t, message, userHandlerWrapper)
}

func TestPluginPact(t *testing.T) {
	v3.SetLogLevel("TRACE")

	// Start plugin
	go startProvider()

	// Plugin provider rules

	// 0. Client is responsible for starting the plugin.
	// Passing the instructions to start and stop the plugin is just another level of indirection that
	// is unlikely to make it easier to use and especially debug. Whether the service is starting within the test
	// or beforehand in a helper script is

	// 1. Each session must be thread safe, allowing multiple parallel sessions to work in isolation

	// 2. Must communicate over HTTP, implementing the following routes
	// POST /session
	// DELETE /session/:id
	// POST /session/:id/interactions
	// GET /session/:id/mismatches
	// GET /session/:id/log

	// 3. Starting a session with the plugin is the responsibility of the framework, and must conform to a standard API (with optional additional config).

	provider, err := v3.NewPluginProvider(v3.PluginProviderConfig{
		Consumer: "V3MessageConsumer",
		Provider: "V3MessageProvider", // must be different to the HTTP one, can't mix both interaction styles
		Port:     4444,                // Communication port to the provider
	})

	if err != nil {
		t.Fatal(err)
	}

	type tcpInteraction struct {
		Message   string `json:"message"`   // consumer request
		Response  string `json:"response"`  // expected response
		Delimeter string `json:"delimeter"` // how to determine message boundary
	}

	// Plugin providers could create language specific interfaces that except well defined types
	// The raw plugin interface accepts an interface{}
	provider.AddInteraction(tcpInteraction{
		Message:   "hello111",
		Response:  "world!",
		Delimeter: "\r\n",
	})

	// Execute pact test
	if err := provider.ExecuteTest(tcpHelloWorldTest); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}

type User struct {
	ID       int    `json:"id" pact:"example=27"`
	Name     string `json:"name" pact:"example=billy"`
	LastName string `json:"lastName" pact:"example=sampson"`
	Date     string `json:"datetime" pact:"example=2020-01-01'T'08:00:45,format=yyyy-MM-dd'T'HH:mm:ss,generator=datetime"`
	// Date     string `json:"datetime" pact:"example=2020-01-01'T'08:00:45,regex=[0-9-]+,format=yyyy-MM-dd'T'HH:mm:ss,generator=datetime"`
	// Date     string `json:"datetime" pact:"example=20200101,regex=[0-9a-z-A-Z]+"`
}

// Pass in test case
var test = func(config v3.MockServerConfig) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config.TLSConfig,
		},
	}
	req := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Host:     fmt.Sprintf("%s:%d", "localhost", config.Port),
			Scheme:   "https",
			Path:     "/foobar",
			RawQuery: "baz=bat&baz=foo&baz=something", // Default behaviour
			// RawQuery: "baz[]=bat&baz[]=foo&baz[]=something", // TODO: Rust v3 does not support this syntax
		},
		Body:   ioutil.NopCloser(strings.NewReader(`{"id": 27, "name":"billy", "lastName":"sampson", "datetime":"2020-01-01'T'08:00:45"}`)),
		Header: make(http.Header),
	}

	// NOTE: by default, request bodies are expected to be sent with a Content-Type
	// of application/json. If you don't explicitly set the content-type, you
	// will get a mismatch during Verification.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 1234")

	_, err := client.Do(req)

	return err
}

// Message Pact - wrapped handler extracts the message
var userHandlerWrapper = func(m v3.Message) error {
	return userHandler(*m.Content.(*User))
}

// Message Pact - actual handler
var userHandler = func(u User) error {
	if u.ID == 0 {
		return errors.New("invalid object supplied, missing fields (id)")
	}

	// ... actually consume the message

	return nil
}

// Pass in test case
var tcpHelloWorldTest = func(config v3.MockServerConfig) error {
	fmt.Println("executing TCP test")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	defer conn.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, "hello\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	status = strings.TrimSpace(status)
	fmt.Println("response from server:", status)

	// if status != "world!" {
	// 	return fmt.Errorf("expected 'world!', got '%s'.", status)
	// }

	return nil
}
