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

func newUserWallet(db *gorm.DB, opts ...gen.DOOption) userWallet {
	_userWallet := userWallet{}

	_userWallet.userWalletDo.UseDB(db, opts...)
	_userWallet.userWalletDo.UseModel(&model.UserWallet{})

	tableName := _userWallet.userWalletDo.TableName()
	_userWallet.ALL = field.NewAsterisk(tableName)
	_userWallet.ID = field.NewInt32(tableName, "id")
	_userWallet.UID = field.NewInt64(tableName, "uid")
	_userWallet.Address = field.NewString(tableName, "address")
	_userWallet.ChainID = field.NewInt32(tableName, "chain_id")
	_userWallet.ChainName = field.NewString(tableName, "chain_name")
	_userWallet.CreatedAt = field.NewTime(tableName, "created_at")
	_userWallet.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userWallet.IsDeleted = field.NewInt32(tableName, "is_deleted")

	_userWallet.fillFieldMap()

	return _userWallet
}

// userWallet 用户钱包的绑定
type userWallet struct {
	userWalletDo userWalletDo

	ALL       field.Asterisk
	ID        field.Int32
	UID       field.Int64
	Address   field.String // eth:42|solana:44|ton:33|btc:
	ChainID   field.Int32  // 非evm链，可以私下用负数指定
	ChainName field.String // 需要区分出测试网，例如eth-sepolia
	CreatedAt field.Time
	UpdatedAt field.Time
	IsDeleted field.Int32 // 0:no|1:yes

	fieldMap map[string]field.Expr
}

func (u userWallet) Table(newTableName string) *userWallet {
	u.userWalletDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userWallet) As(alias string) *userWallet {
	u.userWalletDo.DO = *(u.userWalletDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userWallet) updateTableName(table string) *userWallet {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.Address = field.NewString(table, "address")
	u.ChainID = field.NewInt32(table, "chain_id")
	u.ChainName = field.NewString(table, "chain_name")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.IsDeleted = field.NewInt32(table, "is_deleted")

	u.fillFieldMap()

	return u
}

func (u *userWallet) WithContext(ctx context.Context) IUserWalletDo {
	return u.userWalletDo.WithContext(ctx)
}

func (u userWallet) TableName() string { return u.userWalletDo.TableName() }

func (u userWallet) Alias() string { return u.userWalletDo.Alias() }

func (u userWallet) Columns(cols ...field.Expr) gen.Columns { return u.userWalletDo.Columns(cols...) }

func (u *userWallet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userWallet) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["address"] = u.Address
	u.fieldMap["chain_id"] = u.ChainID
	u.fieldMap["chain_name"] = u.ChainName
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["is_deleted"] = u.IsDeleted
}

func (u userWallet) clone(db *gorm.DB) userWallet {
	u.userWalletDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userWallet) replaceDB(db *gorm.DB) userWallet {
	u.userWalletDo.ReplaceDB(db)
	return u
}

type userWalletDo struct{ gen.DO }

