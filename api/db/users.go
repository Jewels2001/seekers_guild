package db

type User struct {
	Id           int     `json:'id'`
	Name         string  `json:'name'`
	Discord_Name string  `json:'discord_name'`
	Discord_Id   string  `json:'discord_id'`
	Cur_Rank     string  `json:'cur_rank'`
	Prestige     float64 `json:'prestige'`
	Tokens       float64 `json:'tokens'`
}

func GetUsers() ([]*User, error) {
	data := make([]*User, 0)

	// Execute select all query
	rows, err := db.Query(get_all_users)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Scan each row
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id,
                        &u.Name,
                        &u.Discord_Name,
                        &u.Discord_Id,
                        &u.Cur_Rank,
                        &u.Prestige,
                        &u.Tokens,
                    )
		if err != nil {
			return data, err
		}
		data = append(data, &u)
	}
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, nil
}

func GetUser(id int) (*User, error) {
    // Execute get user by id query
    stmt, err := db.Prepare(get_user_by_id)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    var u User
    err = stmt.QueryRow(id).Scan(
        &u.Id,
        &u.Name,
        &u.Discord_Name,
        &u.Discord_Id,
        &u.Cur_Rank,
        &u.Prestige,
        &u.Tokens,
    )
    if err != nil {
        return nil, err
    }

    return &u, err
}
