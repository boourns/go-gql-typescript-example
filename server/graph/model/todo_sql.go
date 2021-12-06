

package model

import (
	"github.com/boourns/dblib"
  "database/sql"
  "fmt"
)

func sqlFieldsForTodo() string {
  return "Todo.ID,Todo.Text,Todo.Done,Todo.UserID" // ADD FIELD HERE
}

func loadTodo(rows *sql.Rows) (*Todo, error) {
	ret := Todo{}

	err := rows.Scan(&ret.ID,&ret.Text,&ret.Done,&ret.UserID) // ADD FIELD HERE
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func SelectTodo(tx dblib.DBLike, cond string, condFields ...interface{}) ([]*Todo, error) {
  ret := []*Todo{}
  sql := fmt.Sprintf("SELECT %s from Todo %s", sqlFieldsForTodo(), cond)
	rows, err := tx.Query(sql, condFields...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
    item, err := loadTodo(rows)
    if err != nil {
      return nil, err
    }
    ret = append(ret, item)
  }
  rows.Close()
  return ret, nil
}

func (s *Todo) Update(tx dblib.DBLike) error {
		stmt, err := tx.Prepare(fmt.Sprintf("UPDATE Todo SET ID=?,Text=?,Done=?,UserID=? WHERE Todo.ID = ?", )) // ADD FIELD HERE

		if err != nil {
			return err
		}

    params := []interface{}{s.ID,s.Text,s.Done,s.UserID} // ADD FIELD HERE
    params = append(params, s.ID)

		_, err = stmt.Exec(params...)
		if err != nil {
			return err
		}

    return nil
}

func (s *Todo) Insert(tx dblib.DBLike) error {
		stmt, err := tx.Prepare("INSERT INTO Todo(Text,Done,UserID) VALUES(?,?,?)") // ADD FIELD HERE
		if err != nil {
			return err
		}

		result, err := stmt.Exec(s.Text,s.Done,s.UserID) // ADD FIELD HERE
		if err != nil {
			return err
    }

    s.ID, err = result.LastInsertId()
    if err != nil {
      return err
    }
	  return nil
}

func (s *Todo) Delete(tx dblib.DBLike) error {
		stmt, err := tx.Prepare("DELETE FROM Todo WHERE ID = ?")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(s.ID)
		if err != nil {
			return err
    }

	  return nil
}

func CreateTodoTable(tx dblib.DBLike) error {
		stmt, err := tx.Prepare(`



CREATE TABLE Todo (
  
    ID INTEGER PRIMARY KEY,
  
    Text VARCHAR(255),
  
    Done BOOLEAN,
  
    UserID INTEGER
  
);

`)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	  return nil
}
