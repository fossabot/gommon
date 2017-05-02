package requests

import (
	"testing"
	"net/http"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	asst "github.com/stretchr/testify/assert"
)

func TestRequestsE2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skip requests e2e test")
	}

	// create an echo server
	version := "0.0.1"
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Log("can't read request body")
		}
		fmt.Fprintf(w, "%s", data)
		// The Server will close the request body. The ServeHTTP
		// Handler does not need to.
		//r.Body.Close()
	})
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"version\": \"%s\"}", version)
	})

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	versionURL := testServer.URL + "/version"
	echoURL := testServer.URL + "/echo"

	t.Run("Get", func(t *testing.T) {
		assert := asst.New(t)
		res, err := Get(versionURL)
		assert.Nil(err)
		assert.Equal(http.StatusOK, res.Res.StatusCode)
	})

	t.Run("GetJSONStringMap", func(t *testing.T) {
		assert := asst.New(t)
		data, err := GetJSONStringMap(versionURL)
		assert.Nil(err)
		assert.Equal(version, data["version"])
	})

	t.Run("PostJSONString", func(t *testing.T) {
		assert := asst.New(t)
		payload := fmt.Sprintf("{\"version\": \"%s\"}", version)
		res, err := PostJSONString(echoURL, payload)
		assert.Nil(err)
		assert.Equal(payload, string(res.Text))
	})
}
