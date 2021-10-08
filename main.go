package main

import (
	"image/png"
	"net/http"
	"text/template"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/qrcode/", viewCodeHandler)
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	p := Page{Title: "Qr Code Generator"}

	t, _ := template.ParseFiles("qrcode.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r * http.Request){
	dataString := r.FormValue("datastring")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 100, 100)

	png.Encode(w, qrCode)
}