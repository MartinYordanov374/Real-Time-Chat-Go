import Axios from 'axios'
export default async function Register(username, email, password)
{
        let res = await Axios.post('http://localhost:8080/register', {username:username, email:email, password:password}, {withCredentials: true})
        console.log(res)
        return res
}
