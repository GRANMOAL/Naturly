git add .
git commit -m "Ultimo Commit"
git push
go build main.go
del main.zip
Compress-Archive -Path .\main -DestinationPath main.zip

