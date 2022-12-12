swagger:
	swag init -d .\cmd\,.\common\models\,.\pkg\teachers\.,.\pkg\timetable\. 

run: swagger
	go run .\cmd\main.go  

deploy: swagger
	cd ..
	docker-compose build
	docker-compose up
