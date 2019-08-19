package main
import (
"database/sql"
    "fmt"
        _ "github.com/mattn/go-sqlite3"


)
func main() {
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего добавленного объекта
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}







 library, err := ioutil.ReadFile("archlinux_ru.txt")
         if err != nil {
		                 log.Fatalln(err)
				         }
					         archz := strings.Split(string(library), "\n")
						         lenArch:=len(archz)-1
							         for j := 0; j < lenArch; j++ {
									         arch:=strings.Split(archz[j],"|")
										         fmt.Println(arch[0],arch[1],arch[2],arch[3])
										 }

