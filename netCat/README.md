# TCPChat: A NetCat-Inspired Server-Client Chat Application

## Overview

TCPChat is a server-client chat application inspired by the functionality of the NetCat (`nc`) system command. It supports multiple clients connecting to a server and enables real-time group communication over a TCP connection. The project replicates key features of NetCat while providing a user-friendly group chat experience with additional enhancements.

## Objectives

The primary objective of this project is to recreate the functionality of NetCat with a focus on group chat. The application operates in a server-client architecture, where:

- **Server Mode**: The application listens on a specified port for incoming client connections.
- **Client Mode**: Clients connect to the server on a specified port and participate in a group chat.

### Key Features

1. **Server-Client Architecture**:
   - Establishes TCP connections between a server and multiple clients.
   - Supports up to **10 simultaneous client connections**.

2. **Client Identification**:
   - Requires clients to provide a unique username upon connection.

3. **Group Chat Functionality**:
   - Messages include a timestamp, the sender’s username, and the message content, e.g.,  
     `[2020-01-20 15:48:41][client.name]:[client.message]`.
   - Empty messages are ignored and not broadcasted.

4. **History Management**:
   - New clients receive all previous chat messages upon joining.

5. **Connection Notifications**:
   - Clients are notified when a new user joins or leaves the chat.

6. **Client Name Change**
    - Clients can change their names to new names by writting `--cn` to the chat, then writing new name from the prompt given.

7. **Resilient Design**:
   - Clients can leave the chat without disrupting the connection for others.
   - The server continues to handle connections and broadcasts even with client exits.

8. **Default and Custom Ports**:
   - Defaults to port `8989` if no port is specified.
   - Responds with a usage message if the arguments are incorrect:  
     `[USAGE]: ./TCPChat $port`.

### Usage Scenarios

- **Collaborative group discussions**.
- **Real-time communication across multiple clients**.
- **Learning about TCP sockets, Go concurrency, and network programming**.

## Instructions

### Prerequisites

- Ensure you have **Go** installed on your system.

### Usage

#### Start the Server
```bash
# Default port
$ go run .  
Listening on the port :8989

# Custom port
$ go run . 2525  
Listening on the port :2525
```

#### Incorrect Usage
```bash
$ go run . 2525 localhost
[USAGE]: ./TCPChat $port
```

### Features at a Glance

- **Server**: Listens for connections and handles client communication.
- **Client**: Connects to the server and participates in the chat.

### Requirements met

1. Written entirely in **Go**.
2. Leverages **goroutines** for concurrency.
3. Handles a maximum of **10 simultaneous client connections**.
4. Implements proper error handling on both server and client sides.

### Allowed Go Packages

- `io`, `log`, `os`, `fmt`
- `net`, `sync`, `time`, `bufio`
- `errors`, `strings`, `reflect`

## Learning Outcomes

This project offers practical experience with:

- **TCP and UDP**:
  - Understanding connections, sockets, and protocols.
- **Go Concurrency**:
  - Using goroutines and channels effectively.
  - Synchronizing data with mutexes.
- **Networking Concepts**:
  - Handling IP addresses and ports.
  - Managing client-server communication.

## Example Usage

### Server Output
```bash
$ go run .
Listening on the port :8989
```

### Client Output

client 1

```bash
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:Alice
[2024-12-04 11:59:00][Bob]: Hi there!
[2024-12-04 12:01:00][Alice]: Hello, everyone!
Shadrack joined the chat...
[2024-12-04 12:01:00][Shadrack]: Hello, Shadrack here...
```

Client 2


```bash
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:Shadrack
[2024-12-04 11:59:00][Bob]: Hi there!
[2024-12-04 12:01:00][Alice]: Hello, everyone!
[2024-12-04 12:01:00][Shadrack]: Hello, Shadrack here...
```

## Future Enhancements

- Support for **encrypted communication**.
- **File sharing** between clients.
- Integration with **UDP** for alternate communication modes.

### Collaborators


<table>
<tr>
<td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/jeseekuya >
            <img src=https://learn.zone01kisumu.ke/git/avatars/b990e7e45f5bae8559b9dc30dd411027?size=84 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Jesee_Kuya/>
            <br />
            <sub style="font-size:14px"><b><i>Jesee Kuya</i><span></b></sub>
        </a>
  </td> 
  <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/barrack-kope-064a43244/ >
            <img src=https://learn.zone01kisumu.ke/git/avatars/69ae4e7472c685f60beaf04edb53b624?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Barrack_Otieno/>
            <br />
            <sub style="font-size:14px"><b><i>Barrack Otieno</i></b></sub>
        </a>
  </td> 
   
</tr>

</table>


### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## Conclusion

TCPChat is a robust, NetCat-inspired application built with Go, providing a seamless group chat experience. It is a valuable tool for learning about network programming, Go concurrency, and TCP/UDP connections. Dive into the source code to explore its implementation and expand your understanding of these fundamental concepts.