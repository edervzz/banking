# banking-go
Simple banking app

## Docker

| Comandos básicos contenedores                | Descripción                |
| -                                    | -                          |
| Descarga y lanza un contenedor desde una imagen (-d=detached) | $ docker run IMAGE [-d] |
| Listar los contenedores en ejecución (-a=todos) | $ docker ps [-a]           |
| Detener un contenedor | $ docker stop CONTAINER\|CONTAINER_ID |
| Remover un contenedor | $ docker rm CONTAINER\|CONTAINER_ID |

| Comandos Imagenes                | Descripción                |
| -                                    | -                          |
| Listar imagenes | $ docker images |
| Remover imagen| $ docker rmi IMAGE_ID |
| Solo descargar una imagen | $ docker pull IMAGE |

| Comandos avanzados contenedores                | Descripción                |
| -                                    | -                          |
| Espera 5s antes de lanzar  | $ docker run IMAGE sleep 5 |
| Ejecuta un comando en el contenedor | $ docker exec banking-app echo 'hello' |
| Adjunta un contenedor al 1er plano | $ docker attach CONTAINER_ID |
| Interative mode | $ docker run -it |
| Tag inidica la versión de la imagen, por defecto es latest | $ docker run IMAGE:TAG |
| Mapea un puerto externo a interno | $ docker run -p 80:8080 IMAGE |
| Mapea un volumen dentro de docker host hacia el contenedor, de esta forma el volumen queda separado de la imagen | $ docker run -v /opt/datadir:/var/lib/mysql/mysql IMAGE |
| Inspecciona un contenedor | $ docker inspect CONTAINER |