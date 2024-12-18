# Backend

Hola,
Para poder seguir en el proceso de selección se ha encargado un reto a realizar para medir tus capacidades. El objetivo principal es medir tu proactividad frente a los obstáculos que se te ponen en el cual se valorarán diferentes puntos:

    - Creatividad: Valoramos más la intención de como se quiere resolverlo más que el resultado en sí.
    - Buena práctica de programación: Modelo Vista Controllador(MVC), código limpio y óptimo (en inglés), buen uso y declaración de variables y funciones etc.
    - Documentación: Funcionalidad del proyecto y finalidad.

## BASE
Visual Studio Code -> Editor de código.

GIT -> Tencología de control de versiones.
https://learngitbranching.js.org/

Golang -> Lenguaje de programación.
https://go.dev/tour

Golang RESTful API -> Gin paquete básico para programar una API.
https://go.dev/doc/tutorial/web-service-gin

## RETO

El reto consiste en programar un RESTful API en Golang que hará una funcionalidad básica (GET, CREATE, UPDATE y DELETE) como por ejemplo con un modelo de datos de User(name,surname and email). Los usuarios serán almacenados en una base de datos que dejo a tú disposición elegir (te aconsejo SQLite, MySQL o MongoDB). 

Este proyecto tiene que ir acompañado por un docker-compose que adjunta el proyecto de Golang con la base de datos elegida.

Y por último, se analizarán las diferentes vulnerabilidades de la API y se va a proponer alguna solución para ello. Esta parte es más creativa ya que puede haber infinidad de casos como por ejemplo: para evitar ataques de Denegación se Servicio (DoS) se pretende adjuntar rate limits en los endpoints con el fin de limitar la cantidad de peticiones.

Suerte!