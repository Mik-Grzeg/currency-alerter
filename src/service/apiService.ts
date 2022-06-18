import axios from "axios";

const apiClinet:any = axios.create({
    baseURL: '',
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
    timeout: 10000

})

export default {
    postAlert(alert:object){
        return apiClinet.post('',alert)
    }
}