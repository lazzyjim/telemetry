# Telemetry system
Backed system for getting data from sensors

---
### Check available CLI options:
```bash
go run ./cmd/sensor --help
```
```bash
go run ./cmd/sink --help
```
### Run:
1. Sink
```bash
go run ./cmd/sink \
  --bind=:8080 \
  --file=telemetry.log \
  --buffer=32768 \
  --flush=100ms \
  --rate-limit=1048576 \
  --encrypt-key=12345678901234567890123456789012
```
- --bind: address to listen on (host:port)

- --file: path to log file

- --buffer: size of in-memory buffer in bytes

- --flush: interval for automatic buffer flush

- --rate-limit: max input flow rate in bytes/sec

- --encrypt-key: optional 32-byte key for message encryption
2. Sensor
```bash
go run ./cmd/sensor \
  --rate=3 \
  --name=temp \
  --sink=http://localhost:8080/ingest
```

- --rate: messages per second

- --name: sensor name

- --sink: Sink HTTP endpoint