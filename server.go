package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/nakagami/firebirdsql"
)

type Rasp struct {
	NOTD       string
	NMPP       string
	RNAME      string
	EVEN_DAY   sql.NullString
	NOEVEN_DAY sql.NullString
	SATURDAY   sql.NullString
	SUNDAY     sql.NullString
}

var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query(`select np_otd.notd,n_mpp.nmpp,(select rname from room where room.id=it_rasp.room),it_rasp.even_day,it_rasp.noeven_day,saturday,sunday
	from it_rasp,np_otd,n_doc,n_mpp
	where (it_rasp.otd=np_otd.otd) and (it_rasp.doc=n_doc.doc) and (n_doc.mpp=n_mpp.mpp) and (n_doc.pv=1)
	order by it_rasp.lpu,notd,nmpp`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	raspisanie := []Rasp{}
	for rows.Next() {
		p := Rasp{}
		err := rows.Scan(&p.NOTD, &p.NMPP, &p.RNAME, &p.EVEN_DAY, &p.NOEVEN_DAY, &p.SATURDAY, &p.SUNDAY)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !p.EVEN_DAY.Valid {
			p.EVEN_DAY.String = ""
		}
		if !p.NOEVEN_DAY.Valid {
			p.NOEVEN_DAY.String = ""
		}
		if !p.SATURDAY.Valid {
			p.SATURDAY.String = ""
		}
		if !p.SUNDAY.Valid {
			p.SUNDAY.String = ""
		}

		raspisanie = append(raspisanie, p)
		fmt.Println(raspisanie)
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, raspisanie)
}

func main() {

	db, err := sql.Open("firebirdsql", "sysdba:masterkey@192.168.100.9/C:/DB/ARENA.GDB")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
