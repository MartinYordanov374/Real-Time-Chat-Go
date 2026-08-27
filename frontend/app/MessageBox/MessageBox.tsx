import React from 'react'
type message ={
    id: number,
    content: string
    sender: boolean
}
//TODO: Make the types global
//TODO: Figure out a more aesthetic way of showing the timestamp
export default function MessageBox(MessageProps : message) {
  return (
    <div className='MessageWrapper'>
        {MessageProps.sender 
        ?
        <div className='flex w-full justify-end max-w-[75%]'>
            <div className='bg-sky-600 rounded-lg p-4'>
                {MessageProps.content}
            </div>
        </div>
        :
        <div className='flex w-full justify-start max-w-[75%]'>
            <div className='bg-white rounded-lg p-4 text-gray-900'>
                {MessageProps.content}
            </div>
        </div>
        }
    </div>
  )
}
