import React from 'react'

export default function ChatList() {
  // TODO: Make it so that the chat box shows the ID corresponding to the user on the other end of the chat, 
  // i.e., if User A is retrieving the chats, they shall see user B in the contacts list
  const contacts = [{
    Username: "John Doe",
    LastMessage: "This is the last message from john doe",
    LastMessageDate: "Friday"
  },
  {
    Username: "The real John Doe",
    LastMessage: "This is the last message from THE REAL john doe and he is not playing games!",
    LastMessageDate: "12:15 PM"
  }]
  return (
    <div className="justify-left w-[50%] bg-white text-gray-900">
      {/* TODO: Populate the user header with data about the current user */}
      {/* 1. User Header */}
      <div className='flex shadow-xs p-4'>
        <div className='flex flex-col size-12 rounded-full bg-blue-500 p-2 justify-center items-center'>PFP</div>
        <div className='pl-4 flex flex-col'>
          <span className='font-bold'>Username</span>
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
      {/* TODO: List all of the chats that the user is a member of. Populate the data accordingly. */}
      {/* 3. Contacts List */}
      {contacts.map((Contact) => (
        <div className='flex flex-col gap-2 p-4 hover:bg-gray-100 cursor-pointer'>
          <div className='flex'>
            <div className=' flex size-12 rounded-full bg-blue-500 p-2 justify-center items-center'>
              PFP
            </div>
            <h2 className='font-semibold pl-2'>{Contact.Username}</h2>
            {/* TODO: for mobile resolutions, move the timestamp under the message*/}
            <p className='ml-auto text-gray-400 text-sm'>{Contact.LastMessageDate}</p>
          </div>
          {/* TODO: Cut out the message after a certain length */}
          <div className='flex-row'>
            <p className='text-sm'>{Contact.LastMessage.split('').length > 83 ? Contact.LastMessage.slice(0, 83)+"..." : Contact.LastMessage}</p>
          </div>
        </div>
      ))}
      
    </div>
  )
}
