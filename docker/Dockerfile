FROM swaggerapi/swagger-ui:v3.21.0
COPY ./src /tmp/

docker run -p 80:8080 -e SWAGGER_JSON=/mnt/swagger.yaml -v ~/Desktop:/mnt swaggerapi/swagger-ui