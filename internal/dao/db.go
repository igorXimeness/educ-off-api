package dao

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
)



//O ponteiro *pgxpool.Pool é usado para representar o pool de conexões.
//Ao retornar um ponteiro, a função ConnectDB() permite que o chamador
//utilize o mesmo pool de conexões em várias partes do código,
// compartilhando efetivamente o mesmo conjunto de conexões.

//windows, DB_USER = "Agatha"
func ConnectDB() (*pgxpool.Pool, error) {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    return pgxpool.Connect(context.Background(), dbURL)
}


// serv_cotuca = u22311 (servidor linux que hospeda o banco de dados)  

// docker = vou criar um sistema linux isolado dentro do meu sistema windows
// BD = hospedar dentro do sistema linux 

