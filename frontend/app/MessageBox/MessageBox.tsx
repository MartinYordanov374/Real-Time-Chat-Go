import React from 'react'
import {Message} from '@/Types/Types'

//TODO: Figure out a more aesthetic way of showing the timestamp
export default function MessageBox(MessageProps : Message) {
    return (
    <div className='MessageWrapper' key={MessageProps.id}>
        {MessageProps.IsCurrentUserSender 
        ?
        <div className='flex w-full justify-end'>
            <div className='bg-blue-500 rounded-lg p-4 mt-2 min-w-[70px]'>
                {MessageProps.textContent}
                <p className='text-xs text-right'>5:40 PM</p>
            </div>
        </div>
        :
        <div className='flex w-full justify-start'>
            <div className='bg-white rounded-lg p-4 text-gray-900 mt-2 min-w-[70px] shadow-lg'>
                {MessageProps.textContent}
                <p className='text-xs text-left'>5:40 PM</p>
            </div>
        </div>
        }
    </div>
  )
}
