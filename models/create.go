package models

import "github.com/lkhadi/MVC_golang-practice/views"

func Artikel(data views.SetArtikel) error {
	query, err := con.Query("INSERT INTO artikel (artikel,judul,author,tanggal) VALUES(?,?,?,?)", data.Artikel, data.Judul, data.Author, data.Tanggal)
	defer query.Close()
	if err != nil {
		return err
	}
	return nil
}
