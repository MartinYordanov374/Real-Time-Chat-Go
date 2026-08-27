import ChatBox from "../ChatBox/ChatBox";
import ChatList from "../ChatList/ChatList";

export default function page() {
  return (
    <div className="flex">
        <ChatList/>
        {/* TODO: Selecting a contact from the chat list 
        will open the chat box with the corresponding conversation */}
        <ChatBox/>
    </div>
  )
}
