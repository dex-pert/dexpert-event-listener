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

func newUserSwapTx(db *gorm.DB, opts ...gen.DOOption) userSwapTx {
	_userSwapTx := userSwapTx{}

	_userSwapTx.userSwapTxDo.UseDB(db, opts...)
	_userSwapTx.userSwapTxDo.UseModel(&model.UserSwapTx{})

	tableName := _userSwapTx.userSwapTxDo.TableName()
	_userSwapTx.ALL = field.NewAsterisk(tableName)
	_userSwapTx.ID = field.NewInt32(tableName, "id")
	_userSwapTx.UID = field.NewInt64(tableName, "uid")
	_userSwapTx.Address = field.NewString(tableName, "address")
	_userSwapTx.Tx = field.NewString(tableName, "tx")
	_userSwapTx.TokenSymbol = field.NewString(tableName, "token_symbol")
	_userSwapTx.AmountIn = field.NewString(tableName, "amount_in")
	_userSwapTx.TokenIn = field.NewString(tableName, "token_in")
	_userSwapTx.TransactionTime = field.NewTime(tableName, "transaction_time")
	_userSwapTx.Fee = field.NewString(tableName, "fee")
	_userSwapTx.BlockNumber = field.NewInt64(tableName, "block_number")
	_userSwapTx.ChainID = field.NewInt32(tableName, "chain_id")
	_userSwapTx.ChainName = field.NewString(tableName, "chain_name")
	_userSwapTx.SwapType = field.NewInt32(tableName, "swap_type")
	_userSwapTx.TokenName = field.NewString(tableName, "token_name")
	_userSwapTx.Decimal = field.NewInt32(tableName, "decimal")
	_userSwapTx.FeeTokenName = field.NewString(tableName, "fee_token_name")
	_userSwapTx.FeeDecimal = field.NewInt32(tableName, "fee_decimal")
	_userSwapTx.FeeTokenSymbol = field.NewString(tableName, "fee_token_symbol")
	_userSwapTx.CreatedAt = field.NewTime(tableName, "created_at")

	_userSwapTx.fillFieldMap()

	return _userSwapTx
}

