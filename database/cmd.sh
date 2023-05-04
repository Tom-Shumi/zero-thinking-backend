## go-migrate
# migrate create -ext sql -dir database/migrations -seq alter_thinking_tree_table
# migrate -database "mysql://root:root@tcp(localhost:3306)/zeroThinking" -path database/migrations up

## sqlboiler
# export GOPATH=$HOME/go
# export PATH=$PATH:$GOPATH/bin
# sqlboiler mysql -c sqlboiler.toml -o models -p models --no-tests --wipe