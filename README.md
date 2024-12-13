

https://courses.calhoun.io/lessons/les_wdv2_basic_web_app

To run test
[air](https://github.com/air-verse/air)

$ air init (create default .air.toml)
$ air main (main being main.go no .go necessary)

Go Mud 2025


## API Generation

Our goal is to write all API as documentation which is then generated using OAS2 (https://openapi-generator.tech/). 

To generate a new Ruby client from a valid api.yaml example

The localhost path is http://localhost:9999/swagger/petstore.yaml (openapi.json, etc in /doc/)

1. openapi-generator-cli generate -i petstore.yaml -g ruby
    *  https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml

2. Via docker
    * 
   ``` bash
   docker run --rm \
    -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/petstore.yaml \
    -g go \
    -o /local/out/go 
  
3. A go example
    * openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true \
    -o out -i petstore.yaml

4. npx example
    * npx @openapitools/openapi-generator-cli generate -i doc/user.yaml -g go -o /tmp/test/ 


