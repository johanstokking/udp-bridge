# UDP Bridge

This is a very simple UDP bridge that prints traffic sent between a client and a server.

## Configuration

Export the following environment variables:

- `LOCAL_ADDR`: the address to listen on (default `:1701`)
- `REMOTE_ADDR`: the address of the remote host 

## Usage

```
$ go run github.com/johanstokking/udp-bridge
```

## Example

The example output of a LoRa gateway running the Semtech UDP Packet Forwarder and a LoRaWAN Network Server:

```
> "\x01@\x0e\x02d\u007f\xda\xff\xfe\x00R`"
< "\x01@\x0e\x04"
> "\x01?~\x00d\u007f\xda\xff\xfe\x00R`{\"rxpk\":[{\"time\":\"2019-05-01T16:24:25.000000Z\",\"tmst\":323386195,\"chan\":2,\"rfch\":1,\"freq\":868.500000,\"stat\":1,\"modu\":\"LORA\",\"datr\":\"SF7BW125\",\"codr\":\"4/5\",\"lsnr\":7.0,\"rssi\":-47,\"size\":23,\"data\":\"AAAAAAAAAABCAAAAAAAAAEJeTGm74ag=\"}]}"
< "\x01?~\x01"
> "\x01\u078a\x02d\u007f\xda\xff\xfe\x00R`"
< "\x01\u078a\x04"
< "\x01\x00\x02\x03{\"txpk\":{\"imme\":false,\"tmst\":328386195,\"freq\":868.5,\"rfch\":0,\"powe\":14,\"modu\":\"LORA\",\"datr\":\"SF7BW125\",\"codr\":\"4/5\",\"ipol\":true,\"size\":33,\"ncrc\":true,\"data\":\"IPQA/dyvZL5kEryksf4iSrzB1oSVDSGNnSJdmHP6VpQD\"}}"
> "\x01\xd0T\x02d\u007f\xda\xff\xfe\x00R`"
< "\x01\xd0T\x04"
```
