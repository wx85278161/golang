# MySQL to be installed

function logging() {
    echo "";echo $1; echo "";
}

#Pre-requisite Mysql to be installed
logging "GET http://localhost:8099/"
curl -s -X GET http://localhost:8099/

logging "DELETE http://localhost:8099/"
curl -s -X DELETE http://localhost:8099/

typeset JSON_FILE=$(cat file.json)
logging "POST http://localhost:8099 -d \"$JSON_FILE\""
#curl -s -X POST http://localhost:8099/ -d '[{"Name":"Dave","Description":"Swimming"}]'
curl -s -X POST http://localhost:8099/ -d "$JSON_FILE"

logging "GET http://localhost:8099/"
curl -s -X GET http://localhost:8099/

logging "PUT http://localhost:8099/Cooking -d '[{\"Name\": \"Dave\"}]'"
curl -s -X PUT http://localhost:8099/Cooking -d '[{"Name":"Dave"}]'

logging "GET http://localhost:8099/"
curl -s -X GET http://localhost:8099/
