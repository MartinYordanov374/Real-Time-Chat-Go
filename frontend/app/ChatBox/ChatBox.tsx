import React from 'react'
import MessageBox from '../MessageBox/MessageBox'

type message ={
    id: number,
    content: string
    sender: boolean,
}
export default function ChatBox() {
    let online : boolean = true
    let Messages : message[] = [{
        id: 1,
        content: "test 1",
        sender: true,
    },
    {
        id: 2,
        content: "test 2",
        sender: false,
    },
    {
        id: 3,
        content: "test 3",
        sender: false,
    },
    {
        id: 4,
        content: "test 4",
        sender: true,
    },
    {
        id: 5,
        content: "test 5",
        sender: false
    },
    {
        id: 6,
        content: "test 6",
        sender: true
    },
    {
        id: 7,
        content: "test 7",
        sender: true
    },
    {
        id: 8,
        content: "test 8",
        sender: true
    },
    {
        id: 9,
        content: "test 9",
        sender: true
    }]
  return (
    <div className="w-full bg-gray-100 max-h-screen overflow-scroll">
        <div className='border border-gray-900sticky top-0 bg-white'>
            <h2 className='text-gray-900 pl-8 pt-8'>Username</h2>
            {online 
                ?
                <p className='text-emerald-500 font-bold pl-8 pb-4 border-bottom border-gray-900'> Online </p>
                :
                <p className='text-gray-400 pl-8 pb-4 border-bottom border-gray-900'> Offline </p>
            }
        </div>
        <div className='flex flex-col-reverse max-w h-screen p-10'>
            {Messages.map((message) => (
                <MessageBox 
                key ={message.id}
                content={message.content}
                sender={message.sender}/>
            ))}
        </div>
    </div>
  )
}
