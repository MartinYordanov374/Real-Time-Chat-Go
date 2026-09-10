"use client"
import ChatBox from "../ChatBox/ChatBox";
import ChatList from "../ChatList/ChatList";
import Axios from 'axios'
import {useState, useEffect} from 'react'
export default function page() {
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
        <ChatList chats = {chats} onSelect={setSelectedChat} currentUsername={currentUsername}/>
        <ChatBox chat = {selectedChat}/>
    </div>
    :
    <div>
      <p>Loading</p>
    </div>
  )
}
