package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestHandlerBasic(t *testing.T) {

	input, _ := json.Marshal("{City:'Amsterdam'}")

	ans, err := handleRequest(context.TODO(), input)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Log(ans.Body)

	var output City
	_ = json.Unmarshal([]byte(ans.Body), &output)

	if !strings.EqualFold(output.Name, "Amsterdam") {
		t.Errorf("We didn't find Amsterdam city")
	}
}
