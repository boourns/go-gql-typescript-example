

package model

import (
	"github.com/boourns/dbutil"
  "database/sql"
  "fmt"
)

func sqlFieldsForUser() string {
  return "User.ID,User.Name" // ADD FIELD HERE
}

func loadUser(rows *sql.Rows) (*User, error) {
	ret := User{}

	err := rows.Scan(&ret.ID,&ret.Name) // ADD FIELD HERE
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func SelectUser(tx dbutil.DBLike, cond string, condFields ...interface{}) ([]*User, error) {
  ret := []*User{}
  sql := fmt.Sprintf("SELECT %s from User %s", sqlFieldsForUser(), cond)
	rows, err := tx.Query(sql, condFields...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
    item, err := loadUser(rows)
    if err != nil {
      return nil, err
    }
    ret = append(ret, item)
  }
  rows.Close()
  return ret, nil
}

func (s *User) Update(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare(fmt.Sprintf("UPDATE User SET ID=?,Name=? WHERE User.ID = ?", )) // ADD FIELD HERE

		if err != nil {
			return err
		}

    params := []interface{}{s.ID,s.Name} // ADD FIELD HERE
    params = append(params, s.ID)

		_, err = stmt.Exec(params...)
		if err != nil {
			return err
		}

    return nil
}

func (s *User) Insert(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare("INSERT INTO User(Name) VALUES(?)") // ADD FIELD HERE
		if err != nil {
			return err
		}

		result, err := stmt.Exec(s.Name) // ADD FIELD HERE
		if err != nil {
			return err
    }

    s.ID, err = result.LastInsertId()
    if err != nil {
      return err
    }
	  return nil
}

func (s *User) Delete(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare("DELETE FROM User WHERE ID = ?")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(s.ID)
		if err != nil {
			return err
    }

	  return nil
}

func CreateUserTable(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare(`



CREATE TABLE User (
  
    ID VARCHAR(255) PRIMARY KEY,
  
    Name VARCHAR(255)
  
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
