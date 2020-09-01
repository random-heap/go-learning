package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id int
	name string
	gender string
	age int
	position string
}

type DBConnect struct {
	DBUrl string
	DB *sql.DB
	student student
}


func main() {

	db := DBConnect{
		DBUrl:"root:root1243@tcp(127.0.0.1:3306)/test?charset=utf8mb4",
	}
	var err error
	db.DB, err = sql.Open("mysql", db.DBUrl)
	if err != nil {
		fmt.Printf("sql connect error : %s", err)
		return
	}
	fmt.Println(db.DBUrl)
	fmt.Println(db.DB.Stats())
	db.queryStudent()
	newStudent := &student{
		id:10003,
		name:"赵小海",
		gender:"男",
		age:22,
		position:"无",
	}
	db.insertStudent(newStudent)
	db.deleteStudent(10001)
	updateStudent := &student{
		id:10002,
		name:"刘小红",
		gender:"女",
		age:18,
		position:"团支书",
	}
	db.updateStudent(updateStudent)
}


func (db *DBConnect) queryStudent() {
	// pre用法
	//sqlPre, _ := db.DB.Prepare("select * from student where age > ? and age < ?")
	//students, err := sqlPre.Query(18, 22)
	students, err := db.DB.Query("select * from student")
	if err != nil{
		fmt.Printf("sql query error : %s", err)
	}
	defer students.Close()
	for i:= 0; students.Next(); i ++ {
		students.Scan(&db.student.id, &db.student.name, &db.student.gender, &db.student.age, &db.student.position,)
		fmt.Println("id:",db.student.id, "name:", db.student.name, "gender:", db.student.gender, "age", db.student.age, "position:", db.student.position)
	}
}

func (db *DBConnect) insertStudent(student *student)  {
	insertStu, _ := db.DB.Prepare("insert into student values(?,?,?,?,?)")
	defer insertStu.Close()
	result, err := insertStu.Exec(&student.id, &student.name, &student.gender, &student.age, &student.position)
	if err != nil {
		fmt.Printf("sql insert error : %s", err)
		return
	}
	rowNum, _ := result.RowsAffected()
	fmt.Printf("影响行数：%d\n", rowNum)
}

func (db *DBConnect) deleteStudent(id int){
	deleteStu, _ := db.DB.Prepare("delete from student where id = ?")
	defer deleteStu.Close()
	result, err := deleteStu.Exec(id)
	if err != nil {
		fmt.Printf("sql delete error : %s", err)
	}
	rowNum, _ := result.RowsAffected()
	fmt.Printf("影响行数：%d\n", rowNum)
}

func (db *DBConnect) updateStudent(newstudent *student)  {
	updateStu, _ := db.DB.Prepare("update student set name = ?, gender = ?, age = ?, position = ? where id = ?")
	defer updateStu.Close()
	result, err := updateStu.Exec(&newstudent.name, &newstudent.gender, &newstudent.age, &newstudent.position, &newstudent.id)
	if err != nil {
		fmt.Printf("sql insert error : %s", err)
		return
	}
	rowNum, _ := result.RowsAffected()
	fmt.Printf("影响行数：%d\n", rowNum)
}
