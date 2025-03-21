import { ElMessage } from 'element-plus'
import HTTP from './http'

const http = new HTTP('http://127.0.0.1:3000')

if (window.fetchURL) {
    fetchURL().then(baseURL => {
        http._baseURL = baseURL
    })
}

export default (route, data, callback) => {
    return http.postJson(route, data).then(res => {
        return res.json()
    }).then(result => {
        if (result.code) {
            console.warn(result)
        } else if (callback) {
            callback(result)
        }
    }).catch(() => {
        ElMessage({ message: "网络错误", type: "error" })
    })
}