<h1 align="center">WoL VE</h1>

<p align="center">
🖥️ Start VMs and LXCs · HTTP & WebSocket API · Multi-Platform
</p>

<div align="center">
  <a href="https://github.com/codeshelldev/wol-ve/releases">
    <img 
        src="https://img.shields.io/github/v/release/codeshelldev/wol-ve?sort=semver&logo=github&label=Release" 
        alt="GitHub release"
    >
  </a>
  <a href="https://github.com/codeshelldev/wol-ve/stargazers">
    <img 
        src="https://img.shields.io/github/stars/codeshelldev/wol-ve?style=flat&logo=github&label=Stars" 
        alt="GitHub stars"
    >
  </a>
  <a href="./LICENSE">
    <img 
        src="https://img.shields.io/badge/License-MIT-green.svg"
        alt="License: MIT"
    >
  </a>
</div>

---

**WoL VE** is a lightweight Go program exposing an **HTTP and WebSocket API** for starting virtual environments such as **virtual machines** and **LXC containers**.

Use standalone or together with  
[**WoL-Redirect**](https://github.com/codeshelldev/wol-redirect) for a web-based UI.

## Installation

Download the latest binary from the Releases page, mark it executable, and run:

```bash
chmod +x wol-ve
./wol-ve
```

## Usage

Start a VM or LXC instance by sending a POST request to `/wake`.  
Example: start an instance with ID `100`:

```
curl -X POST "http://wol-ve:9000/wake" \
     -H "Content-Type: application/json" \
     -d '{
           "id": "100",
           "ip": "192.168.1.1",
           "startupTime": 5
         }'
```

> [!NOTE]
> - `startupTime` is optional.
> - If `startupTime` is supplied, it acts as a **maximum wait time** while still allowing the ping-based readiness check.
> - If `startupTime` is omitted, readiness depends entirely on the ping logic using [`PING_INTERVAL`](#ping_interval) and [`PING_RETRIES`](#ping_retries).
> - `ip` is optional and is only needed if ping-based readiness should be used.

## WebSocket Updates

The `/wake` endpoint returns a `client_id`.  
Use it to open a WebSocket connection:

```
ws://wol-ve:9000/ws
```

The WebSocket sends structured updates during the startup sequence:

- `success`: `true` when the process completes
- `error`: `true` if startup fails
- `message`: descriptive status or error details

## Configuration

### `PING_INTERVAL`

Interval in seconds for pinging when `startupTime` is not provided.

### `PING_RETRIES`

Number of retries for pinging when `startupTime` is not provided before timing out.

## Contributing

Have suggestions or improvements? Feel free to open an issue or submit a Pull Request.

## License

This project is licensed under the [MIT License](./LICENSE).
