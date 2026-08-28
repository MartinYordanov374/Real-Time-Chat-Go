import React from 'react'

export default function ChatList() {
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
    <div className="justify-left w-[40%] bg-white text-gray-900 border-r">
      {/* 1. User Header */}
      <div className='border-b pt-4'>
        <h2>Username</h2>
        <p>Active status</p>
      </div>  
      {/* 2. Search Bar */}
      <div className='border-b'>
        <input placeholder='Search users by username' className='p-4 w-full'/>
      </div>
      {/* 3. Contacts List */}
      {contacts.map((Contact) => (
        <div className='flex flex-col gap-2 p-2 hover:bg-gray-100 cursor-pointer'>
          <div className='flex'>
            <h2 className='font-semibold'>{Contact.Username}</h2>
            <p className='ml-auto text-gray-400 text-sm'>{Contact.LastMessageDate}</p>
          </div>
          {/* TODO: Cut out the message after a certain length */}
          <div className='flex-row'>
            <p className='text-sm'>{Contact.LastMessage}</p>
          </div>
        </div>
      ))}
      
    </div>
  )
}
