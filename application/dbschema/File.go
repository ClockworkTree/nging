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

type Slice_File []*File

func (s Slice_File) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_File) RangeRaw(fn func(m *File) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// File 文件表
type File struct {
	base    factory.Base
	objects []*File

	Id         uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"文件ID" json:"id" xml:"id"`
	OwnerType  string `db:"owner_type" bson:"owner_type" comment:"用户类型" json:"owner_type" xml:"owner_type"`
	OwnerId    uint64 `db:"owner_id" bson:"owner_id" comment:"用户ID" json:"owner_id" xml:"owner_id"`
	Name       string `db:"name" bson:"name" comment:"原始文件名" json:"name" xml:"name"`
	SaveName   string `db:"save_name" bson:"save_name" comment:"保存名称" json:"save_name" xml:"save_name"`
	SavePath   string `db:"save_path" bson:"save_path" comment:"文件保存路径" json:"save_path" xml:"save_path"`
	ViewUrl    string `db:"view_url" bson:"view_url" comment:"查看链接" json:"view_url" xml:"view_url"`
	Ext        string `db:"ext" bson:"ext" comment:"文件后缀" json:"ext" xml:"ext"`
	Mime       string `db:"mime" bson:"mime" comment:"文件mime类型" json:"mime" xml:"mime"`
	Type       string `db:"type" bson:"type" comment:"文件类型" json:"type" xml:"type"`
	Size       uint64 `db:"size" bson:"size" comment:"文件大小" json:"size" xml:"size"`
	Width      uint   `db:"width" bson:"width" comment:"宽度(像素)" json:"width" xml:"width"`
	Height     uint   `db:"height" bson:"height" comment:"高度(像素)" json:"height" xml:"height"`
	Dpi        uint   `db:"dpi" bson:"dpi" comment:"分辨率" json:"dpi" xml:"dpi"`
	Md5        string `db:"md5" bson:"md5" comment:"文件md5" json:"md5" xml:"md5"`
	StorerName string `db:"storer_name" bson:"storer_name" comment:"文件保存位置" json:"storer_name" xml:"storer_name"`
	StorerId   string `db:"storer_id" bson:"storer_id" comment:"位置ID" json:"storer_id" xml:"storer_id"`
	Created    uint   `db:"created" bson:"created" comment:"上传时间" json:"created" xml:"created"`
	Updated    uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Project    string `db:"project" bson:"project" comment:"项目名称" json:"project" xml:"project"`
	TableId    string `db:"table_id" bson:"table_id" comment:"关联表数据id" json:"table_id" xml:"table_id"`
	TableName  string `db:"table_name" bson:"table_name" comment:"关联表名称" json:"table_name" xml:"table_name"`
	FieldName  string `db:"field_name" bson:"field_name" comment:"关联表字段名" json:"field_name" xml:"field_name"`
	Sort       int64  `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Status     int    `db:"status" bson:"status" comment:"状态(1-已审核/0-未审核)" json:"status" xml:"status"`
	CategoryId uint   `db:"category_id" bson:"category_id" comment:"分类ID" json:"category_id" xml:"category_id"`
	UsedTimes  uint   `db:"used_times" bson:"used_times" comment:"被使用的次数" json:"used_times" xml:"used_times"`
}

// - base function

func (a *File) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *File) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *File) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *File) Context() echo.Context {
	return a.base.Context()
}

func (a *File) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *File) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *File) Namer() func(string) string {
	return a.base.Namer()
}

func (a *File) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *File) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *File) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *File) Objects() []*File {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *File) NewObjects() factory.Ranger {
	return &Slice_File{}
}

func (a *File) InitObjects() *[]*File {
	a.objects = []*File{}
	return &a.objects
}

func (a *File) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *File) Short_() string {
	return "file"
}

func (a *File) Struct_() string {
	return "File"
}

func (a *File) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *File) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *File) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *File) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *File) GroupBy(keyField string, inputRows ...[]*File) map[string][]*File {
	var rows []*File
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*File{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*File{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *File) KeyBy(keyField string, inputRows ...[]*File) map[string]*File {
	var rows []*File
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*File{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *File) AsKV(keyField string, valueField string, inputRows ...[]*File) map[string]interface{} {
	var rows []*File
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

func (a *File) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *File) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Type) == 0 {
		a.Type = "image"
	}
	if len(a.TableId) == 0 {
		a.TableId = "0"
	}
	err = DBI.Fire("creating", a, nil)
	if err != nil {
		return
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *File) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Type) == 0 {
		a.Type = "image"
	}
	if len(a.TableId) == 0 {
		a.TableId = "0"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *File) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *File) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["owner_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["owner_type"] = "user"
		}
	}
	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "image"
		}
	}
	if val, ok := kvset["table_id"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["table_id"] = "0"
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

func (a *File) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Type) == 0 {
			a.Type = "image"
		}
		if len(a.TableId) == 0 {
			a.TableId = "0"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Type) == 0 {
			a.Type = "image"
		}
		if len(a.TableId) == 0 {
			a.TableId = "0"
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *File) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *File) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *File) Reset() *File {
	a.Id = 0
	a.OwnerType = ``
	a.OwnerId = 0
	a.Name = ``
	a.SaveName = ``
	a.SavePath = ``
	a.ViewUrl = ``
	a.Ext = ``
	a.Mime = ``
	a.Type = ``
	a.Size = 0
	a.Width = 0
	a.Height = 0
	a.Dpi = 0
	a.Md5 = ``
	a.StorerName = ``
	a.StorerId = ``
	a.Created = 0
	a.Updated = 0
	a.Project = ``
	a.TableId = ``
	a.TableName = ``
	a.FieldName = ``
	a.Sort = 0
	a.Status = 0
	a.CategoryId = 0
	a.UsedTimes = 0
	return a
}

