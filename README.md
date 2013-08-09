rtb-server-dsp
==============

Implement RTB Server for DSP side.

# Example
    $ go run app.go -port 8080
    $ curl -i 'http://localhost:8080/ad?id=vVcz45AhEII77xApR5PNyv5Qn0vdWDcJ2y7vmkFpZEwMCI42EHwxiKYEwzsb59bQ&p=351000&s=1&t=10'
    HTTP/1.1 200 OK
    Content-Type: application/json
    Content-Length: 99
    Date: Fri, 09 Aug 2013 05:35:59 GMT

    {"id":"vVcz45AhEII77xApR5PNyv5Qn0vdWDcJ2y7vmkFpZEwMCI42EHwxiKYEwzsb59bQ","status":1,"price":351000}
