<template>
    <div class="footer">
        <img src = "https://pic1.zhimg.com/80/v2-ac6cd5297968d638be2ddd9bfc810e98_720w.webp"/>
        <button @click="onClick">测试按钮</button>
    </div>
</template>


<script setup lang="ts">
import request from '../utils/request'
import {useRouter} from 'vue-router'
const router = useRouter()

const onClick = async ()=> {
    const response = await request.post('/login',{'a':'b'})
    console.log(response.data)
    if (response.data === 'login'){
        router.push('/index')
    }
    
}

</script>