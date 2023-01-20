package db

func UpdatePrestige(id int, prestige_change float64) (float64, error) {
	// Execute prestige update query
	if _, err := db.Exec(update_prestige, prestige_change, id); err != nil {
		return 0.0, err
	}

	// Get updated user
	user, err := GetUser(id)
	if err != nil {
		return 0.0, err
	}

	return user.Prestige, nil
}

func UpdateTokens(id int, tokens_change float64) (float64, error) {
	// Execute prestige update query
	if _, err := db.Exec(update_tokens, tokens_change, id); err != nil {
		return 0.0, err
	}

	// Get updated user
	user, err := GetUser(id)
	if err != nil {
		return 0.0, err
	}

	return user.Tokens, nil
}
