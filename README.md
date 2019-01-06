## filebeat output to pulsar

### install

#### install pulsar
Reference https://pulsar.apache.org/docs/en/standalone/

#### install pulsar go client
Reference https://pulsar.apache.org/docs/en/client-libraries-go/

If you encounter problems with dynamic libraries，please reference:https://pulsar.apache.org/docs/en/client-libraries-cpp/

#### install filebeat
```
go get github.com/elastic/beats
```

#### install filebeat-output-pulsar
Use root users
```
su
go get github.com/AmateurEvents/filebeat-ouput-pulsar
cd $GOPATH/src/github.com/AmateurEvents/filebeat-ouput-pulsar
go build
./filebeat-ouput-pulsar -e
```
