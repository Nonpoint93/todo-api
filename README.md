POST /tasks

curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
    "title": "Comprar comida",
    "done": false
}'


GET /tasks

curl -X GET http://localhost:8080/tasks


#Mocks

Generated by:

└──╼ $mockgen -source=repositories/task_repository.go -destination=mocks/mock_task_repository.go -package=mocks