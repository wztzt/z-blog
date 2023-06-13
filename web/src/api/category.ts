
import { resultObj } from '@/modle/ResultObj'
import request from '@/utils/request'


const getCateGoryList = async ()=>{
    const response = await request.get<resultObj<Array<string>>>('/category')
    console.log(response)
    return response.data
}

const getCateGory =async (catename:string) => {
    const response = await request.get('/category/'+catename)
    return response.data
}

export {getCateGoryList, getCateGory}