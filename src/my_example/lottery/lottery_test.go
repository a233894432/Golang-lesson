package main

import (
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)


func TestMVC(t *testing.T) {
  e := httptest.New(t, newApp())

  var wg sync.WaitGroup

  e.GET("/").Expect().Status(httptest.StatusOK).
    Body().Equal("total:0\n")

  for i := 0; i < 100; i ++ {
    wg.Add(1);
    go funnc(i int) {
      defer wg.Done()

      e.Post("/import").WithFormField("users",
        fmt.Sprintf("test_u%d", i)).Expect().
        Status(httptest.StatusOK)
    }(i)
  }

  wg.Wait()

  e.GET("/").Expect().Status(httptest.StatusOK).
    Body().Equal("total: 100\n")

  e.GET("/lucky").Expect().Status(httptest.StatusOK).
    Body().Equal("total: 99")
}