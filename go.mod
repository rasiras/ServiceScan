module example.com/myproject

go 1.16

require (
    github.com/gorilla/securecookie v1.1.1 // Vulnerable version
    github.com/gorilla/mux v1.6.2 // Vulnerable version
    github.com/go-sql-driver/mysql v1.5.0 // Vulnerable version
    golang.org/x/text v0.3.0 // Vulnerable version
    golang.org/x/net v0.0.0-20230101000000-abcdefabcdef // Non-vulnerable example
)
