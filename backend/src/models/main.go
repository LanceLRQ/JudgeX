package main

import (
    "github.com/wejudge/wejudge-polygon/src/models/polygon"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    dsn := "root:12345678@tcp(127.0.0.1:13306)/polygon_test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    db.AutoMigrate(&polygon.Account{})
    db.AutoMigrate(&polygon.LoginLog{})
    db.AutoMigrate(&polygon.WeJudgeSSOBinding{})
    db.AutoMigrate(&polygon.ProblemTag{})
    db.AutoMigrate(&polygon.Problem{})
    db.AutoMigrate(&polygon.JudgeRecord{})
    db.AutoMigrate(&polygon.ApplicationClient{})
    db.AutoMigrate(&polygon.GlobalSettings{})
    db.AutoMigrate(&polygon.JudgeMachineNodeGroup{})
    db.AutoMigrate(&polygon.JudgeMachineNode{})
}