// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"dexpert-event-listener/gorm/model"
)

func newListenerNewestBlocknumber(db *gorm.DB, opts ...gen.DOOption) listenerNewestBlocknumber {
	_listenerNewestBlocknumber := listenerNewestBlocknumber{}

	_listenerNewestBlocknumber.listenerNewestBlocknumberDo.UseDB(db, opts...)
	_listenerNewestBlocknumber.listenerNewestBlocknumberDo.UseModel(&model.ListenerNewestBlocknumber{})

	tableName := _listenerNewestBlocknumber.listenerNewestBlocknumberDo.TableName()
	_listenerNewestBlocknumber.ALL = field.NewAsterisk(tableName)
	_listenerNewestBlocknumber.ID = field.NewInt32(tableName, "id")
	_listenerNewestBlocknumber.ContractAddress = field.NewString(tableName, "contract_address")
	_listenerNewestBlocknumber.ChainID = field.NewInt32(tableName, "chain_id")
	_listenerNewestBlocknumber.BlockNumber = field.NewInt64(tableName, "block_number")
	_listenerNewestBlocknumber.CreatedAt = field.NewTime(tableName, "created_at")
	_listenerNewestBlocknumber.UpdatedAt = field.NewTime(tableName, "updated_at")

	_listenerNewestBlocknumber.fillFieldMap()

	return _listenerNewestBlocknumber
}

// listenerNewestBlocknumber 事件监听器区块记录表，存放监听器已监听的最新区块数
type listenerNewestBlocknumber struct {
	listenerNewestBlocknumberDo listenerNewestBlocknumberDo

	ALL             field.Asterisk
	ID              field.Int32
	ContractAddress field.String // 合约地址
	ChainID         field.Int32
	BlockNumber     field.Int64 // 最新爬到的区块
	CreatedAt       field.Time
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (l listenerNewestBlocknumber) Table(newTableName string) *listenerNewestBlocknumber {
	l.listenerNewestBlocknumberDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l listenerNewestBlocknumber) As(alias string) *listenerNewestBlocknumber {
	l.listenerNewestBlocknumberDo.DO = *(l.listenerNewestBlocknumberDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *listenerNewestBlocknumber) updateTableName(table string) *listenerNewestBlocknumber {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.ContractAddress = field.NewString(table, "contract_address")
	l.ChainID = field.NewInt32(table, "chain_id")
	l.BlockNumber = field.NewInt64(table, "block_number")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")

	l.fillFieldMap()

	return l
}

func (l *listenerNewestBlocknumber) WithContext(ctx context.Context) IListenerNewestBlocknumberDo {
	return l.listenerNewestBlocknumberDo.WithContext(ctx)
}

func (l listenerNewestBlocknumber) TableName() string {
	return l.listenerNewestBlocknumberDo.TableName()
}

func (l listenerNewestBlocknumber) Alias() string { return l.listenerNewestBlocknumberDo.Alias() }

func (l listenerNewestBlocknumber) Columns(cols ...field.Expr) gen.Columns {
	return l.listenerNewestBlocknumberDo.Columns(cols...)
}

func (l *listenerNewestBlocknumber) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *listenerNewestBlocknumber) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["id"] = l.ID
	l.fieldMap["contract_address"] = l.ContractAddress
	l.fieldMap["chain_id"] = l.ChainID
	l.fieldMap["block_number"] = l.BlockNumber
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
}

func (l listenerNewestBlocknumber) clone(db *gorm.DB) listenerNewestBlocknumber {
	l.listenerNewestBlocknumberDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l listenerNewestBlocknumber) replaceDB(db *gorm.DB) listenerNewestBlocknumber {
	l.listenerNewestBlocknumberDo.ReplaceDB(db)
	return l
}

type listenerNewestBlocknumberDo struct{ gen.DO }

type IListenerNewestBlocknumberDo interface {
	gen.SubQuery
	Debug() IListenerNewestBlocknumberDo
	WithContext(ctx context.Context) IListenerNewestBlocknumberDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IListenerNewestBlocknumberDo
	WriteDB() IListenerNewestBlocknumberDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IListenerNewestBlocknumberDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IListenerNewestBlocknumberDo
	Not(conds ...gen.Condition) IListenerNewestBlocknumberDo
	Or(conds ...gen.Condition) IListenerNewestBlocknumberDo
	Select(conds ...field.Expr) IListenerNewestBlocknumberDo
	Where(conds ...gen.Condition) IListenerNewestBlocknumberDo
	Order(conds ...field.Expr) IListenerNewestBlocknumberDo
	Distinct(cols ...field.Expr) IListenerNewestBlocknumberDo
	Omit(cols ...field.Expr) IListenerNewestBlocknumberDo
	Join(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo
	RightJoin(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo
	Group(cols ...field.Expr) IListenerNewestBlocknumberDo
	Having(conds ...gen.Condition) IListenerNewestBlocknumberDo
	Limit(limit int) IListenerNewestBlocknumberDo
	Offset(offset int) IListenerNewestBlocknumberDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IListenerNewestBlocknumberDo
	Unscoped() IListenerNewestBlocknumberDo
	Create(values ...*model.ListenerNewestBlocknumber) error
	CreateInBatches(values []*model.ListenerNewestBlocknumber, batchSize int) error
	Save(values ...*model.ListenerNewestBlocknumber) error
	First() (*model.ListenerNewestBlocknumber, error)
	Take() (*model.ListenerNewestBlocknumber, error)
	Last() (*model.ListenerNewestBlocknumber, error)
	Find() ([]*model.ListenerNewestBlocknumber, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ListenerNewestBlocknumber, err error)
	FindInBatches(result *[]*model.ListenerNewestBlocknumber, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ListenerNewestBlocknumber) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IListenerNewestBlocknumberDo
	Assign(attrs ...field.AssignExpr) IListenerNewestBlocknumberDo
	Joins(fields ...field.RelationField) IListenerNewestBlocknumberDo
	Preload(fields ...field.RelationField) IListenerNewestBlocknumberDo
	FirstOrInit() (*model.ListenerNewestBlocknumber, error)
	FirstOrCreate() (*model.ListenerNewestBlocknumber, error)
	FindByPage(offset int, limit int) (result []*model.ListenerNewestBlocknumber, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IListenerNewestBlocknumberDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l listenerNewestBlocknumberDo) Debug() IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Debug())
}

func (l listenerNewestBlocknumberDo) WithContext(ctx context.Context) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l listenerNewestBlocknumberDo) ReadDB() IListenerNewestBlocknumberDo {
	return l.Clauses(dbresolver.Read)
}

