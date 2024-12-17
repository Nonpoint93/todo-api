POST /tasks

curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
    "title": "Comprar comida",
    "done": false
}'


GET /tasks

curl -X GET http://localhost:8080/tasks
