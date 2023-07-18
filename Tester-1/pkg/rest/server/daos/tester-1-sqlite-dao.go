package daos

import (
	"database/sql"
	"errors"
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/daos/clients/sqls"
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type Tester1Dao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTester1s(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS tester1s(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Tester TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTester1Dao() (*Tester1Dao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTester1s(sqlClient)
	if err != nil {
		return nil, err
	}
	return &Tester1Dao{
		sqlClient,
	}, nil
}

func (tester1Dao *Tester1Dao) CreateTester1(m *models.Tester1) (*models.Tester1, error) {
	insertQuery := "INSERT INTO tester1s(Tester)values(?)"
	res, err := tester1Dao.sqlClient.DB.Exec(insertQuery, m.Tester)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("tester1 created")
	return m, nil
}

func (tester1Dao *Tester1Dao) UpdateTester1(id int64, m *models.Tester1) (*models.Tester1, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	tester1, err := tester1Dao.GetTester1(id)
	if err != nil {
		return nil, err
	}
	if tester1 == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE tester1s SET Tester = ? WHERE Id = ?"
	res, err := tester1Dao.sqlClient.DB.Exec(updateQuery, m.Tester, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("tester1 updated")
	return m, nil
}

func (tester1Dao *Tester1Dao) DeleteTester1(id int64) error {
	deleteQuery := "DELETE FROM tester1s WHERE Id = ?"
	res, err := tester1Dao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("tester1 deleted")
	return nil
}

func (tester1Dao *Tester1Dao) ListTester1s() ([]*models.Tester1, error) {
	selectQuery := "SELECT * FROM tester1s"
	rows, err := tester1Dao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var tester1s []*models.Tester1
	for rows.Next() {
		m := models.Tester1{}
		if err = rows.Scan(&m.Id, &m.Tester); err != nil {
			return nil, err
		}
		tester1s = append(tester1s, &m)
	}
	if tester1s == nil {
		tester1s = []*models.Tester1{}
	}

	log.Debugf("tester1 listed")
	return tester1s, nil
}

func (tester1Dao *Tester1Dao) GetTester1(id int64) (*models.Tester1, error) {
	selectQuery := "SELECT * FROM tester1s WHERE Id = ?"
	row := tester1Dao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Tester1{}
	if err := row.Scan(&m.Id, &m.Tester); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("tester1 retrieved")
	return &m, nil
}