func (l listenerNewestBlocknumberDo) WriteDB() IListenerNewestBlocknumberDo {
	return l.Clauses(dbresolver.Write)
}

func (l listenerNewestBlocknumberDo) Session(config *gorm.Session) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Session(config))
}

func (l listenerNewestBlocknumberDo) Clauses(conds ...clause.Expression) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l listenerNewestBlocknumberDo) Returning(value interface{}, columns ...string) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l listenerNewestBlocknumberDo) Not(conds ...gen.Condition) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l listenerNewestBlocknumberDo) Or(conds ...gen.Condition) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l listenerNewestBlocknumberDo) Select(conds ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l listenerNewestBlocknumberDo) Where(conds ...gen.Condition) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l listenerNewestBlocknumberDo) Order(conds ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l listenerNewestBlocknumberDo) Distinct(cols ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l listenerNewestBlocknumberDo) Omit(cols ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l listenerNewestBlocknumberDo) Join(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l listenerNewestBlocknumberDo) LeftJoin(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l listenerNewestBlocknumberDo) RightJoin(table schema.Tabler, on ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l listenerNewestBlocknumberDo) Group(cols ...field.Expr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l listenerNewestBlocknumberDo) Having(conds ...gen.Condition) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l listenerNewestBlocknumberDo) Limit(limit int) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l listenerNewestBlocknumberDo) Offset(offset int) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l listenerNewestBlocknumberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l listenerNewestBlocknumberDo) Unscoped() IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Unscoped())
}

func (l listenerNewestBlocknumberDo) Create(values ...*model.ListenerNewestBlocknumber) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l listenerNewestBlocknumberDo) CreateInBatches(values []*model.ListenerNewestBlocknumber, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l listenerNewestBlocknumberDo) Save(values ...*model.ListenerNewestBlocknumber) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l listenerNewestBlocknumberDo) First() (*model.ListenerNewestBlocknumber, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListenerNewestBlocknumber), nil
	}
}

func (l listenerNewestBlocknumberDo) Take() (*model.ListenerNewestBlocknumber, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListenerNewestBlocknumber), nil
	}
}

func (l listenerNewestBlocknumberDo) Last() (*model.ListenerNewestBlocknumber, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListenerNewestBlocknumber), nil
	}
}

func (l listenerNewestBlocknumberDo) Find() ([]*model.ListenerNewestBlocknumber, error) {
	result, err := l.DO.Find()
	return result.([]*model.ListenerNewestBlocknumber), err
}

func (l listenerNewestBlocknumberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ListenerNewestBlocknumber, err error) {
	buf := make([]*model.ListenerNewestBlocknumber, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l listenerNewestBlocknumberDo) FindInBatches(result *[]*model.ListenerNewestBlocknumber, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l listenerNewestBlocknumberDo) Attrs(attrs ...field.AssignExpr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l listenerNewestBlocknumberDo) Assign(attrs ...field.AssignExpr) IListenerNewestBlocknumberDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l listenerNewestBlocknumberDo) Joins(fields ...field.RelationField) IListenerNewestBlocknumberDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l listenerNewestBlocknumberDo) Preload(fields ...field.RelationField) IListenerNewestBlocknumberDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l listenerNewestBlocknumberDo) FirstOrInit() (*model.ListenerNewestBlocknumber, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListenerNewestBlocknumber), nil
	}
}

func (l listenerNewestBlocknumberDo) FirstOrCreate() (*model.ListenerNewestBlocknumber, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListenerNewestBlocknumber), nil
	}
}

func (l listenerNewestBlocknumberDo) FindByPage(offset int, limit int) (result []*model.ListenerNewestBlocknumber, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l listenerNewestBlocknumberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l listenerNewestBlocknumberDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l listenerNewestBlocknumberDo) Delete(models ...*model.ListenerNewestBlocknumber) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *listenerNewestBlocknumberDo) withDO(do gen.Dao) *listenerNewestBlocknumberDo {
	l.DO = *do.(*gen.DO)
	return l
}
