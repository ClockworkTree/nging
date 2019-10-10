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

type Slice_CodeInvitation []*CodeInvitation

func (s Slice_CodeInvitation) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_CodeInvitation) RangeRaw(fn func(m *CodeInvitation) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// CodeInvitation 邀请码
type CodeInvitation struct {
	base    factory.Base
	objects []*CodeInvitation

	Id       uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid      uint   `db:"uid" bson:"uid" comment:"创建者" json:"uid" xml:"uid"`
	RecvUid  uint   `db:"recv_uid" bson:"recv_uid" comment:"使用者" json:"recv_uid" xml:"recv_uid"`
	Code     string `db:"code" bson:"code" comment:"邀请码" json:"code" xml:"code"`
	Created  uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Used     uint   `db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Start    uint   `db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End      uint   `db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	RoleIds  string `db:"role_ids" bson:"role_ids" comment:"注册为角色(多个用“,”分隔开)" json:"role_ids" xml:"role_ids"`
}

// - base function

func (a *CodeInvitation) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *CodeInvitation) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *CodeInvitation) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *CodeInvitation) Context() echo.Context {
	return a.base.Context()
}

func (a *CodeInvitation) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *CodeInvitation) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *CodeInvitation) Namer() func(string) string {
	return a.base.Namer()
}

func (a *CodeInvitation) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *CodeInvitation) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *CodeInvitation) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *CodeInvitation) Objects() []*CodeInvitation {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *CodeInvitation) NewObjects() factory.Ranger {
	return &Slice_CodeInvitation{}
}

func (a *CodeInvitation) InitObjects() *[]*CodeInvitation {
	a.objects = []*CodeInvitation{}
	return &a.objects
}

func (a *CodeInvitation) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *CodeInvitation) Short_() string {
	return "code_invitation"
}

func (a *CodeInvitation) Struct_() string {
	return "CodeInvitation"
}

func (a *CodeInvitation) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *CodeInvitation) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *CodeInvitation) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *CodeInvitation) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *CodeInvitation) GroupBy(keyField string, inputRows ...[]*CodeInvitation) map[string][]*CodeInvitation {
	var rows []*CodeInvitation
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*CodeInvitation{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CodeInvitation{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *CodeInvitation) KeyBy(keyField string, inputRows ...[]*CodeInvitation) map[string]*CodeInvitation {
	var rows []*CodeInvitation
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*CodeInvitation{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *CodeInvitation) AsKV(keyField string, valueField string, inputRows ...[]*CodeInvitation) map[string]interface{} {
	var rows []*CodeInvitation
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

func (a *CodeInvitation) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *CodeInvitation) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *CodeInvitation) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *CodeInvitation) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *CodeInvitation) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
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

func (a *CodeInvitation) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
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

func (a *CodeInvitation) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *CodeInvitation) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *CodeInvitation) Reset() *CodeInvitation {
	a.Id = 0
	a.Uid = 0
	a.RecvUid = 0
	a.Code = ``
	a.Created = 0
	a.Used = 0
	a.Start = 0
	a.End = 0
	a.Disabled = ``
	a.RoleIds = ``
	return a
}

func (a *CodeInvitation) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Uid"] = a.Uid
	r["RecvUid"] = a.RecvUid
	r["Code"] = a.Code
	r["Created"] = a.Created
	r["Used"] = a.Used
	r["Start"] = a.Start
	r["End"] = a.End
	r["Disabled"] = a.Disabled
	r["RoleIds"] = a.RoleIds
	return r
}

func (a *CodeInvitation) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "recv_uid":
			a.RecvUid = param.AsUint(value)
		case "code":
			a.Code = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "used":
			a.Used = param.AsUint(value)
		case "start":
			a.Start = param.AsUint(value)
		case "end":
			a.End = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "role_ids":
			a.RoleIds = param.AsString(value)
		}
	}
}

func (a *CodeInvitation) Set(key interface{}, value ...interface{}) {
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
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "RecvUid":
			a.RecvUid = param.AsUint(vv)
		case "Code":
			a.Code = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Used":
			a.Used = param.AsUint(vv)
		case "Start":
			a.Start = param.AsUint(vv)
		case "End":
			a.End = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "RoleIds":
			a.RoleIds = param.AsString(vv)
		}
	}
}

func (a *CodeInvitation) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["uid"] = a.Uid
	r["recv_uid"] = a.RecvUid
	r["code"] = a.Code
	r["created"] = a.Created
	r["used"] = a.Used
	r["start"] = a.Start
	r["end"] = a.End
	r["disabled"] = a.Disabled
	r["role_ids"] = a.RoleIds
	return r
}

func (a *CodeInvitation) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *CodeInvitation) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
