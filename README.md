# beego-mysql-api
```
# Create database
mysql -u root -p
mysql> source testdb.sql

# Generate code from mysql
bee generate appcode -driver=mysql -conn="root:test@tcp(127.0.0.1:3306)/testdb" -level=3

# Add main.go

# Generate Swagger Spec and UI
bee run -downdoc=true -gendoc=true

```
Need to change MySQL password to the real password in main.go:
```
 orm.RegisterDataBase("default", "mysql", "root:<my password>@/testdb?charset=utf8", 30)
``` 

Then the swagger UI is accessible at: http://localhost:8080/swagger/

N.B. table id needs to be integer.
