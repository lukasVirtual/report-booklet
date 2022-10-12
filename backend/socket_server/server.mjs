import express from "express"
import http from "http"
import { Server } from "socket.io"
const app = express()
const server = http.createServer(app)
const io = new Server(server, { cors: {
    origin: "http://127.0.0.1:3000",
    methods: ["GET", "POST"]
}})

app.get('/', (req, res) => {
    res.send("<h1>Working</h1>")
})

const socketUsers = new Map()
io.on('connection', (socket) => {
    socket.on('login', (user) => {
        console.log(`User: ${user.user_name} with id ${user.user_id} socket id ${socket.id} connected`)
        socketUsers.set(user.user_name, socket.id)
    })

    socket.on('logout', (user) => {
        console.log(`User: ${user.user_name} with id ${user.user_id} disconnected`)
        if (socketUsers.get(user.user_name) !== undefined) socketUsers.delete(user.user_name)
    })
    
    socket.on('submit', (data, sendTo) => {
        console.log("data", data)
        //Need to be done
        const reciever = socketUsers.get(sendTo)
        if (reciever) {
            console.log(reciever)
            /* TODO
             * sending to specific user is currently not
             * working need to fix that
            */   
            //socket.broadcast.to(reciever).emit("test", data) 
            io.emit("test", JSON.parse(data))
            console.debug("message sent")
        } else {
            console.error("error: ", reciever)
        }
    })
})

server.listen(7000, () => {
    console.log("listening...")
})