package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "/data/users.db") // Ruta fija para evitar crear subcarpetas no deseadas
    if err != nil {
        log.Fatalf("Error al abrir la base de datos: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }

    // Crear tabla "usuarios" si no existe
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        surname TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`
    _, err = DB.Exec(createTableQuery)
    if err != nil {
        log.Fatalf("Error al crear la tabla 'usuarios': %v", err)
    }

    // Insertar datos iniciales si la tabla está vacía
    insertDataQuery := `
    INSERT INTO usuarios (name, surname, email)
    SELECT 'John', 'Doe', 'john.doe@example.com'
    WHERE NOT EXISTS (SELECT 1 FROM usuarios);
    `
    _, err = DB.Exec(insertDataQuery)
    if err != nil {
        log.Fatalf("Error al insertar datos iniciales en 'usuarios': %v", err)
    }
}

// GetDB devuelve la instancia de la base de datos
func GetDB() *sql.DB {
    return DB
}
