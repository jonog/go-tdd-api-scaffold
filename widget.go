package main

import (
	"errors"
	"strings"
	"time"

	"github.com/go-gorp/gorp"
)

type Widget struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
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

var WidgetColumns string = strings.Join([]string{"id", "name", "created_at", "updated_at"}, ",")

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

func CreateWidget(db *gorp.DbMap, params *WidgetParams) (*Widget, error) {
	resource := BuildWidget(params)
	err := resource.Save(db)
	return resource, err
}

func GetAllWidgets(db *gorp.DbMap) (widgets []Widget, err error) {
	_, err = db.Select(&widgets, "select "+WidgetColumns+" from widgets")
	return widgets, err
}

func FindWidget(db *gorp.DbMap, id int64) (*Widget, error) {
	resource := Widget{}
	err := db.SelectOne(&resource, "select "+WidgetColumns+" from widgets where id=$1", id)
	return &resource, err
}

func (c *Widget) Save(db *gorp.DbMap) (err error) {
	if c.Id == 0 {
		err = db.Insert(c)
	} else {
		_, err = db.Update(c)
	}
	return err
}

func (c *Widget) Validate() error {
	if c.Name == "" {
		return errors.New("Validation error")
	}
	return nil
}

func (c *Widget) Delete(db *gorp.DbMap) (err error) {
	_, err = db.Delete(c)
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
