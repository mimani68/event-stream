# Blockchain Notifier

## Development

```bash
ENv=production go run main.go
```

## Production

### Build
```bash
go build -o dist/event_stream main.go
```

### Run
In first step modify the `.env` file.
```bash
vi .env
```
After that run executable server
```bash
chmod +x ./dist/event_stream
ENv=production ./dist/event_stream
```