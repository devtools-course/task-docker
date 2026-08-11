# Ping pong

You are provided with two servers: [back.go](./back.go) and [front.go](./front.go).

Add a parameterised Dockerfile that can be used to build both servers. 

> Non-parameterised solutions won't be accepted!

Create two scripts:
* `mkenv.sh` - Creates the environment
* `rmenv.sh` - Destroys the environment

## Environment

* Shared network `korea`
* Volume `hanja`
* Two containers: `south` & `north`

### South

The volume must be mounted at `/data`

### North

The container's `7676` port must be forwarded to the host's `9999` port


---
* How can the parameterisation be achieved with Dockerfiles?
* What are the alternatives to the parameterisation?
* How can a container be accessed through a network?
* When are volumes needed?
