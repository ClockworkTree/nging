// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_CollectorRule []*CollectorRule

func (s Slice_CollectorRule) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_CollectorRule) RangeRaw(fn func(m *CollectorRule) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// CollectorRule 页面中的元素采集规则
type CollectorRule struct {
	base    factory.Base
	objects []*CollectorRule

	Id      uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	PageId  uint   `db:"page_id" bson:"page_id" comment:"页面ID" json:"page_id" xml:"page_id"`
	Name    string `db:"name" bson:"name" comment:"保存匹配结果的变量名" json:"name" xml:"name"`
	Rule    string `db:"rule" bson:"rule" comment:"规则" json:"rule" xml:"rule"`
	Type    string `db:"type" bson:"type" comment:"数据类型" json:"type" xml:"type"`
	Filter  string `db:"filter" bson:"filter" comment:"过滤器" json:"filter" xml:"filter"`
	Created uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Sort    int    `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
}

// - base function

func (a *CollectorRule) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *CollectorRule) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *CollectorRule) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *CollectorRule) Context() echo.Context {
	return a.base.Context()
}

func (a *CollectorRule) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *CollectorRule) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *CollectorRule) Namer() func(string) string {
	return a.base.Namer()
}

func (a *CollectorRule) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *CollectorRule) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *CollectorRule) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *CollectorRule) Objects() []*CollectorRule {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *CollectorRule) NewObjects() factory.Ranger {
	return &Slice_CollectorRule{}
}

func (a *CollectorRule) InitObjects() *[]*CollectorRule {
	a.objects = []*CollectorRule{}
	return &a.objects
}

func (a *CollectorRule) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *CollectorRule) Short_() string {
	return "collector_rule"
}

func (a *CollectorRule) Struct_() string {
	return "CollectorRule"
}

func (a *CollectorRule) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *CollectorRule) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *CollectorRule) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *CollectorRule) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *CollectorRule) GroupBy(keyField string, inputRows ...[]*CollectorRule) map[string][]*CollectorRule {
	var rows []*CollectorRule
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*CollectorRule{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CollectorRule{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *CollectorRule) KeyBy(keyField string, inputRows ...[]*CollectorRule) map[string]*CollectorRule {
	var rows []*CollectorRule
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*CollectorRule{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *CollectorRule) AsKV(keyField string, valueField string, inputRows ...[]*CollectorRule) map[string]interface{} {
	var rows []*CollectorRule
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (a *CollectorRule) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *CollectorRule) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Type) == 0 {
		a.Type = "string"
	}
	err = DBI.Fire("creating", a, nil)
	if err != nil {
		return
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *CollectorRule) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.Type) == 0 {
		a.Type = "string"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *CollectorRule) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *CollectorRule) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "string"
		}
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *CollectorRule) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.Type) == 0 {
			a.Type = "string"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Type) == 0 {
			a.Type = "string"
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *CollectorRule) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *CollectorRule) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *CollectorRule) Reset() *CollectorRule {
	a.Id = 0
	a.PageId = 0
	a.Name = ``
	a.Rule = ``
	a.Type = ``
	a.Filter = ``
	a.Created = 0
	a.Sort = 0
	return a
}

func (a *CollectorRule) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["PageId"] = a.PageId
	r["Name"] = a.Name
	r["Rule"] = a.Rule
	r["Type"] = a.Type
	r["Filter"] = a.Filter
	r["Created"] = a.Created
	r["Sort"] = a.Sort
	return r
}

func (a *CollectorRule) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "page_id":
			a.PageId = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "rule":
			a.Rule = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "filter":
			a.Filter = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "sort":
			a.Sort = param.AsInt(value)
		}
	}
}

func (a *CollectorRule) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "PageId":
			a.PageId = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Rule":
			a.Rule = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Filter":
			a.Filter = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Sort":
			a.Sort = param.AsInt(vv)
		}
	}
}

func (a *CollectorRule) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["page_id"] = a.PageId
	r["name"] = a.Name
	r["rule"] = a.Rule
	r["type"] = a.Type
	r["filter"] = a.Filter
	r["created"] = a.Created
	r["sort"] = a.Sort
	return r
}

func (a *CollectorRule) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *CollectorRule) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
