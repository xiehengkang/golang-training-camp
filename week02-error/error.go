package main

import (
	"database/sql"
	"errors"
	"fmt"
)
//可以在Dao层warp sql.ErrorNoRows抛给上层，在上层通过Is或As来做判定。

//定义哨兵变量
var NotFound = errors.New("data not found")

func Biz() error {
	err := Dao("")
	//包装后的错误，使用Is做判定
	//Is方法会Unwarp到Dao层封装的NotFound，不会追踪到sql.ErrNoRows
	if errors.Is(err,NotFound){
		//NotFound处理逻辑...
		return nil
	}

	if err != nil{
		//处理sql系统其他错误或抛给上层
	}
	//do something
	return nil
}


func Dao(query string)error  {
	//mockError指代调用sql
	//判断两种可能，sql.ErrorNoRows和其他错误
	err := mockError()
	if err == sql.ErrNoRows {
		//使用errors.Wrapf包装查询参数，带上了堆栈信息。
		return errors.Wrapf(Notfound,fmt.Sprintf("data node found, sql is :%s",query))


	}

	if err != nil{
		return errors.Wrapf(err,fmt.Sprintf("db query system error sql :%s",query))
	}
	//do something

	return nil

}

