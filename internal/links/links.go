package links

import (
	database "github.com/GodKimba/cuddly-golang-server/internal/pkg/db/mysql"
	"github.com/GodKimba/cuddly-golang-server/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// This function insert a link object into the database and return it's ID
func (link Link) Save() int64 {
	// Used prepare here before exec for security(?)
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUE(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	// Actually inserting
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	// Getting the id of the iserted link
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

// Function to get all the links from the database and pass it to the graphql server
func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