func (a *File) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["OwnerType"] = a.OwnerType
	r["OwnerId"] = a.OwnerId
	r["Name"] = a.Name
	r["SaveName"] = a.SaveName
	r["SavePath"] = a.SavePath
	r["ViewUrl"] = a.ViewUrl
	r["Ext"] = a.Ext
	r["Mime"] = a.Mime
	r["Type"] = a.Type
	r["Size"] = a.Size
	r["Width"] = a.Width
	r["Height"] = a.Height
	r["Dpi"] = a.Dpi
	r["Md5"] = a.Md5
	r["StorerName"] = a.StorerName
	r["StorerId"] = a.StorerId
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["Project"] = a.Project
	r["TableId"] = a.TableId
	r["TableName"] = a.TableName
	r["FieldName"] = a.FieldName
	r["Sort"] = a.Sort
	r["Status"] = a.Status
	r["CategoryId"] = a.CategoryId
	r["UsedTimes"] = a.UsedTimes
	return r
}

func (a *File) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "owner_type":
			a.OwnerType = param.AsString(value)
		case "owner_id":
			a.OwnerId = param.AsUint64(value)
		case "name":
			a.Name = param.AsString(value)
		case "save_name":
			a.SaveName = param.AsString(value)
		case "save_path":
			a.SavePath = param.AsString(value)
		case "view_url":
			a.ViewUrl = param.AsString(value)
		case "ext":
			a.Ext = param.AsString(value)
		case "mime":
			a.Mime = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "size":
			a.Size = param.AsUint64(value)
		case "width":
			a.Width = param.AsUint(value)
		case "height":
			a.Height = param.AsUint(value)
		case "dpi":
			a.Dpi = param.AsUint(value)
		case "md5":
			a.Md5 = param.AsString(value)
		case "storer_name":
			a.StorerName = param.AsString(value)
		case "storer_id":
			a.StorerId = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "project":
			a.Project = param.AsString(value)
		case "table_id":
			a.TableId = param.AsString(value)
		case "table_name":
			a.TableName = param.AsString(value)
		case "field_name":
			a.FieldName = param.AsString(value)
		case "sort":
			a.Sort = param.AsInt64(value)
		case "status":
			a.Status = param.AsInt(value)
		case "category_id":
			a.CategoryId = param.AsUint(value)
		case "used_times":
			a.UsedTimes = param.AsUint(value)
		}
	}
}

func (a *File) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "OwnerType":
			a.OwnerType = param.AsString(vv)
		case "OwnerId":
			a.OwnerId = param.AsUint64(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "SaveName":
			a.SaveName = param.AsString(vv)
		case "SavePath":
			a.SavePath = param.AsString(vv)
		case "ViewUrl":
			a.ViewUrl = param.AsString(vv)
		case "Ext":
			a.Ext = param.AsString(vv)
		case "Mime":
			a.Mime = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Size":
			a.Size = param.AsUint64(vv)
		case "Width":
			a.Width = param.AsUint(vv)
		case "Height":
			a.Height = param.AsUint(vv)
		case "Dpi":
			a.Dpi = param.AsUint(vv)
		case "Md5":
			a.Md5 = param.AsString(vv)
		case "StorerName":
			a.StorerName = param.AsString(vv)
		case "StorerId":
			a.StorerId = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Project":
			a.Project = param.AsString(vv)
		case "TableId":
			a.TableId = param.AsString(vv)
		case "TableName":
			a.TableName = param.AsString(vv)
		case "FieldName":
			a.FieldName = param.AsString(vv)
		case "Sort":
			a.Sort = param.AsInt64(vv)
		case "Status":
			a.Status = param.AsInt(vv)
		case "CategoryId":
			a.CategoryId = param.AsUint(vv)
		case "UsedTimes":
			a.UsedTimes = param.AsUint(vv)
		}
	}
}

func (a *File) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["owner_type"] = a.OwnerType
	r["owner_id"] = a.OwnerId
	r["name"] = a.Name
	r["save_name"] = a.SaveName
	r["save_path"] = a.SavePath
	r["view_url"] = a.ViewUrl
	r["ext"] = a.Ext
	r["mime"] = a.Mime
	r["type"] = a.Type
	r["size"] = a.Size
	r["width"] = a.Width
	r["height"] = a.Height
	r["dpi"] = a.Dpi
	r["md5"] = a.Md5
	r["storer_name"] = a.StorerName
	r["storer_id"] = a.StorerId
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["project"] = a.Project
	r["table_id"] = a.TableId
	r["table_name"] = a.TableName
	r["field_name"] = a.FieldName
	r["sort"] = a.Sort
	r["status"] = a.Status
	r["category_id"] = a.CategoryId
	r["used_times"] = a.UsedTimes
	return r
}

func (a *File) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *File) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
