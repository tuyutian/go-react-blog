export default class LoginService {
    static Login(data:FormData){
        return fetch('http://localhost:9000/api/v1/login',{body:data,method:"post"},).then(res=>res.json()).catch(err=>{
            console.log(err);
        })
    }
}