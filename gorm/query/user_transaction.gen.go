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

func newUserTransaction(db *gorm.DB, opts ...gen.DOOption) userTransaction {
	_userTransaction := userTransaction{}

	_userTransaction.userTransactionDo.UseDB(db, opts...)
	_userTransaction.userTransactionDo.UseModel(&model.UserTransaction{})

	tableName := _userTransaction.userTransactionDo.TableName()
	_userTransaction.ALL = field.NewAsterisk(tableName)
	_userTransaction.ID = field.NewInt32(tableName, "id")
	_userTransaction.UID = field.NewInt64(tableName, "uid")
	_userTransaction.Tid = field.NewInt32(tableName, "tid")
	_userTransaction.SwapType = field.NewInt32(tableName, "swap_type")
	_userTransaction.Volume = field.NewString(tableName, "volume")
	_userTransaction.Timestamp = field.NewTime(tableName, "timestamp")
	_userTransaction.SwapName = field.NewString(tableName, "swap_name")
	_userTransaction.ChainName = field.NewString(tableName, "chain_name")
	_userTransaction.ChainID = field.NewInt32(tableName, "chain_id")
	_userTransaction.Fee = field.NewString(tableName, "fee")
	_userTransaction.FeeTokenSymbol = field.NewString(tableName, "fee_token_symbol")
	_userTransaction.FeeTokenDecimal = field.NewInt32(tableName, "fee_token_decimal")
	_userTransaction.IdentifyAddress = field.NewString(tableName, "identify_address")
	_userTransaction.CreatedAt = field.NewTime(tableName, "created_at")

	_userTransaction.fillFieldMap()

	return _userTransaction
}

type userTransaction struct {
	userTransactionDo userTransactionDo

	ALL             field.Asterisk
	ID              field.Int32  // 记录编号
	UID             field.Int64  // 用户id
	Tid             field.Int32  // 关联的table的 id
	SwapType        field.Int32  // 交易类型0:swap 1:sniper 99: launch
	Volume          field.String // 交易量
	Timestamp       field.Time   // 交易时间
	SwapName        field.String // 交易类型的名字
	ChainName       field.String // 链名
	ChainID         field.Int32  // 链id
	Fee             field.String // 手续费
	FeeTokenSymbol  field.String // 手续费symbol
	FeeTokenDecimal field.Int32  // 手续费token的decimal
	IdentifyAddress field.String // 可以是 user_launch_tx 表的 contract_address, 也可以是 user_swap_tx 的 tx, 用于构建唯一索引
	CreatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (u userTransaction) Table(newTableName string) *userTransaction {
	u.userTransactionDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTransaction) As(alias string) *userTransaction {
	u.userTransactionDo.DO = *(u.userTransactionDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTransaction) updateTableName(table string) *userTransaction {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.Tid = field.NewInt32(table, "tid")
	u.SwapType = field.NewInt32(table, "swap_type")
	u.Volume = field.NewString(table, "volume")
	u.Timestamp = field.NewTime(table, "timestamp")
	u.SwapName = field.NewString(table, "swap_name")
	u.ChainName = field.NewString(table, "chain_name")
	u.ChainID = field.NewInt32(table, "chain_id")
	u.Fee = field.NewString(table, "fee")
	u.FeeTokenSymbol = field.NewString(table, "fee_token_symbol")
	u.FeeTokenDecimal = field.NewInt32(table, "fee_token_decimal")
	u.IdentifyAddress = field.NewString(table, "identify_address")
	u.CreatedAt = field.NewTime(table, "created_at")

	u.fillFieldMap()

	return u
}

func (u *userTransaction) WithContext(ctx context.Context) IUserTransactionDo {
	return u.userTransactionDo.WithContext(ctx)
}

func (u userTransaction) TableName() string { return u.userTransactionDo.TableName() }

func (u userTransaction) Alias() string { return u.userTransactionDo.Alias() }

func (u userTransaction) Columns(cols ...field.Expr) gen.Columns {
	return u.userTransactionDo.Columns(cols...)
}

func (u *userTransaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTransaction) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 14)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["tid"] = u.Tid
	u.fieldMap["swap_type"] = u.SwapType
	u.fieldMap["volume"] = u.Volume
	u.fieldMap["timestamp"] = u.Timestamp
	u.fieldMap["swap_name"] = u.SwapName
	u.fieldMap["chain_name"] = u.ChainName
	u.fieldMap["chain_id"] = u.ChainID
	u.fieldMap["fee"] = u.Fee
	u.fieldMap["fee_token_symbol"] = u.FeeTokenSymbol
	u.fieldMap["fee_token_decimal"] = u.FeeTokenDecimal
	u.fieldMap["identify_address"] = u.IdentifyAddress
	u.fieldMap["created_at"] = u.CreatedAt
}

func (u userTransaction) clone(db *gorm.DB) userTransaction {
	u.userTransactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userTransaction) replaceDB(db *gorm.DB) userTransaction {
	u.userTransactionDo.ReplaceDB(db)
	return u
}

