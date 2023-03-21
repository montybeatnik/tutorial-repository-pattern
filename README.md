# Repository tutorial 

## Curl examples 
```bash
# create a device
curl -X POST localhost:9080/new-device \
    -H 'Content-Type: application/json' \
    -d '{"hostname": "curl_test", "ip": "7.7.7.42"}'
# get a device by ip
curl -X GET localhost:9080/device?ip=7.7.7.7
curl -X GET localhost:9080/device
```