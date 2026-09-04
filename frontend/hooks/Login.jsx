import Axios from 'axios'
export default async function Login(username, password)
{
        let res = await Axios.post('http://localhost:8080/login', {username:username, password:password}, {withCredentials: true})
        return res
}
