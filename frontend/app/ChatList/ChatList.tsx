import React from 'react'

export default function ChatList({chats, onSelect, currentUsername}) {
  return (
    <div className="justify-left w-[50%] bg-white text-gray-900">
      {/* 1. User Header */}
      <div className='flex shadow-xs p-4'>
        <div className='flex flex-col size-12 rounded-full bg-blue-500 p-2 justify-center items-center'>PFP</div>
        <div className='pl-4 flex flex-col'>
          <span className='font-bold'>{currentUsername}</span>
          <span className='text-sm text-gray-400 pt-0'>Online</span>
        </div>
      </div>  
      {/* TODO: Implement search existing AND new contacts */}
      {/* 2. Search Bar */}
      <div className='p-4'>
        <input 
        className="p-4 w-full focus:outline-none bg-gray-200 rounded-xl text-sm"
        placeholder='Find contacts by username'/>
      </div>
      {/* 3. Contacts List */}
      {chats.map((Chat) => (
        <div className='flex flex-col gap-2 p-4 hover:bg-gray-100 cursor-pointer' onClick={() => onSelect(Chat)} key={Chat.id}>
          <div className='flex'>
            <div className=' flex size-12 rounded-full bg-blue-500 p-2 justify-center items-center'>
              PFP
            </div>
            <h2 className='font-semibold pl-2'>{Chat.DisplayedUsername}</h2>
            {/* TODO: for mobile resolutions, move the timestamp under the message*/}
            <p className='ml-auto text-gray-400 text-sm'>{Chat.messages[0].timeStamp}</p>
          </div>
          {/* TODO: Cut out the message after a certain length */}
          <div className='flex-row'>
            <p className='text-sm'>{Chat.messages[0].textContent}</p>
          </div>
        </div>
      ))}
      
    </div>
  )
}
