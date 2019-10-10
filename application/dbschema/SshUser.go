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

type Slice_SshUser []*SshUser

func (s Slice_SshUser) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_SshUser) RangeRaw(fn func(m *SshUser) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// SshUser 数据库账号
type SshUser struct {
	base    factory.Base
	objects []*SshUser

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid         uint   `db:"uid" bson:"uid" comment:"UID" json:"uid" xml:"uid"`
	Host        string `db:"host" bson:"host" comment:"主机名" json:"host" xml:"host"`
	Port        int    `db:"port" bson:"port" comment:"端口" json:"port" xml:"port"`
	Charset     string `db:"charset" bson:"charset" comment:"字符集" json:"charset" xml:"charset"`
	Username    string `db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password    string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Name        string `db:"name" bson:"name" comment:"账号名称" json:"name" xml:"name"`
	Options     string `db:"options" bson:"options" comment:"其它选项(JSON)" json:"options" xml:"options"`
	PrivateKey  string `db:"private_key" bson:"private_key" comment:"私钥内容" json:"private_key" xml:"private_key"`
	Passphrase  string `db:"passphrase" bson:"passphrase" comment:"私钥口令" json:"passphrase" xml:"passphrase"`
	Protocol    string `db:"protocol" bson:"protocol" comment:"连接协议" json:"protocol" xml:"protocol"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	GroupId     uint   `db:"group_id" bson:"group_id" comment:"组" json:"group_id" xml:"group_id"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

// - base function

func (a *SshUser) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *SshUser) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *SshUser) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *SshUser) Context() echo.Context {
	return a.base.Context()
}

func (a *SshUser) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *SshUser) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *SshUser) Namer() func(string) string {
	return a.base.Namer()
}

func (a *SshUser) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *SshUser) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *SshUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *SshUser) Objects() []*SshUser {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *SshUser) NewObjects() factory.Ranger {
	return &Slice_SshUser{}
}

func (a *SshUser) InitObjects() *[]*SshUser {
	a.objects = []*SshUser{}
	return &a.objects
}

func (a *SshUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *SshUser) Short_() string {
	return "ssh_user"
}

func (a *SshUser) Struct_() string {
	return "SshUser"
}

func (a *SshUser) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *SshUser) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *SshUser) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *SshUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *SshUser) GroupBy(keyField string, inputRows ...[]*SshUser) map[string][]*SshUser {
	var rows []*SshUser
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*SshUser{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*SshUser{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *SshUser) KeyBy(keyField string, inputRows ...[]*SshUser) map[string]*SshUser {
	var rows []*SshUser
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*SshUser{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *SshUser) AsKV(keyField string, valueField string, inputRows ...[]*SshUser) map[string]interface{} {
	var rows []*SshUser
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

func (a *SshUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *SshUser) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Host) == 0 {
		a.Host = "localhost"
	}
	if len(a.Username) == 0 {
		a.Username = "root"
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

func (a *SshUser) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Host) == 0 {
		a.Host = "localhost"
	}
	if len(a.Username) == 0 {
		a.Username = "root"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *SshUser) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *SshUser) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["host"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["host"] = "localhost"
		}
	}
	if val, ok := kvset["username"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["username"] = "root"
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

func (a *SshUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Host) == 0 {
			a.Host = "localhost"
		}
		if len(a.Username) == 0 {
			a.Username = "root"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Host) == 0 {
			a.Host = "localhost"
		}
		if len(a.Username) == 0 {
			a.Username = "root"
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

func (a *SshUser) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *SshUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *SshUser) Reset() *SshUser {
	a.Id = 0
	a.Uid = 0
	a.Host = ``
	a.Port = 0
	a.Charset = ``
	a.Username = ``
	a.Password = ``
	a.Name = ``
	a.Options = ``
	a.PrivateKey = ``
	a.Passphrase = ``
	a.Protocol = ``
	a.Description = ``
	a.GroupId = 0
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *SshUser) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Uid"] = a.Uid
	r["Host"] = a.Host
	r["Port"] = a.Port
	r["Charset"] = a.Charset
	r["Username"] = a.Username
	r["Password"] = a.Password
	r["Name"] = a.Name
	r["Options"] = a.Options
	r["PrivateKey"] = a.PrivateKey
	r["Passphrase"] = a.Passphrase
	r["Protocol"] = a.Protocol
	r["Description"] = a.Description
	r["GroupId"] = a.GroupId
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	return r
}

func (a *SshUser) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "host":
			a.Host = param.AsString(value)
		case "port":
			a.Port = param.AsInt(value)
		case "charset":
			a.Charset = param.AsString(value)
		case "username":
			a.Username = param.AsString(value)
		case "password":
			a.Password = param.AsString(value)
		case "name":
			a.Name = param.AsString(value)
		case "options":
			a.Options = param.AsString(value)
		case "private_key":
			a.PrivateKey = param.AsString(value)
		case "passphrase":
			a.Passphrase = param.AsString(value)
		case "protocol":
			a.Protocol = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *SshUser) Set(key interface{}, value ...interface{}) {
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
		case "Host":
			a.Host = param.AsString(vv)
		case "Port":
			a.Port = param.AsInt(vv)
		case "Charset":
			a.Charset = param.AsString(vv)
		case "Username":
			a.Username = param.AsString(vv)
		case "Password":
			a.Password = param.AsString(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Options":
			a.Options = param.AsString(vv)
		case "PrivateKey":
			a.PrivateKey = param.AsString(vv)
		case "Passphrase":
			a.Passphrase = param.AsString(vv)
		case "Protocol":
			a.Protocol = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *SshUser) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["uid"] = a.Uid
	r["host"] = a.Host
	r["port"] = a.Port
	r["charset"] = a.Charset
	r["username"] = a.Username
	r["password"] = a.Password
	r["name"] = a.Name
	r["options"] = a.Options
	r["private_key"] = a.PrivateKey
	r["passphrase"] = a.Passphrase
	r["protocol"] = a.Protocol
	r["description"] = a.Description
	r["group_id"] = a.GroupId
	r["created"] = a.Created
	r["updated"] = a.Updated
	return r
}

func (a *SshUser) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *SshUser) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
