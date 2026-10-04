"use client"
import ChatBox from "../ChatBox/ChatBox";
import ChatList from "../ChatList/ChatList";
import Axios from 'axios'
import {useState, useEffect} from 'react'
export default function page() {
  const [chats, setChats] = useState([])
  const [loading, setLoading] = useState(true)
  const [selectedChat, setSelectedChat] = useState(undefined)
  const [currentUsername, setCurrentUsername] = useState('')
  const [selectedChatMessages, setSelectedChatMessages] = useState([])
  const [selectedUser, setSelectedUser] = useState(undefined)

  useEffect(() => {
    async function RetrieveUserChats(){
      let AllCurrentUserChats = await Axios.get('http://localhost:8080/GetAllChats', {withCredentials: true})
      .then((res) =>{
          if (res.data != null && res.data.chats != null){
            setChats(res.data.chats)
        }
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
    RetrieveCurrentUsername()
  }, [])

  useEffect(() => {
    if (!selectedChat?.id){ 
      return 
    }
    const Socket = new WebSocket("ws://localhost:8080/ws")
    Socket.onopen = () => {
        console.log('Socket connection established')
    }
    Socket.onmessage = (event) => {
      console.log(event)
      let parsedData = JSON.parse(event.data)
      setSelectedChatMessages((prevState) => [
        parsedData,
        ...prevState
      ])

    }

    Socket.onclose = () => {
        console.log('socket connection closed')
    }
    
    return () => {
      Socket.close()
    }
  }, [selectedChat?.id])

  useEffect(() => {
    const Socket = new WebSocket("ws://localhost:8080/ws")
    Socket.onopen = () => {
        console.log('Socket connection established')
    }
    Socket.onmessage = (event) => {
      let parsedData = JSON.parse(event.data)
      console.log('Setting Chat data')
      // TODO: Fix the initial double rendering of items when a chat request shows
      // I think it may have something to do with more than one web socket connection being open
      // at a time for the same user

      setChats((prevState) => [
          parsedData,
          ...prevState
      ])
    }

  }, [])

  return (
    loading == false ?
    <div className="flex">
        <ChatList chats = {chats} onSelectChat={setSelectedChat} currentUsername={currentUsername} onSelectUser={setSelectedUser}/>
        <ChatBox chat = {selectedChat} chatMessages={selectedChatMessages} setChatMessages={setSelectedChatMessages} user={selectedUser}/>
    </div>
    :
    <div>
      <p>Loading</p>
    </div>
  )
}
