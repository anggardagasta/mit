## How To Run
- cd {YourGoPath}/src/github.com/anggardagasta
- git clone https://github.com/anggardagasta/mit.git
- cd mit
- go run main.go

In needs some package:
- go get -u github.com/asaskevich/govalidator
- go get -u github.com/go-chi/chi
- go get -u github.com/go-chi/cors
- go get -u github.com/go-sql-driver/mysql
- go get -u github.com/joho/godotenv

## Endpoint
Check User Phone
- Endpoint: http://localhost:8082/v1/users/check/phone [POST]
- Request: {"phone_number": "+6296533812343"}

Check User Email
- Endpoint: http://localhost:8082/v1/users/check/email [POST]
- Request: {"email": "yourmail@email.com"}

Register
- Endpoint: http://localhost:8082/v1/users/register [POST]
- Request: {"first_name":"First","last_name":"Last","phone_number":"+6285123456789","email": "yourmail@email.com","gender":"m","birth_date":"2000-01-01"}
