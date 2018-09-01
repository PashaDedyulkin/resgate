package test

import (
	"encoding/json"
	"testing"
)

// Test that a client can unsubscribe to a model
func TestUnsubscribeModel(t *testing.T) {
	runTest(t, func(s *Session) {
		model := resource["test.model"]

		c := s.Connect()
		creq := c.Request("subscribe.test.model", nil)

		// Handle model get and access request
		mreqs := s.GetParallelRequests(t, 2)
		req := mreqs.GetRequest(t, "access.test.model")
		req.RespondSuccess(json.RawMessage(`{"get":true}`))
		req = mreqs.GetRequest(t, "get.test.model")
		req.RespondSuccess(json.RawMessage(`{"model":` + model + `}`))

		// Get client response
		creq.GetResponse(t)

		// Call unsubscribe
		c.Request("unsubscribe.test.model", nil)
	})
}

// Test that a client can unsubscribe to a collection
func TestUnsubscribeCollection(t *testing.T) {
	runTest(t, func(s *Session) {
		collection := resource["test.collection"]

		c := s.Connect()
		creq := c.Request("subscribe.test.collection", nil)

		// Handle collection get and access request
		mreqs := s.GetParallelRequests(t, 2)
		req := mreqs.GetRequest(t, "access.test.collection")
		req.RespondSuccess(json.RawMessage(`{"get":true}`))
		req = mreqs.GetRequest(t, "get.test.collection")
		req.RespondSuccess(json.RawMessage(`{"collection":` + collection + `}`))

		// Get client response
		creq.GetResponse(t)

		// Call unsubscribe
		c.Request("unsubscribe.test.collection", nil)
	})
}
