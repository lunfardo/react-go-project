package server

import (
	"apirest/Controllers"
	"apirest/Entities"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func getDatabaseConnection() *sql.DB {
	connStr := "host=10.5.0.5 user=postgres password=mysecretpassword dbname=my_arts_database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Serve() {
	fmt.Println("Connection to Database...")
	//depency injection to the entities
	conn := getDatabaseConnection()
	PainterDI := &Entities.PainterEntity
	PainterDI.AddDB(conn)
	PaintDI := &Entities.PaintEntity
	PaintDI.AddDB(conn)

	PaintersControllerCapsule := &Controllers.ResourceDemux{Controllers.PaintersController}
	PaintsControllerCapsule := &Controllers.ResourceDemux{Controllers.PaintsController}

	fmt.Println("Starting server...")
	// 	PaintersController := (*painters.PaintersController)(structs.MakeController(conn))

	mux := http.NewServeMux()
	PaintersControllerCapsule.AddRoutesToMux("/api/v1/painters", mux)
	PaintsControllerCapsule.AddRoutesToMux("/api/v1/paints", mux)

	http.ListenAndServe(":8080", mux)
}
