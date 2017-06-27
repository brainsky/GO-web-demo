package models

import (
	"database/sql"
	"fmt"
)

//对数据库的CRUD的操作
type Cruder interface {
	AddCus(Customer)
	DeleteCus(int)
	Update(Customer)
	QueryCus(int) Customer
	QueryNamePass(string, string) Customer
}

//创建一个Customer类，字段与数据库一致
type Customer struct {
	Id       int    `form:"-"`
	Name     string `form:"name"  valid:"Required"`
	Password string `form:"password" valid:"Required"; MinSize(6)`
	Email    string `form:"email" valid:"Email; MaxSize(100)"`
	Phone    string `form:"phone" valid:"Mobile"`
}

//定义一个实现CRUD的类
type CruderImp struct {
	Con *sql.DB
}

//实现接口Cruder的增删改查操作
func (this *CruderImp) AddCus(c *Customer) {
	stm, err := this.Con.Prepare(`INSERT INTO customer (name,password,email,phone) values (?,?,?,?)`)
	defer stm.Close()
	checkerr(err)
	stm.Exec(c.Name, c.Password, c.Email, c.Phone)
}

//根据id删除用户
func (this *CruderImp) DeleteCus(id int) {
	stm, err := this.Con.Prepare(`DELETE FROM customer WHERE id=?`)
	defer stm.Close()
	checkerr(err)
	stm.Exec(id)
}

func (this *CruderImp) Update(c *Customer) {
	stm, err := this.Con.Prepare(`UPDATE customer SET name=?,password=?,phone=?,email=? WHERE id=?`)
	defer stm.Close()
	checkerr(err)
	stm.Exec(c.Name, c.Password, c.Phone, c.Email, c.Id)

}

func (this *CruderImp) QueryCus(id int) *Customer {
	//fmt.Println("QueryCus")
	rows, err := this.Con.Query(`select name,password,email,phone from customer where id=?`, id)
	checkerr(err)
	//fmt.Println("err")
	defer rows.Close()
	c := new(Customer)
	for rows.Next() {
		err := rows.Scan(&c.Name, &c.Password, &c.Email, &c.Phone)
		checkerr(err)
	}
	c.Id = id
	return c
}

func (this *CruderImp) QueryNamePass(email string, password string) *Customer {
	//fmt.Println("QueryNamePass")
	c := new(Customer)
	rows, err := this.Con.Query(`select name,password,email,phone from customer where email=? and password=?`, email, password)
	fmt.Println(rows)
	if err != nil {
		c = nil
		//fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.Name, &c.Password, &c.Email, &c.Phone)
		checkerr(err)
	}
	return c
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
