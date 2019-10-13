package models

import "github.com/lkhadi/MVC_golang-practice/views"

func Update(data views.UpdateArtikel) error {
	query, err := con.Query("UPDATE artikel SET judul=? ,artikel=? ,author=? ,tanggal=? WHERE id_artikel=?", data.Judul, data.Artikel, data.Author, data.Tanggal, data.IDArtikel)
	defer query.Close()
	if err != nil {
		return err
	}
	return nil
}
