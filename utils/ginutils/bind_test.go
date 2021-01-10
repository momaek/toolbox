package ginutils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

type testReq struct {
	ID     string    `in:"path:id"`
	Status *string   `in:"query:status,required,notnull"`
	Time   time.Time `in:"query:time"`

	Name string `json:"name"`
	Desc string `json:"desc"`
}

func TestBind(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	engine.POST("/v1/instances/:id", func(c *gin.Context) {
		var args testReq
		err := Bind(c, &args)
		if err != nil {
			log.Println(err)
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.AbortWithStatus(http.StatusOK)
	})

	svr := httptest.NewServer(engine)

	body := map[string]string{
		"name": "helloworld",
		"desc": "desc",
	}

	b, _ := json.Marshal(&body)

	req, _ := http.NewRequest(http.MethodPost, svr.URL+"/v1/instances/299?status=success&time=2021-01-02", bytes.NewBuffer(b))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("unexpected error ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("unexpected code")
		return
	}

	req, _ = http.NewRequest(http.MethodPost, svr.URL+"/v1/instances/299?status=&time=2021-01-02", bytes.NewBuffer(b))

	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("unexpected error ", err)
		return
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusBadRequest {
		t.Fatal("unexpected code")
		return
	}

	req, _ = http.NewRequest(http.MethodPost, svr.URL+"/v1/instances/299?time=2021-01-02", bytes.NewBuffer(b))

	resp3, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("unexpected error ", err)
		return
	}
	defer resp3.Body.Close()

	if resp3.StatusCode != http.StatusBadRequest {
		t.Fatal("unexpected code")
		return
	}
}
