# Changelog
## 0.1.0
- Allows user to define maximum UDP packet size through `--udp-max-packet-size` (default **8192**) [#22](https://github.com/CotaPreco/Horus/issues/22);
- Removed `EXPOSE` from Dockerfile [#16](https://github.com/CotaPreco/Horus/issues/16);
- Fixed bug when message is `nil` before attempt to `.Send` (TaggedConnectionHub) [#14](https://github.com/CotaPreco/Horus/issues/14);
- Display details on startup (such as WebSocket `address:port` and UdpReceiver `address:port`) [#19](https://github.com/CotaPreco/Horus/issues/19).

## 0.1.0
- `UdpReceiver` & `NullByteReceiveStrategy`;
- `Message` *(usually a broadcast)*, `TaggedMessage` & `TagSequencedMessage` *(strictly must contain all tags)*;
- `Tag` now supports a few more special characters: `A-Z a-z 0-9 - . _ ~ ! $ & ' ( ) * + , ; = : @ / ?`.