// userSwapTx 记录用户swap产生的交易
type userSwapTx struct {
	userSwapTxDo userSwapTxDo

	ALL             field.Asterisk
	ID              field.Int32
	UID             field.Int64  // 用户编号
	Address         field.String // 引用user_wallet表的address
	Tx              field.String // 交易哈希值
	TokenSymbol     field.String // 第一个代币的名称
	AmountIn        field.String // 第一个代币的数量
	TokenIn         field.String // 第一个代币的地址
	TransactionTime field.Time   // 时间
	Fee             field.String // 手续费
	BlockNumber     field.Int64  // 区块Id
	ChainID         field.Int32  // 链id
	ChainName       field.String // 链名
	SwapType        field.Int32  // 交易类型0：swap，1：Sniper
	TokenName       field.String // 代币名称
	Decimal         field.Int32
	FeeTokenName    field.String // 费率token
	FeeDecimal      field.Int32
	FeeTokenSymbol  field.String
	CreatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (u userSwapTx) Table(newTableName string) *userSwapTx {
	u.userSwapTxDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userSwapTx) As(alias string) *userSwapTx {
	u.userSwapTxDo.DO = *(u.userSwapTxDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userSwapTx) updateTableName(table string) *userSwapTx {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.Address = field.NewString(table, "address")
	u.Tx = field.NewString(table, "tx")
	u.TokenSymbol = field.NewString(table, "token_symbol")
	u.AmountIn = field.NewString(table, "amount_in")
	u.TokenIn = field.NewString(table, "token_in")
	u.TransactionTime = field.NewTime(table, "transaction_time")
	u.Fee = field.NewString(table, "fee")
	u.BlockNumber = field.NewInt64(table, "block_number")
	u.ChainID = field.NewInt32(table, "chain_id")
	u.ChainName = field.NewString(table, "chain_name")
	u.SwapType = field.NewInt32(table, "swap_type")
	u.TokenName = field.NewString(table, "token_name")
	u.Decimal = field.NewInt32(table, "decimal")
	u.FeeTokenName = field.NewString(table, "fee_token_name")
	u.FeeDecimal = field.NewInt32(table, "fee_decimal")
	u.FeeTokenSymbol = field.NewString(table, "fee_token_symbol")
	u.CreatedAt = field.NewTime(table, "created_at")

	u.fillFieldMap()

	return u
}

func (u *userSwapTx) WithContext(ctx context.Context) IUserSwapTxDo {
	return u.userSwapTxDo.WithContext(ctx)
}

func (u userSwapTx) TableName() string { return u.userSwapTxDo.TableName() }

func (u userSwapTx) Alias() string { return u.userSwapTxDo.Alias() }

func (u userSwapTx) Columns(cols ...field.Expr) gen.Columns { return u.userSwapTxDo.Columns(cols...) }

func (u *userSwapTx) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userSwapTx) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 19)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["address"] = u.Address
	u.fieldMap["tx"] = u.Tx
	u.fieldMap["token_symbol"] = u.TokenSymbol
	u.fieldMap["amount_in"] = u.AmountIn
	u.fieldMap["token_in"] = u.TokenIn
	u.fieldMap["transaction_time"] = u.TransactionTime
	u.fieldMap["fee"] = u.Fee
	u.fieldMap["block_number"] = u.BlockNumber
	u.fieldMap["chain_id"] = u.ChainID
	u.fieldMap["chain_name"] = u.ChainName
	u.fieldMap["swap_type"] = u.SwapType
	u.fieldMap["token_name"] = u.TokenName
	u.fieldMap["decimal"] = u.Decimal
	u.fieldMap["fee_token_name"] = u.FeeTokenName
	u.fieldMap["fee_decimal"] = u.FeeDecimal
	u.fieldMap["fee_token_symbol"] = u.FeeTokenSymbol
	u.fieldMap["created_at"] = u.CreatedAt
}

func (u userSwapTx) clone(db *gorm.DB) userSwapTx {
	u.userSwapTxDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userSwapTx) replaceDB(db *gorm.DB) userSwapTx {
	u.userSwapTxDo.ReplaceDB(db)
	return u
}

type userSwapTxDo struct{ gen.DO }

type IUserSwapTxDo interface {
	gen.SubQuery
	Debug() IUserSwapTxDo
	WithContext(ctx context.Context) IUserSwapTxDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserSwapTxDo
	WriteDB() IUserSwapTxDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserSwapTxDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserSwapTxDo
	Not(conds ...gen.Condition) IUserSwapTxDo
	Or(conds ...gen.Condition) IUserSwapTxDo
	Select(conds ...field.Expr) IUserSwapTxDo
	Where(conds ...gen.Condition) IUserSwapTxDo
	Order(conds ...field.Expr) IUserSwapTxDo
	Distinct(cols ...field.Expr) IUserSwapTxDo
	Omit(cols ...field.Expr) IUserSwapTxDo
	Join(table schema.Tabler, on ...field.Expr) IUserSwapTxDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserSwapTxDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserSwapTxDo
	Group(cols ...field.Expr) IUserSwapTxDo
	Having(conds ...gen.Condition) IUserSwapTxDo
	Limit(limit int) IUserSwapTxDo
	Offset(offset int) IUserSwapTxDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserSwapTxDo
	Unscoped() IUserSwapTxDo
	Create(values ...*model.UserSwapTx) error
	CreateInBatches(values []*model.UserSwapTx, batchSize int) error
	Save(values ...*model.UserSwapTx) error
	First() (*model.UserSwapTx, error)
	Take() (*model.UserSwapTx, error)
	Last() (*model.UserSwapTx, error)
	Find() ([]*model.UserSwapTx, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserSwapTx, err error)
	FindInBatches(result *[]*model.UserSwapTx, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserSwapTx) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserSwapTxDo
	Assign(attrs ...field.AssignExpr) IUserSwapTxDo
	Joins(fields ...field.RelationField) IUserSwapTxDo
	Preload(fields ...field.RelationField) IUserSwapTxDo
	FirstOrInit() (*model.UserSwapTx, error)
	FirstOrCreate() (*model.UserSwapTx, error)
	FindByPage(offset int, limit int) (result []*model.UserSwapTx, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserSwapTxDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userSwapTxDo) Debug() IUserSwapTxDo {
	return u.withDO(u.DO.Debug())
}

func (u userSwapTxDo) WithContext(ctx context.Context) IUserSwapTxDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userSwapTxDo) ReadDB() IUserSwapTxDo {
	return u.Clauses(dbresolver.Read)
}

