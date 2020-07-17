package services

// GetQuote gets a quote from the database
func GetQuote(id int) (string, error) {
	var quote string
	sql := "SELECT text FROM quote WHERE id = ?"
	err := dbm.QueryRow(sql, id).Scan(&quote)
	if err != nil {
		return "", err
	}
	return quote, nil
}
