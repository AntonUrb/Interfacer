# Milrem Internship Tooling Assignment

This microservice application provides information about network interfaces through a simple HTTP API. It consists of two components: an HTTP server and an HTTP client, both running in separate Docker containers.
**The application is written in Go and deployed on Docker.**

## Setup

1. **Prerequisites**: Make sure you have Docker installed on your machine. If not, you can download and install it from [Docker's official website](https://www.docker.com/get-started).

---

2. **Clone the Repository**: Clone this repository to your local machine using the following command:
   
```
git clone github.com/AntonUrb/Milrem-Internship-Tooling.git
```

---

3. **Navigate to the Project Root Directory**: Move into the cloned repository's directory:

```
cd Milrem-Internship-Tooling
```

---

4. **Build and Run the Containers**: Use Docker Compose to build and run the server and client containers:

```
sudo docker compose -f docker-compose.yml build && sudo docker compose -f docker-compose.yml up
```

This command will build the Docker images for both the server and client components, create containers, and start the services.

- *Troubleshooting*: if you get this or any other error
```
Error response from daemon: network c1a87c76f774fc1ed617c48f4e0884d40b804461de60802386aa4c7df929b58e not found
```

**run** `sudo docker compose down` and try the first command again.

---

5. **Monitoring Network Interfaces:**: Once the containers are up and running, you should see the http-client start polling http-server for interfaces and printing them out in the terminal.

---

## API Overview

### List Network Interfaces

- **Endpoint**: `/network`
- **Method**: `GET`
- **Query Parameters** (Optional):

    `?interface={interface_name}`: Specifies the name of a network interface. If provided, the API will return details only for the specified interface. Otherwise it returns all interfaces.

- **Response Structure**:

The API returns a JSON response with details about network interfaces.

**200 OK**: Returns details about the requested network interface(s) in JSON format.

    Fields include:
        Name: The name of the network interface.
        IP Addresses: The list of IP addresses associated with the interface.
        MAC Address: The MAC (Media Access Control) address of the interface.
        MTU: The Maximum Transmission Unit of the interface.
        Speed: The speed of the interface.
        Duplex: The duplex mode of the interface (e.g., full or half).
        Admin Status: The administrative status of the interface (enabled, disabled, or unknown).
        Operational Status: The operational status of the interface (UP, DOWN, or unknown).

- **Response Example**:
```
[
  {
    "name": "eth0",
    "ip_addresses": [
        "192.168.1.10",
        "10.0.0.1"
    ],
    "mac_address": "00:11:22:33:44:55",
    "mtu": 1500,
    "speed": "1 Gbps",
    "duplex": "Full",
    "admin_status": "enabled",
    "operational_status": "UP"
  },
  {
    "name": "wlan0",
    "ip_addresses": [
        "192.168.2.20"
    ],
    "mac_address": "a1:b2:c3:d4:e5:f6",
    "mtu": 1200,
    "speed": "100 Mbps",
    "duplex": "Half",
    "admin_status": "disabled",
    "operational_status": "DOWN"
  }
]
```

---

### Error Handling

**404 Not Found** is returned with an error message, if the specified interface doesn't exist.

`{
  "error": "interface not found"
}`

**400 Bad Request** is returned with an error message, if an invalid query parameter is provided.

`{
  "error": "only ?interface={interface_name} input format is allowed"
}`

**500 Internal Server Error** is returned with an error message, if an internal server error occurs.

`{
  "error": "internal server error"
}`

---

## Important Notes

- The server runs on [port :8080] http://localhost:8080/network. (With this link you can also access the server in your browser)

- The HTTP-client periodically calls the server's endpoint to fetch network interface details.

- The code itself is documented and readable whenever you're curious about how something works.

## Additional Information

**Interval Configuration**

You can modify the interval at which the HTTP client calls the server's endpoint by changing the `INTERVAL` environment value of http-client (e.g., 3s for 3 seconds) in the `docker-compose.yml` file that's located in the root directory. The default value is 5 seconds.

**Interface Parameter**

You can search for a specific interface which the API should display details about, by changing the `INTERFACE` environment value in the http-client (e.g. eth0 or wlan0) in the `docker-compose.yml` file that's located in the root directory. By default the value is empty, which means that all interfaces get displayed.

**System with Go installed**

If you have Go installed on your system, you can run both http-server and http-client directly on your system.

  1. Navigate to the task's root directory.

  2. Type `cd client/ && make` into the first terminal and `cd server/ && make` into the second.

This will use the Makefiles in the respective directories and run the http-server and http-client with preset environment variables. The Makefiles are also configured to automatically run test files on the API and client before launching.

---

### API Test Suite

- **Build and Run the HTTP Test-Server & Test-Client Containers using this command**:

```
sudo docker compose -f docker-compose-test.yml build && sudo docker compose -f docker-compose-test.yml up

```

- **In case of errors**, run: `sudo docker compose down` and try the first command again.

Description:

`sudo docker compose -f docker-compose-test.yml build`: This part of the command builds the Docker containers for the HTTP test-server & test-client. It compiles the source code and creates the container images.

`sudo docker compose -f docker-compose-test.yml up`: This part of the command runs the built Docker containers for the HTTP test-server & test-client. It starts the containers and executes the defined commands within the containers, which are to use the integrated test files to test the functionality of the http-server and http-client and it's edge-cases.

---

### Author

<a href="www.linkedin.com/in/anton-urban-4544522b5" alt="LinkedIn"><strong>Anton Urban</strong></a>