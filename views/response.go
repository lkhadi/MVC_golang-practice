package views

type Response struct {
	Code int         "json:code"
	Body interface{} "json:body"
}

type Artikel struct {
	Judul   string "json:judul"
	Artikel string "json:artikel"
	Author  string "json:author"
	Tanggal string "json:tanggal"
}
