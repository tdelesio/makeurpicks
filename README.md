Install MongoDB
brew tap mongodb/brew
brew install mongodb-community@5.0

Start MongoDB
brew services start mongodb/brew/mongodb-community

Init The Source Directory
cd /src/main
go mod init makeurpicks

Install GO Dependencies
go get github.com/google/uuid
go get go.mongodb.org/mongo-driver/mongo

MySQL (Not Needed ATM)
go get -u github.com/go-sql-driver/mysql
brew install mysql
brew services start mysql