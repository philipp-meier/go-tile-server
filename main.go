package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db    *sql.DB
	dbErr error
)

const tileQuery = `
	SELECT tile_data
	FROM tiles
	WHERE zoom_level = ? AND tile_column = ? AND tile_row = ?
	LIMIT 1
`

func main() {
	db, dbErr = sql.Open("sqlite3", "./data/map_data.mbtiles")
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/tiles/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}", getTileImage).Methods("GET")

	log.Fatal(http.ListenAndServe(":5296", router))
}

func getTileImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	z, _ := strconv.Atoi(params["z"])
	x, _ := strconv.Atoi(params["x"])
	y, _ := strconv.Atoi(params["y"])

	tileBytes, err := queryTileData(db, z, x, y)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Cache-Control", "max-age=58362, stale-while-revalidate=604800, stale-if-error=604800")
	w.Header().Add("Content-Type", "image/png")
	w.Header().Add("Content-Length", strconv.Itoa(len(tileBytes)))

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Error code: 900001", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(tileBytes)
	flusher.Flush()
}

func queryTileData(db *sql.DB, z int, x int, y int) ([]byte, error) {
	b := make([]byte, 16)
	err := db.QueryRow(tileQuery, z, x, y).Scan(&b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
