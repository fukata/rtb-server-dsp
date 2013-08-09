rtb-server-dsp
==============

Implement RTB Server for DSP side.

# Example
    $ go run app.go -port 8080
    $ curl -i -XPOST -d'id=vVcz45AhEII77xApR5PNyv5Qn0vdWDcJ2y7vmkFpZEwMCI42EHwxiKYEwzsb59bQ' -d'p=351000' -d's=1' -d't=1' 'http://localhost:8080/ad'
    HTTP/1.1 200 OK
    Content-Type: application/json
    Content-Length: 99
    Date: Fri, 09 Aug 2013 04:47:22 GMT
    
    {"id":"vVcz45AhEII77xApR5PNyv5Qn0vdWDcJ2y7vmkFpZEwMCI42EHwxiKYEwzsb59bQ","status":1,"price":351000}
