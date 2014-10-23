package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keyserver Suite")
}

func HandleTestError(err error) {
	if err != nil {
		panic(err)
	}
}

var _ = BeforeSuite(func() {

	api.InitDB()
	api.InitRoutes()

})

var _ = AfterSuite(func() {

	api.DB.Db.Close()

})

var _ = BeforeEach(func() {

	err := api.DB.TruncateTables()
	if err != nil {
		panic(err)
	}

})
