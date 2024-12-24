package sacloud

import (
	"os"
	"testing"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/testutil"
)

var testZone string
var testCaller *iaas.Client

func TestMain(m *testing.M) {
	// this is for to use fake driver on iaas-api-go
	os.Setenv("TESTACC", "")

	testZone = testutil.TestZone()
	testCaller = testutil.SingletonAPICaller().(*iaas.Client)

	ret := m.Run()
	os.Exit(ret)
}
