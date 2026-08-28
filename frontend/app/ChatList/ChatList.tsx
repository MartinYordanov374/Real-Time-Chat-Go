import React from 'react'

export default function ChatList() {
  return (
    <div className="justify-left w-[40%] bg-white text-gray-900 border-r">
      {/* 1. User Header */}
      <div className='border-b '>
        <h2>Username</h2>
        <p>Active status</p>
      </div>
      {/* 2. Search Bar */}
      <div>
        <input placeholder='Search users by username'/>
      </div>
      {/* 3. Contacts List */}
      <div>
        <h2>Username</h2>
        <p>Last message</p>
        <p>Sent date</p>
      </div>
    </div>
  )
}
