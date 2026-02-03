module github.com/lalka1231/todo-list1

go 1.21

require (
    github.com/go-chi/chi/v5 v5.2.4
    gorm.io/driver/postgres v1.6.0
    gorm.io/gorm v1.31.1
)

replace github.com/lalka1231/todo-list1/models => ./models
replace github.com/lalka1231/todo-list1/database => ./database
