import React from 'react'
type message ={
    id: number,
    content: string
    sender: boolean
}
//TODO: Make the types global
export default function MessageBox(MessageProps : message) {
  return (
    <div className='MessageWrapper'>
        {MessageProps.sender 
        ?
        <div>
            <div className='bg-sky-600 rounded-sm p-4'>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
            </div>
            <p className='text-gray-400 text-sm p-2 pt-0'>5:40 PM</p>)
        </div>
        :
        <div>
            <div className='bg-gray-300 rounded-sm p-4 text-gray-900'>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
            </div>
            <p className='text-gray-400 text-sm p-2 pt-0'>5:40 PM</p>
        </div>
        }
    </div>
  )
}
