package main

import (
	"errors"
	"time"

	"github.com/go-gorp/gorp"
)

type Widget struct {
	Id        int64
	Name      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type WidgetPublic struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type WidgetParams struct {
	Name string `json:"name"`
}

func (c *Widget) Export() WidgetPublic {
	return WidgetPublic{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ExportWidgets(widgets []Widget) []WidgetPublic {
	widgetsPublic := make([]WidgetPublic, len(widgets))
	for idx, resource := range widgets {
		widgetsPublic[idx] = resource.Export()
	}
	return widgetsPublic
}

func BuildWidget(params *WidgetParams) *Widget {
	resource := Widget{Name: params.Name}
	return &resource
}

func CreateWidget(params *WidgetParams) (*Widget, error) {
	resource := BuildWidget(params)
	err := resource.Save()
	return resource, err
}

func GetAllWidgets() (widgets []Widget, err error) {
	_, err = api.DB.Select(&widgets, "SELECT * FROM widgets")
	return widgets, err
}

func FindWidget(id int64) (*Widget, error) {
	resource := Widget{}
	err := api.DB.SelectOne(&resource, "select * from widgets where id=$1", id)
	return &resource, err
}

func (c *Widget) Save() (err error) {
	if c.Id == 0 {
		err = api.DB.Insert(c)
	} else {
		_, err = api.DB.Update(c)
	}
	return err
}

func (c *Widget) Validate() error {
	if c.Name == "" {
		return errors.New("Validation error")
	}
	return nil
}

func (c *Widget) Delete() (err error) {
	_, err = api.DB.Delete(c)
	return err
}

func (r *Widget) PreInsert(s gorp.SqlExecutor) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = r.CreatedAt
	return nil
}

func (r *Widget) PreUpdate(s gorp.SqlExecutor) error {
	r.UpdatedAt = time.Now()
	return nil
}
