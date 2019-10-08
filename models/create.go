package models

func Artikel() error {
	query, err := con.Query("INSERT INTO artikel (artikel,judul,author,tanggal) VALUES(?,?,?,?)", "ulalala", "judul1", "Lalu Kismara Hadi", "2019-10-09")
	defer query.Close()
	if err != nil {
		return err
	}
	return nil
}
