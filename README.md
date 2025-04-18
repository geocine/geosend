# geosend

`geosend` is a command-line tool for secure file and folder transfers between computers. It uses [croc](https://github.com/schollz/croc), extending it to support JSON-based relay configuration. Relay settings are embedded during build but can be overridden at runtime.

## Features  
- Secure, end-to-end encrypted transfers  
- Send or receive files and folders  
- Multiplexed connections for faster transfer (multiple ports per relay)  
- Built-in relay configuration, with optional override via `relays.json`

## Usage

### Send
```sh
geosend send [file0] [file1] ...
```
### Receive
```sh
geosend receive [code]
```

## Relay Configuration

The `relays.json` configuration will be embedded at build time.

Example:
```json
{
  "relays": [
    {
      "address": "localhost",
      "password": "geocine",
      "ports": "12347"
    }
  ]
}
```

- `address`: Hostname or IP of the relay (may include a port)  
- `password`: Password for relay authentication  
- `ports`: Comma-separated list of ports for multiplexed transfer

### Address and Port Logic

- If `address` includes a port (e.g. `localhost:12345`):  
  - Handshake uses port `12345`  
  - Multiplexing uses the ports listed in `ports`

- If `address` does **not** include a port (e.g. `localhost`):  
  - Handshake uses the default port `9009`  
  - Multiplexing uses the ports listed in `ports`

## Hosting Your Own Relay
```sh
relay --ports 12345,12346 --pass geocine
```
- Match the ports and password in your `relays.json`  
- Ensure those ports are open and accessible on your network

## Security

- End-to-end encryption for all transfers  
- Only users with the correct codephrase **and** relay credentials can access the data

## License  
MIT