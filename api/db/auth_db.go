package db

import "database/sql"

func TokenActive(aid int) (bool, error) {
	// Execute check token query
    var auth_id string
    err := db.QueryRow(check_token, aid).Scan(&auth_id)
	if err != nil {
        if err != sql.ErrNoRows {
            return false, err
        }
		return false, nil
	}

	return true, nil
}

func AddToken(uid, aid string) (sql.Result, error) {
	// Execute insert token query
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare(insert_token)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(uid, aid)
	if err != nil {
		return res, err
	}
	err = tx.Commit()
	if err != nil {
		return res, err
	}

	return res, nil
}

func RemoveToken(id int) error {
	// Execute delete Query
	if _, err := db.Exec(delete_token, id); err != nil {
		return err
	}

	return nil
}
