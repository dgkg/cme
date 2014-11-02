cme
===

[![wercker status](https://app.wercker.com/status/af308c0f5a0853e269c1cf9164cc06e6/m "wercker status")](https://app.wercker.com/project/bykey/af308c0f5a0853e269c1cf9164cc06e6)

Site de l'association des élèves de Maryse Éloy


## Production

[CME Production](http://www.la-communaute-de-maryse-eloy.com/)

## Testing

[CME Test](http://test.la-communaute-de-maryse-eloy.com/)

## Packages dependency

### All go get
go get github.com/gorilla/mux
go get github.com/gorilla/sessions
go get github.com/kennygrant/sanitize
go get github.com/nfnt/resize
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql

### Serveur Mux (router et dispatcheur)
[Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux/)
go get github.com/gorilla/mux

### Session
[Gorilla Session](http://www.gorillatoolkit.org/pkg/sessions)
go get github.com/gorilla/sessions

### Simplify text (url, filename...)
[Drivers MySQL](https://github.com/kennygrant/sanitize)
go get github.com/kennygrant/sanitize

### Image resize
[Image resize](github.com/nfnt/resize)
go get github.com/nfnt/resize

### ORM
[GORM ORM](https://github.com/jinzhu/gorm)
go get github.com/jinzhu/gorm

### Drivers MySQL 
[Drivers MySQL](github.com/go-sql-driver/mysql)
go get github.com/go-sql-driver/mysql
