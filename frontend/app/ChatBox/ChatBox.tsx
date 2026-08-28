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
    <div className="flex h-screen flex-col w-full bg-gray-100">
        <div className='sticky top-0 bg-white flex p-4'>
            <div className='flex size-12 rounded-full bg-blue-500 p-4 justify-center items-center'>PFP</div>
            <div className='flex-col'>
                <h2 className='text-gray-900 pl-4'>Username</h2>
                {online 
                    ?
                    <p className='text-emerald-500 font-bold pl-4 pb-4'> Online </p>
                    :
                    <p className='text-gray-400 pl-4 pb-4'> Offline </p>
                }
            </div>
        </div>
        <div className='flex flex-col-reverse max-w h-screen p-10 max-h-screen overflow-scroll'>
            {Messages.map((message) => (
                <MessageBox 
                key ={message.id}
                content={message.content}
                sender={message.sender}/>
            ))}
        </div>
        <div className='sticky bottom-0 bg-white p-4 gap-2'>
            <div className='flex items-end bg-gray-200 rounded-lg text-gray-900 p-4 flex-1'>
                <textarea 
                className='
                w-full
                resize-none
                focus:outline-none
                p-2'
                placeholder='Send a message'/>
            
                <button className='
                text-blue-500 
                hover:cursor-pointer
                hover:text-blue-600
                flex size-10'>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="size-6">
                        <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
                    </svg>
                </button>
            </div>
        </div>
    </div>
  )
}
