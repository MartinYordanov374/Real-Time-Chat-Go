"use client"
import React from 'react'
import MessageBox from '../MessageBox/MessageBox'
import SelectedUser from '../SelectedUser/SelectedUser'
import SelectedChat from '../SelectedChat/SelectedChat'
import Axios from 'axios'
import {useState, useEffect} from 'react'
import {Message} from '@/Types/Types'

export default function ChatBox({chat, chatMessages, setChatMessages, user}) {
  let online : boolean = true
  const [message, setMessage] = useState<string>('')
  useEffect(() => {
    setChatMessages(chat?.messages ?? [])
  }, [chat])
 // TODO: Test the send message function for the non-contacts.
 // 1. If a non-contact is sent a message, then they should receive a notification in their inbox
 // 2. A chatbox should appear for the receiver and sender user with the said chat
 // 3. A "request is pending" field should appear for the sender and they shall not be able to send any more messages until the receiver accepts their request
 // 4. Upon opening the chat request chatbox, the receiver can either accept or reject the request. If a request is rejected, the chat is deleted for both users.
 // 5. Test if there are any bugs when a user sends a message to a non-contact and then attempts to message a contact.
 useEffect(() => {
    console.log('user selected! ', user)
 }, [user])
 async function SendMessage(){
    let URI = user ? `http://localhost:8080/SendMessage/${user.id}`
 : `http://localhost:8080/SendMessage/${chat.receiverId}`
    let res = await Axios.post(URI, {"TextContent": message}, {withCredentials: true})

    .then((res) => {
        console.log(res)
    })
    .catch((err) => {
        console.log(err)
    })
  }
  return (
    chat == undefined && user == undefined ?
    <div className="flex h-screen flex-col w-full bg-gray-100">
        <div className='flex flex-1 justify-center items-center'>Start or select a chat and it will appear here</div>
    </div>
    :
    (user == undefined ?
    <SelectedChat chat={chat} chatMessages={chatMessages} setMessage={setMessage} sendMessage={sendMessage}/>
        :
    <SelectedUser user={user} setMessage={setMessage} sendMessage={sendMessage}/>
    )
  )
}