type IUserWalletDo interface {
	gen.SubQuery
	Debug() IUserWalletDo
	WithContext(ctx context.Context) IUserWalletDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserWalletDo
	WriteDB() IUserWalletDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserWalletDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserWalletDo
	Not(conds ...gen.Condition) IUserWalletDo
	Or(conds ...gen.Condition) IUserWalletDo
	Select(conds ...field.Expr) IUserWalletDo
	Where(conds ...gen.Condition) IUserWalletDo
	Order(conds ...field.Expr) IUserWalletDo
	Distinct(cols ...field.Expr) IUserWalletDo
	Omit(cols ...field.Expr) IUserWalletDo
	Join(table schema.Tabler, on ...field.Expr) IUserWalletDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserWalletDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserWalletDo
	Group(cols ...field.Expr) IUserWalletDo
	Having(conds ...gen.Condition) IUserWalletDo
	Limit(limit int) IUserWalletDo
	Offset(offset int) IUserWalletDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserWalletDo
	Unscoped() IUserWalletDo
	Create(values ...*model.UserWallet) error
	CreateInBatches(values []*model.UserWallet, batchSize int) error
	Save(values ...*model.UserWallet) error
	First() (*model.UserWallet, error)
	Take() (*model.UserWallet, error)
	Last() (*model.UserWallet, error)
	Find() ([]*model.UserWallet, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserWallet, err error)
	FindInBatches(result *[]*model.UserWallet, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserWallet) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserWalletDo
	Assign(attrs ...field.AssignExpr) IUserWalletDo
	Joins(fields ...field.RelationField) IUserWalletDo
	Preload(fields ...field.RelationField) IUserWalletDo
	FirstOrInit() (*model.UserWallet, error)
	FirstOrCreate() (*model.UserWallet, error)
	FindByPage(offset int, limit int) (result []*model.UserWallet, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserWalletDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userWalletDo) Debug() IUserWalletDo {
	return u.withDO(u.DO.Debug())
}

func (u userWalletDo) WithContext(ctx context.Context) IUserWalletDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userWalletDo) ReadDB() IUserWalletDo {
	return u.Clauses(dbresolver.Read)
}

func (u userWalletDo) WriteDB() IUserWalletDo {
	return u.Clauses(dbresolver.Write)
}

func (u userWalletDo) Session(config *gorm.Session) IUserWalletDo {
	return u.withDO(u.DO.Session(config))
}

func (u userWalletDo) Clauses(conds ...clause.Expression) IUserWalletDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userWalletDo) Returning(value interface{}, columns ...string) IUserWalletDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userWalletDo) Not(conds ...gen.Condition) IUserWalletDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userWalletDo) Or(conds ...gen.Condition) IUserWalletDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userWalletDo) Select(conds ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userWalletDo) Where(conds ...gen.Condition) IUserWalletDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userWalletDo) Order(conds ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userWalletDo) Distinct(cols ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userWalletDo) Omit(cols ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userWalletDo) Join(table schema.Tabler, on ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userWalletDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userWalletDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userWalletDo) Group(cols ...field.Expr) IUserWalletDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userWalletDo) Having(conds ...gen.Condition) IUserWalletDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userWalletDo) Limit(limit int) IUserWalletDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userWalletDo) Offset(offset int) IUserWalletDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userWalletDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserWalletDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userWalletDo) Unscoped() IUserWalletDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userWalletDo) Create(values ...*model.UserWallet) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userWalletDo) CreateInBatches(values []*model.UserWallet, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userWalletDo) Save(values ...*model.UserWallet) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userWalletDo) First() (*model.UserWallet, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserWallet), nil
	}
}

func (u userWalletDo) Take() (*model.UserWallet, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserWallet), nil
	}
}

func (u userWalletDo) Last() (*model.UserWallet, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserWallet), nil
	}
}

func (u userWalletDo) Find() ([]*model.UserWallet, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserWallet), err
}

func (u userWalletDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserWallet, err error) {
	buf := make([]*model.UserWallet, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userWalletDo) FindInBatches(result *[]*model.UserWallet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userWalletDo) Attrs(attrs ...field.AssignExpr) IUserWalletDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userWalletDo) Assign(attrs ...field.AssignExpr) IUserWalletDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userWalletDo) Joins(fields ...field.RelationField) IUserWalletDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userWalletDo) Preload(fields ...field.RelationField) IUserWalletDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userWalletDo) FirstOrInit() (*model.UserWallet, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserWallet), nil
	}
}

func (u userWalletDo) FirstOrCreate() (*model.UserWallet, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserWallet), nil
	}
}

func (u userWalletDo) FindByPage(offset int, limit int) (result []*model.UserWallet, count int64, err error) {
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

func (u userWalletDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userWalletDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userWalletDo) Delete(models ...*model.UserWallet) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userWalletDo) withDO(do gen.Dao) *userWalletDo {
	u.DO = *do.(*gen.DO)
	return u
}
