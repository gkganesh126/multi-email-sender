
curl -X POST \
  http://localhost:8080/send-email \
  -H 'Content-Type: application/json' \
  -H 'X-Auth-Token: a2abe96296ce' \
  -d '{
	"provider": ["mailjet"],
    "from": {
        "name": "Ganesh",
        "emailID": "gkganesh127@gmail.com"
    },
    "to": [{
        "name": "Ganesh",
        "emailID": "gkganesh127@gmail.com"
    }],
    "cc": [{
        "name": "Ganesh",
        "emailID": "gkganesh127@gmail.com"
    }],
    "bcc": [{
        "name": "Ganesh",
        "emailID": "gkganesh127@gmail.com"
    }],
    "subject": "hey",
    "message": "dude" 
}'