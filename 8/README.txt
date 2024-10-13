Для запуска сервера и бд нужно запустить на исполнение файл сервера 
После у нас будут подобные маршруты 
	
	
	mux.HandleFunc("/users", usersHandler) для запросов на добавление и вывод всех
	mux.HandleFunc("/users/id", usersFilteredByIdHolder) для поиска по 1ому параметру
	mux.HandleFunc("/users/name", usersFilteredByNameHolder) для поиска по 2ому параметру
	mux.HandleFunc("/user", userHandler)для любых модификаций по айди



все тесты в скрипте cl.go 
которые проверяют все маршруты с нормальными значениями и вызывают запрос с ошибкой для теста её обработчика