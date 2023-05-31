import axios from 'axios'

const service = axios.create(
    {
        baseURL: '/api/v1',
    }
)


export default service
