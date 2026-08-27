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
            <div className='bg-sky-600 rounded-lg p-4 mt-2 min-w-[70px]'>
                {MessageProps.content}
                <p className='text-xs text-right'>5:40 PM</p>
            </div>
        </div>
        :
        <div className='flex w-full justify-start max-w-[75%]'>
            <div className='bg-white rounded-lg p-4 text-gray-900 mt-2 min-w-[70px] shadow-lg'>
                {MessageProps.content}
                <p className='text-xs text-left'>5:40 PM</p>
            </div>
        </div>
        }
    </div>
  )
}