type userTransactionDo struct{ gen.DO }

type IUserTransactionDo interface {
	gen.SubQuery
	Debug() IUserTransactionDo
	WithContext(ctx context.Context) IUserTransactionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserTransactionDo
	WriteDB() IUserTransactionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserTransactionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserTransactionDo
	Not(conds ...gen.Condition) IUserTransactionDo
	Or(conds ...gen.Condition) IUserTransactionDo
	Select(conds ...field.Expr) IUserTransactionDo
	Where(conds ...gen.Condition) IUserTransactionDo
	Order(conds ...field.Expr) IUserTransactionDo
	Distinct(cols ...field.Expr) IUserTransactionDo
	Omit(cols ...field.Expr) IUserTransactionDo
	Join(table schema.Tabler, on ...field.Expr) IUserTransactionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserTransactionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserTransactionDo
	Group(cols ...field.Expr) IUserTransactionDo
	Having(conds ...gen.Condition) IUserTransactionDo
	Limit(limit int) IUserTransactionDo
	Offset(offset int) IUserTransactionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTransactionDo
	Unscoped() IUserTransactionDo
	Create(values ...*model.UserTransaction) error
	CreateInBatches(values []*model.UserTransaction, batchSize int) error
	Save(values ...*model.UserTransaction) error
	First() (*model.UserTransaction, error)
	Take() (*model.UserTransaction, error)
	Last() (*model.UserTransaction, error)
	Find() ([]*model.UserTransaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTransaction, err error)
	FindInBatches(result *[]*model.UserTransaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserTransaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserTransactionDo
	Assign(attrs ...field.AssignExpr) IUserTransactionDo
	Joins(fields ...field.RelationField) IUserTransactionDo
	Preload(fields ...field.RelationField) IUserTransactionDo
	FirstOrInit() (*model.UserTransaction, error)
	FirstOrCreate() (*model.UserTransaction, error)
	FindByPage(offset int, limit int) (result []*model.UserTransaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserTransactionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userTransactionDo) Debug() IUserTransactionDo {
	return u.withDO(u.DO.Debug())
}

func (u userTransactionDo) WithContext(ctx context.Context) IUserTransactionDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTransactionDo) ReadDB() IUserTransactionDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTransactionDo) WriteDB() IUserTransactionDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTransactionDo) Session(config *gorm.Session) IUserTransactionDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTransactionDo) Clauses(conds ...clause.Expression) IUserTransactionDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTransactionDo) Returning(value interface{}, columns ...string) IUserTransactionDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTransactionDo) Not(conds ...gen.Condition) IUserTransactionDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTransactionDo) Or(conds ...gen.Condition) IUserTransactionDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTransactionDo) Select(conds ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTransactionDo) Where(conds ...gen.Condition) IUserTransactionDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTransactionDo) Order(conds ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTransactionDo) Distinct(cols ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTransactionDo) Omit(cols ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTransactionDo) Join(table schema.Tabler, on ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTransactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTransactionDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTransactionDo) Group(cols ...field.Expr) IUserTransactionDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTransactionDo) Having(conds ...gen.Condition) IUserTransactionDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTransactionDo) Limit(limit int) IUserTransactionDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTransactionDo) Offset(offset int) IUserTransactionDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTransactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTransactionDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTransactionDo) Unscoped() IUserTransactionDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTransactionDo) Create(values ...*model.UserTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTransactionDo) CreateInBatches(values []*model.UserTransaction, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTransactionDo) Save(values ...*model.UserTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTransactionDo) First() (*model.UserTransaction, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTransaction), nil
	}
}

func (u userTransactionDo) Take() (*model.UserTransaction, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTransaction), nil
	}
}

func (u userTransactionDo) Last() (*model.UserTransaction, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTransaction), nil
	}
}

func (u userTransactionDo) Find() ([]*model.UserTransaction, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserTransaction), err
}

func (u userTransactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTransaction, err error) {
	buf := make([]*model.UserTransaction, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTransactionDo) FindInBatches(result *[]*model.UserTransaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTransactionDo) Attrs(attrs ...field.AssignExpr) IUserTransactionDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTransactionDo) Assign(attrs ...field.AssignExpr) IUserTransactionDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTransactionDo) Joins(fields ...field.RelationField) IUserTransactionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTransactionDo) Preload(fields ...field.RelationField) IUserTransactionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTransactionDo) FirstOrInit() (*model.UserTransaction, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTransaction), nil
	}
}

func (u userTransactionDo) FirstOrCreate() (*model.UserTransaction, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTransaction), nil
	}
}

func (u userTransactionDo) FindByPage(offset int, limit int) (result []*model.UserTransaction, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userTransactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTransactionDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTransactionDo) Delete(models ...*model.UserTransaction) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTransactionDo) withDO(do gen.Dao) *userTransactionDo {
	u.DO = *do.(*gen.DO)
	return u
}
