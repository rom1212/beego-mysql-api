# beego-mysql-api
```
# Create database
mysql -u root -p
mysql> source testdb.sql

# Generate code from mysql
bee generate appcode -driver=mysql -conn="root:i2test@tcp(127.0.0.1:3306)/testdb" -level=3

# Add main.go

# Generate Swagger Spec and UI
bee run -downdoc=true -gendoc=true

```
