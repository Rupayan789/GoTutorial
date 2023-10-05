const express = require('express')
const { json, urlencoded } = express
// import express, { json, urlencoded } from "express"

const app = express()

app.use(json())
app.use(urlencoded({extended:true}))

const users = [
    {
        id:1,
        name: 'John',
        age: 34
    },
    {
        id:2,
        name: 'Dane',
        age: 40
    },
    {
        id:3,
        name: 'Bob',
        age: 45
    }
]

app.get("/",(req,res)=>{
    return res.status(200).json({"message":"Welcome to my server"})
})

app.get('/:id',(req,res)=>{
    const id = req.params.id
    const user = users.find(item => item.id == id)
    if(!user){
        return res.status(404).json({"message":`No user found with id ${id}`})
    }
    return res.status(200).json({"message":"User found","data":user})
})


app.post("/",(req,res)=>{
    const user = req.body
    const doesExist = users.find(item=>item.name==user.name)
    if(doesExist){
        return res.status(404).json({"message":"User already exists"})
    }
    console.log(user)
    console.log(users)
    users.push(user)
    return res.status(200).json({"message":"User successfully added"})
})

app.listen(3000,()=>{
    console.log("Server up and running")
})