## Onboard Service
### Insert Message
```
grpcurl -plaintext -d \
'{"body": "test"}' \
localhost:8081 \
MessageService/InsertMessage
```
### Get All Messages
```
grpcurl -plaintext \
localhost:8081 \
MessageService/GetAllMessages
```