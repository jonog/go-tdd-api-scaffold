package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {

	Describe("GET /widgets", func() {

		It("returns the list of widgets", func() {

			r1, err := CreateWidget("Yo 1")
			HandleTestError(err)

			r2, err := CreateWidget("Yo 2")
			HandleTestError(err)

			request, _ := http.NewRequest("GET", "/widgets", nil)
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(200))

			body, _ := ioutil.ReadAll(res.Body)
			widgets := make([]Widget, 0)
			err = json.Unmarshal(body, &widgets)
			HandleTestError(err)

			Ω(widgets).Should(HaveLen(2))

			var receivedWidgetIds []int64
			for _, co := range widgets {
				receivedWidgetIds = append(receivedWidgetIds, co.Id)
			}

			Ω(receivedWidgetIds).Should(ConsistOf(r1.Id, r2.Id))

		})

	})

	Describe("POST /widgets", func() {

		It("creates a widget", func() {

			request, _ := http.NewRequest("POST", "/widgets", bytes.NewReader([]byte(`{"name":"Yo Widget"}`)))
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(201))

			body, _ := ioutil.ReadAll(res.Body)
			widget := Widget{}
			err := json.Unmarshal(body, &widget)
			HandleTestError(err)

			Ω(widget.Name).To(Equal("Yo Widget"))

			var testInt int64
			Ω(widget.Id).Should(BeAssignableToTypeOf(testInt))

		})

		It("returns a 400 if params are invalid", func() {

			request, _ := http.NewRequest("POST", "/widgets", bytes.NewReader([]byte(`{"another_param":"Yo Widget"}`)))
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(400))

		})

	})

	Describe("GET /widgets/:id", func() {

		It("returns a widget", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widget.Id, 10)
			request, _ := http.NewRequest("GET", url, nil)
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(200))

		})

		It("returns a 404 if the widget does not exist", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)
			widgetId := widget.Id
			err = widget.Delete()
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widgetId, 10)
			request, _ := http.NewRequest("GET", url, nil)
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(404))

		})

	})

	Describe("PUT /widgets/:id", func() {

		It("updates a widget", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widget.Id, 10)
			request, _ := http.NewRequest("PUT", url, bytes.NewReader([]byte(`{"name":"Yo 2"}`)))
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(200))

			// check response is the new record
			body, _ := ioutil.ReadAll(res.Body)
			resourceRes := Widget{}
			err = json.Unmarshal(body, &resourceRes)
			HandleTestError(err)
			Ω(resourceRes.Name).To(Equal("Yo 2"))

			// check db has been updated
			updatedResource, err := FindWidget(widget.Id)
			HandleTestError(err)
			Ω(updatedResource.Name).To(Equal("Yo 2"))

		})

		It("returns a 404 if the widget does not exist", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)
			widgetId := widget.Id
			err = widget.Delete()
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widgetId, 10)
			request, _ := http.NewRequest("PUT", url, bytes.NewReader([]byte(`{"name":"Yo 2"}`)))
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(404))

		})

		It("returns a 400 if params are invalid", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widget.Id, 10)

			request, _ := http.NewRequest("PUT", url, bytes.NewReader([]byte(`{"name":""}`)))
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(400))

			updatedResource, err := FindWidget(widget.Id)
			HandleTestError(err)
			Ω(updatedResource.Name).To(Equal("Yo 1"))

		})

	})

	Describe("DELETE /widgets/:id", func() {

		It("deletes a widget", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)

			request, _ := http.NewRequest("DELETE", "/widgets/"+strconv.FormatInt(widget.Id, 10), nil)
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(200))

			_, err = FindWidget(widget.Id)
			Ω(RecordNotFoundError(err)).To(Equal(true))

		})

		It("returns a 404 if the widget does not exist", func() {

			widget, err := CreateWidget("Yo 1")
			HandleTestError(err)
			widgetId := widget.Id
			err = widget.Delete()
			HandleTestError(err)

			url := "/widgets/" + strconv.FormatInt(widgetId, 10)
			request, _ := http.NewRequest("DELETE", url, nil)
			res := httptest.NewRecorder()
			api.Router.ServeHTTP(res, request)
			Ω(res.Code).To(Equal(404))

		})

	})

})
