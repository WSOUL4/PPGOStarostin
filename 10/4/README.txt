>openssl req -new -nodes -x509 -out server.pem -keyout server.key -days 365

>openssl req -new -nodes -x509 -out client.pem -keyout client.key -days 365




тут м создаём сервак чуть иным способом, для начала мы доп прописываем

cert, err := tls.LoadX509KeyPair("server.pem", "server.key")

config := tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAnyClientCert}
	config.Rand = rand.Reader

Для создания конфига с сертификатом, а tcp сервак создаётся не с net  библиотекой,а 
listen, err := tls.Listen(TYPE, HOST+":"+PORT, &config) 

и у нас сервер с сертификатом


Клиентже создаётся

config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true} //ТУТ ПРИКОЛ , НАШ СЕРТИФИКАТ НИФИГА НЕ НОМРАЛЬНЫЙ ПОЭТОМУ ПРОВЕРКУ ОТКЛЮЧИТЬ
conn, err := tls.Dial(TYPE, HOST+":"+PORT, &config)
тоже через тлс с ссылкой на конфиг


InsecureSkipVerify: нужна что бы твой конфиг был оправдан, тк ты можешь быть не подтверждённым автором сертификата


