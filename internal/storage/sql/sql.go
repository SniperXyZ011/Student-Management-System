package sql

import (
	"database/sql"
	"fmt"

	"github.com/SniperXyZ011/Student-Management-System/internal/config"
	"github.com/SniperXyZ011/Student-Management-System/internal/types"
	_ "github.com/go-sql-driver/mysql"
)

type Sql struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sql, error) {
	db, err := sql.Open("mysql", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100) UNIQUE,
		age INT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sql{
		Db: db,
	}, nil
}

func (s *Sql) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}	

func (s *Sql) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")

	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return  types.Student{}, err
		}
		return types.Student{}, fmt.Errorf("Query error: %s", err)
	}

	return  student, nil
}

func (s *Sql) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")

	if err != nil {
		return []types.Student{}, err
	}

	defer stmt.Close()

	var students []types.Student
	rows, err := stmt.Query()
	
	if err != nil {
		if err == sql.ErrNoRows {
			return  nil, fmt.Errorf("No students found")
		}
		return nil, fmt.Errorf("Query error: %s", err)
	}	

	defer rows.Close()

	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *Sql) DeleteStudent(id int64) (string, error) {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return "", err
	}	
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("no student found with id %d", id)
	}

	return "Deleted successfully", nil
}

func (s *Sql) EditStudent(student types.Student) (types.Student, error) {
	id := student.Id
	name := student.Name
	email := student.Email
	age := student.Age

	//we will first get the data of this id
	stmt, err := s.Db.Prepare("SELECT name, email, age FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student_query types.Student
	err = stmt.QueryRow(id).Scan(&student_query.Name, &student_query.Email, &student_query.Age)

	if err != nil {
		return types.Student{}, err
	}

	if(name == ""){
		name = student_query.Name
	}
	if(email == ""){
		email = student_query.Email
	}
	if(age == 0){
		age = student_query.Age
	}

	updateStmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id= ?")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	res, err := updateStmt.Exec(name, email, age, id)

	if err != nil {
		return types.Student{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return types.Student{}, fmt.Errorf("no student found with id %d", id)
	}

	return types.Student{
		Id: id,
		Name: name,
		Email: email,
		Age: age,
	}, nil
}