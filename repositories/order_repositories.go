package repositories

import (
	"database/sql"
	"product/common"
	"product/datamodels"
	"strconv"
)

type IOrderRepositories interface {
	Conn() error
	Insert(*datamodels.Order) (int64,error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

func NewOrderManagerRepository(table string, mysqlConn *sql.DB) IOrderRepositories {
	return &OrderManagerRepository{table:table, mysqlConn:mysqlConn}
}

type OrderManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

func (o *OrderManagerRepository) Insert(order *datamodels.Order) (id int64,err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "insert " + o.table + "set userID=?,productID=?,orderStatus=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return id, errStmt
	}
	result, errResult := stmt.Exec(order.ID, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return id, errResult
	}
	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(productID int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}
	sql := "delete from " + o.table + "where ID=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return false
	}
	_, err := stmt.Exec(productID)
	if err != nil{
		return false
	}
	return true
}

func (o *OrderManagerRepository) Update(order *datamodels.Order) error {
	if err := o.Conn(); err != nil {
		return err
	}
	sql := "update " + o.table + "set userID=?,productID=?,orderStatus=? where ID=" + strconv.FormatInt(order.ID, 10)
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return errStmt
	}
	_, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return errResult
	}
	return nil
}

func (o *OrderManagerRepository) SelectByKey(orderID int64) (order *datamodels.Order,err error) {
	if err = o.Conn(); err != nil {
		return &datamodels.Order{}, err
	}
	sql := "select * from " + o.table + "where ID=" + strconv.FormatInt(orderID, 10)
	row, errRow := o.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.Order{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, errRow
	}
	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}

func (o *OrderManagerRepository) SelectAll() (orderArray []*datamodels.Order,err error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "select * from " + o.table
	rows, errRows := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, errRows
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}
	for _, v := range result{
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v,order)
		orderArray = append(orderArray, order)
	}
	return
}

func (o *OrderManagerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}
	sql := "select o.ID,o.productID,o.orderStatus from ly.order as o left join product as p on o.productID=p.ID"
	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	return common.GetResultRows(rows), nil
}