# REST API mit Go und Fiber

Diese REST API wurde mit der Go-Webframework Fiber erstellt und dient als Vorlage für die Entwicklung von RESTful Webanwendungen. Die API umfasst grundlegende Funktionen für das Verbinden mit einer Datenbank, das Hinzufügen von Middleware, die Dokumentation mit Swagger und das Management von HTTP-Anfragen.

## Middleware-Funktionalität

### Recovery Middleware
Die Recovery-Middleware ist eine zentrale Komponente, die die API vor Abstürzen schützt. Sie fängt Panikfälle (Panic) ab und sorgt dafür, dass die Anwendung nicht abstürzt. Im Falle eines Panikfalls gibt die Middleware einen Fehler zurück und zeichnet den Stacktrace des Panikfalls auf.

### Logging Middleware
Die Logging-Middleware wird verwendet, um alle eingehenden Anfragen und deren Reaktionszeiten zu protokollieren. Sie gibt Informationen über den Zeitpunkt der Anfrage, den Statuscode der Antwort, die HTTP-Methode, den Pfad und eventuelle Abfrageparameter aus. Die Protokolldatei wird in "Logfile.log" gespeichert.

### Authentifizierungs-Middleware
Die Authentifizierungs-Middleware ist für die Sicherstellung der API-Sicherheit verantwortlich. Sie überprüft die Authentifizierung des Benutzers, indem sie das JWT-Token im `Authorization`-Header erwartet. Sie prüft, ob das Token im Bearer-Token-Format vorliegt und validiert das Token, um sicherzustellen, dass der Benutzer authentifiziert ist. Bei einer fehlerhaften Authentifizierung wird eine Fehlermeldung zurückgegeben, andernfalls wird der Token im Kontext gespeichert.

## Verwendung

### API-Endpunkte
Die API enthält verschiedene Endpunkte, darunter:

- `/monitor`: Ein Endpunkt zur Überwachung und Überprüfung des Status der API.

- `/docs`: Hier finden Sie die Swagger-Dokumentation der API.

- `/swagger/*`: Die generierte Swagger-Oberfläche zur Interaktion mit der API.

- `/pfad`: Ein Endpunkt für POST-Anfragen, der eine Authentifizierung erfordert.

- `/pfad/:para`: Ein Endpunkt für GET-Anfragen mit einem optionalen Parameter, der ebenfalls eine Authentifizierung erfordert.

### Authentifizierung
Um die geschützten Endpunkte der API (z. B. `/pfad` und `/pfad/:para`) zu verwenden, ist eine Authentifizierung erforderlich. Dies kann mit einer Bearer-Token-Authentifizierung erreicht werden. Sie können Bibliotheken wie Gocloak oder Firebase verwenden, um die Authentifizierung zu implementieren.

## Anpassung
Sie können diese Vorlage nach Ihren Anforderungen anpassen und erweitern. Fügen Sie weitere Endpunkte hinzu, ändern Sie die Middleware oder integrieren Sie zusätzliche Funktionen in die API.

---

Happy coding! ✨🚀
