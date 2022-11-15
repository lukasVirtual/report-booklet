import express from "express"
import http from "http"
import { Server } from "socket.io"
import fetch from "node-fetch"

// import { createAdapter } from "@socket.io/redis-adapter";
import redis from "redis"
const app = express()
const server = http.createServer(app)
const client = redis.createClient({ url: "redis://localhost:6379"})
const io = new Server(server, { cors: {
    origin: "http://127.0.0.1:3000",
    methods: ["GET", "POST"]
}})

await client.connect()
// io.adapter(createAdapter(client, client.duplicate()))

app.get('/', (req, res) => {
    res.send("<h1>Working</h1>")
})

io.on('connection', (socket) => {
    socket.on('login', async (user) => {
        console.log(`User: ${user.user_name} with id ${user.user_id} socket id ${socket.id} connected`)
        await client.set(user.user_name, socket.id, function(err) { if (err) throw err })
    })

    socket.on('logout', async (user) => {
        const exists = await client.exists(user.user_name)
        if (exists === 1) {
            console.log(`User: ${user.user_name} with id ${user.user_id} disconnected`)
            const res = await client.del(user.user_name)
            console.log(res)
        } else console.log("no one connected")

    })
    
    socket.on('submit', async (data, sendFrom, month) => {
        // console.log("data", data, month)
        console.log(sendFrom)

        // const reciever = await client.get(sendFrom)
        // if (reciever) {
        //     console.log("r: ",reciever)


            /* TODO
             * sending to specific user is currently not
             * working need to fix that
            */   
           
            const values = JSON.stringify({
                name: sendFrom,
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
    
            io.emit("test", sendFrom)
            console.debug("message sent")
        // } else {
        //     console.error("error: ", reciever)
        // }
    })
})

server.listen(7000, () => {
    console.log("listening...")
})