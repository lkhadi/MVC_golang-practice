package models

func Artikel(artikel, judul, author, tanggal string) error {
	query, err := con.Query("INSERT INTO artikel (artikel,judul,author,tanggal) VALUES(?,?,?,?)", artikel, judul, author, tanggal)
	defer query.Close()
	if err != nil {
		return err
	}
	return nil
}
