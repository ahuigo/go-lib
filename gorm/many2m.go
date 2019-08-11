package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)
type User struct {
    gorm.Model
    Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    ID   int    `gorm:"primary_key"`
    Name string
    UserId int
    Users             []*User     `gorm:"many2many:languages;association_jointable_foreignkey:user_id;jointable_foreignkey:user_id;"` 
                /*
                association_jointable_foreignkey: 
                    JOIN ON language.user_id=users.id
                jointable_foreignkey: 
                    Where language.user_id in (2)
                HasMany:
                    association_foreignkey:ID;foreignkey:ID;
                */
}


func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  pf := fmt.Printf
  pf("'1'")
  if err != nil {
    println(err)
    fmt.Println(err)
    panic("连接数据库失败")
  }
  db.LogMode(true)
  //db.DropTable("users", "profiles")
  // 自动迁移模式
  db.AutoMigrate(&User{},&Language{})


    var users []User
    language := Language{}

    db.Create(&Language{Name:"en"})
    db.Create(&Language{Name:"zh",UserId:2})
    db.Create(&User{})
    //db.First(&language, "id = ?", 2)
    language = Language{ID:2}

    db.Model(&language).Related(&users,  "Users")
        //SELECT "users".* FROM "users" INNER JOIN "languages" ON "languages"."user_id" = "users"."id" WHERE ("languages"."user_id" IN ('2'))


    pf("user:%+v;\n:language:%+v\n", users, language)

  defer db.Close()
}
