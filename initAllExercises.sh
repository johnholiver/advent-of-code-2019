# Dumb initializer to all exercises
for i in {1..25}
do
  mkdir $i
  touch $i/main.go
  echo 'package main' >> $i/main.go
  echo 'import "fmt"' >> $i/main.go
  echo 'func main() {' >> $i/main.go
  echo 'fmt.Println("hello world - '$i'")' >> $i/main.go
  echo '}' >> $i/main.go
  go fmt $i/main.go
done
