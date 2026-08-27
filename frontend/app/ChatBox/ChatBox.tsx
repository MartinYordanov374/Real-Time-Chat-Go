import React from 'react'
import MessageBox from '../MessageBox/MessageBox'

type message ={
    id: number,
    content: string
    sender: boolean
}
export default function ChatBox() {
    let Messages : message[] = [{
        id: 1,
        content: "test",
        sender: true
    },
    {
        id: 2,
        content: "Another test",
        sender: false,
    },
    {
        id: 3,
        content: "Another test",
        sender: false,
    },
    {
        id: 4,
        content: "Another test",
        sender: true,
    },
    {
        id: 5,
        content: "Another test",
        sender: false
    }]
  return (
    <div className="max-w-md w-full bg-white max-h-screen overflow-scroll p-10">
        {Messages.map((message) => (
            <MessageBox 
            id ={message.id}
            content={message.content}
            sender={message.sender}/>
        ))}
    </div>
  )
}
