"use client"
import {useState} from 'react'
import Register from '../../Helpers/Register'

export default function RegisterPage() {

  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  function HandleRegister(){
    // TODO: Check if the confirm password field value matches the password field value
    PerformRegister()
  }

  async function PerformRegister(){
    try{
      let UserLoginRes = await Register(username, email, password)
    }
    catch(err){
      console.log(err.response?.data.message)
    }
  }
  
  return (
    <div className="flex items-center min-h-screen justify-center bg-gray-50">
      <div className="w-full max-w-md rounded-lg p-8 shadow-xl">
        <div className="flex justify-center mb-2">
          <div className="bg-blue-600 rounded-2xl p-4">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2.5} stroke="white" className="size-6">
              <path strokeLinecap="round" strokeLinejoin="round" d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785A5.969 5.969 0 0 0 6 21c1.282 0 2.47-.402 3.445-1.087.81.22 1.668.337 2.555.337Z" />
            </svg>
          </div>
        </div>
        <h2 className="text-center text-gray-900 text-lg font-bold ">Golang RTC</h2>
        <p className="text-center text-gray-400 text-sm font-semibold mb-4">Join the Golang RTC platform.</p>
          <label className="text-gray-800 font-semibold">Username</label>
          <input 
            type="text"
            placeholder="Enter username" 
            className="
            placeholder-gray-500 
            text-gray-500
            border 
            border-gray-200
            bg-gray-100
            rounded-lg
            w-full
            text-left
            p-2
            mb-4
            focus:outline-none
            focus:ring-2 
            focus:ring-blue-500"
            onChange={(e) => setUsername(e.target.value)}
          />
          <label className="text-gray-800 font-semibold">Email</label>
          <input 
            type="text"
            placeholder="samplemail@msamplemail.com" 
            className="
            placeholder-gray-500 
            text-gray-500
            border 
            border-gray-200
            bg-gray-100
            rounded-lg
            w-full
            text-left
            p-2
            mb-4
            focus:outline-none
            focus:ring-2 
            focus:ring-blue-500"
            onChange={(e) => setEmail(e.target.value)}
          />
          <label className="text-gray-800 font-semibold">Password</label>
          <input 
            type="password"
            placeholder="Enter password" 
            className="
            placeholder-gray-500 
            text-gray-500
            border 
            border-gray-200
            bg-gray-100
            rounded-lg
            w-full
            p-2
            text-left
            focus:outline-none
            focus:ring-2 
            focus:ring-blue-500
            mb-4"
            onChange={(e) => setPassword(e.target.value)}
          />
          <label className="text-gray-800 font-semibold">Confirm Password</label>
          <input 
            type="password"
            placeholder="Confirm your password" 
            className="
            placeholder-gray-500 
            text-gray-500
            border 
            border-gray-200
            bg-gray-100
            rounded-lg
            w-full
            p-2
            text-left
            focus:outline-none
            focus:ring-2 
            focus:ring-blue-500
            mb-4"
          />
          <button
            type="submit"
            className="
            w-full 
            bg-blue-600
            rounded-lg
            p-2
            hover:bg-blue-700 
            cursor-pointer
            text-white
            font-semibold
            "
            onClick={() => HandleRegister()}
          >
            Sign up
          </button>

          <p className="text-gray-900 text-center mt-4"> Already have an account? 
            <a className="
            text-blue-500
            font-bold
            hover:text-blue-600 cursor:pointer" 
            href='/login'> Log in here.</a> 
          </p>
      </div>
    </div>
  );
}
