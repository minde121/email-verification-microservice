type spyHandler struct {
	wasCalled bool
	}
	
	func (h *spyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.wasCalled = true
	}
	
	func TestApiKeyMiddleware(t *testing.T) {
	spyHandlerCallTest := &spyHandler{}
	handler := NewAPIKeyMiddleware(spyHandlerCallTest, "myprivatekey")
	server := httptest.NewServer(handler)
	defer server.Close()
	request1, _ := http.NewRequest("POST", server.URL, nil)
	request1.Header.Set("X-API-Key", "myprivatekey")
	resp, _ := http.DefaultClient.Do(request1)
	if resp != nil {
	if resp.StatusCode == http.StatusOK {
	t.Log(request1.Header.Values("X-API-Key"), "Pass")
	}
	if resp.StatusCode == http.StatusForbidden {
	t.Error(request1.Header.Values("X-API-Key"), "Key doesnt match")
	}
	} else {
	t.Error("Request failed")
	}
	request2, _ := http.NewRequest("POST", server.URL, nil)
	request2.Header.Set("X-API-Key", "badkey")
	resp2, _ := http.DefaultClient.Do(request2)
	if resp2 != nil {
	if resp2.StatusCode == http.StatusForbidden {
	t.Error(request2.Header.Values("X-API-Key"), "Key doesnt match")
	}
	if resp2.StatusCode == http.StatusOK {
	t.Error(request2.Header.Values("X-API-Key"), "Pass")
	}
	} else {
	t.Error("Request failed")
	}
	request3, _ := http.NewRequest("POST", server.URL, nil)
	request3.Header.Set("Random", "myprivatekey")
	resp3, _ := http.DefaultClient.Do(request3)
	if resp != nil {
	if resp3.StatusCode == http.StatusOK {
	t.Log(request3.Header, "Required header found")
	}
	if resp3.StatusCode == http.StatusForbidden {
	t.Error(request3.Header, "Invalid Header")
	}
	} else {
	t.Error("Request failed")
	}
	if spyHandlerCallTest.wasCalled == false {
	t.Error("failed to call")
	}
	}