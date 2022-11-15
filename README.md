# Registration and Login
- [Installing](#installing)
- [Running Ip and Port](#IpAndPort)
- [Register And Login](#register&login)
- [Unittest](#unittest)
- [Jaeger](#jaeger)

## installing
    untuk menjalankan aplikasi ini sebagai berkut

```sh
    $ git clone https://github.com/rasyidmm/edufund.git
    $ cd edufund
    $ docker-compose up  
```
## IpAndPort
    Aplikasi berjalan pada IP 172.27.0.2:8080
    untuk DB berjalan pada IP 172.27.0.3:3306

## register&login
    dapat melakukan dengen postman
  - [Link Postman](https://github.com/rasyidmm/edufund/blob/master/example/edufund.postman_collection.json)
    
## unittest

![](example/register.png)
![](example/login.png)

## jaeger
    untuk melakukan tracing dan log dengan jaeger dapat menginstall jager pada docker dengan

```sh
    docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp  -p 6831:6831/udp   -p 6832:6832/udp -p 5778:5778  -p 16686:16686  -p 14268:14268 -p 14250:14250 -p 9411:9411 --restart unless-stopped jaegertracing/all-in-one:1.20
```
    buka browser http://localhost:16686/search