func (u userSwapTxDo) WriteDB() IUserSwapTxDo {
	return u.Clauses(dbresolver.Write)
}

func (u userSwapTxDo) Session(config *gorm.Session) IUserSwapTxDo {
	return u.withDO(u.DO.Session(config))
}

func (u userSwapTxDo) Clauses(conds ...clause.Expression) IUserSwapTxDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userSwapTxDo) Returning(value interface{}, columns ...string) IUserSwapTxDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userSwapTxDo) Not(conds ...gen.Condition) IUserSwapTxDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userSwapTxDo) Or(conds ...gen.Condition) IUserSwapTxDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userSwapTxDo) Select(conds ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userSwapTxDo) Where(conds ...gen.Condition) IUserSwapTxDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userSwapTxDo) Order(conds ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userSwapTxDo) Distinct(cols ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userSwapTxDo) Omit(cols ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userSwapTxDo) Join(table schema.Tabler, on ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userSwapTxDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userSwapTxDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userSwapTxDo) Group(cols ...field.Expr) IUserSwapTxDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userSwapTxDo) Having(conds ...gen.Condition) IUserSwapTxDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userSwapTxDo) Limit(limit int) IUserSwapTxDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userSwapTxDo) Offset(offset int) IUserSwapTxDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userSwapTxDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserSwapTxDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userSwapTxDo) Unscoped() IUserSwapTxDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userSwapTxDo) Create(values ...*model.UserSwapTx) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userSwapTxDo) CreateInBatches(values []*model.UserSwapTx, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userSwapTxDo) Save(values ...*model.UserSwapTx) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userSwapTxDo) First() (*model.UserSwapTx, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserSwapTx), nil
	}
}

func (u userSwapTxDo) Take() (*model.UserSwapTx, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserSwapTx), nil
	}
}

func (u userSwapTxDo) Last() (*model.UserSwapTx, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserSwapTx), nil
	}
}

func (u userSwapTxDo) Find() ([]*model.UserSwapTx, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserSwapTx), err
}

func (u userSwapTxDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserSwapTx, err error) {
	buf := make([]*model.UserSwapTx, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userSwapTxDo) FindInBatches(result *[]*model.UserSwapTx, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userSwapTxDo) Attrs(attrs ...field.AssignExpr) IUserSwapTxDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userSwapTxDo) Assign(attrs ...field.AssignExpr) IUserSwapTxDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userSwapTxDo) Joins(fields ...field.RelationField) IUserSwapTxDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userSwapTxDo) Preload(fields ...field.RelationField) IUserSwapTxDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userSwapTxDo) FirstOrInit() (*model.UserSwapTx, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserSwapTx), nil
	}
}

func (u userSwapTxDo) FirstOrCreate() (*model.UserSwapTx, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserSwapTx), nil
	}
}

func (u userSwapTxDo) FindByPage(offset int, limit int) (result []*model.UserSwapTx, count int64, err error) {
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

func (u userSwapTxDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userSwapTxDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userSwapTxDo) Delete(models ...*model.UserSwapTx) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userSwapTxDo) withDO(do gen.Dao) *userSwapTxDo {
	u.DO = *do.(*gen.DO)
	return u
}
