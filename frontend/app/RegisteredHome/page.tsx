"use client"
import ChatBox from "../ChatBox/ChatBox";
import ChatList from "../ChatList/ChatList";
import Axios from 'axios'
import {useState, useEffect} from 'react'
export default function page() {
  // TODO: First fetch the chats and then load the page to ensure no issues with the page loading before the chats are retrieved occur
  const [chats, setChats] = useState()
  const [loading, setLoading] = useState(true)
  const [selectedChat, setSelectedChat] = useState(undefined)
  const [currentUsername, setCurrentUsername] = useState('')
  useEffect(() => {
    async function RetrieveUserChats(){
      let AllCurrentUserChats = await Axios.get('http://localhost:8080/GetAllChats', {withCredentials: true})
      .then((res) =>{
        setChats(res.data?.chats)
        RetrieveCurrentUsername()
      })
      .catch((err) => {
        console.log(err)
      })
    }

    async function RetrieveCurrentUsername(){
        let CurrentUsername = await Axios.get('http://localhost:8080/GetCurrentUserData', {withCredentials: true})
        .then((res) =>{
          setCurrentUsername(res.data.User)
          setLoading(false)
        })
        .catch((err) => {
          console.log(err)
        })
    }
    RetrieveUserChats()
  }, [])

  return (
    loading == false ?
    <div className="flex">
        {/* TODO: Send the chats list to the chat list component and populate the contacts page with the chats */}
        <ChatList chats = {chats} onSelect={setSelectedChat} currentUsername={currentUsername}/>
        {/* TODO: Selecting a contact from the chat list 
        will open the chat box with the corresponding conversation */}
        <ChatBox chat = {selectedChat}/>
    </div>
    :
    <div>
      <p>Loading</p>
    </div>
  )
}
