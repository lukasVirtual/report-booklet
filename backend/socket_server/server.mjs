import express from "express"
import http from "http"
import { Server } from "socket.io"
import fetch from "node-fetch"

// import { createAdapter } from "@socket.io/redis-adapter";
import redis from "redis"
const app = express()
const server = http.createServer(app)
const client = redis.createClient({ url: "redis://localhost:6379"})
await client.connect()
const io = new Server(server, { cors: {
    origin: "http://127.0.0.1:3000",
    methods: ["GET", "POST"]
}})

// io.adapter(createAdapter(client, client.duplicate()))

app.get('/', (req, res) => {
    res.send("<h1>Working</h1>")
})

// const socketUsers = new Map()
io.on('connection', (socket) => {
    socket.on('login', async (user) => {
        console.log(`User: ${user.user_name} with id ${user.user_id} socket id ${socket.id} connected`)
        // socketUsers.set(user.user_name, socket.id)
        await client.set(user.user_name, socket.id, function(err) { if (err) throw err })
    })

    socket.on('logout', async (user) => {
        const exists = await client.exists(user.user_name)
        if (exists === 1) {
            console.log(`User: ${user.user_name} with id ${user.user_id} disconnected`)
            await client.del(user.user_name)
        } else console.log("no one connected")

        // if (socketUsers.get(user.user_name)) socketUsers.delete(user.user_name)
    })
    
    socket.on('submit', async (data, sendTo, month) => {
        console.log("data", data, month)
        //Need to be done
        // const reciever = socketUsers.get(sendTo)
        const reciever = await client.get(sendTo)
        if (reciever) {
            console.log(reciever)
            /* TODO
             * sending to specific user is currently not
             * working need to fix that
            */   
           
            const values = JSON.stringify({
                name: sendTo,
                date: month,
                data: data,
            });
    
            const res = await fetch("http://127.0.0.1:5000/api/saveSubmittedData", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
                },
                body: values,
            });

            console.log(res.status)
    
            io.emit("test", sendTo)
            console.debug("message sent")
        } else {
            console.error("error: ", reciever)
        }
    })
})

server.listen(7000, () => {
    console.log("listening...")
})