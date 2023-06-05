<template>
    <div>
        <form>
            <input v-model="user"> 
            <br>
            <input v-model="password">
        </form>
        <button @click="register">注册</button>
        <button @click="login">登录</button>
        <button @click="systeminfo">系统信息</button>
    </div>
</template>


<script setup lang="ts">
import request from '../utils/request'
import {useRouter} from 'vue-router'
import {ref } from 'vue';
const router = useRouter()

const user = ref<string>()
const password = ref<string>()





const register = async ()=> {
    const response = await request.post('/register',{'user':user.value, 'password': password.value})
    console.log(response.data)
    if (response.data.code === 0){
        router.push('/index')
    }
}

const login = async ()=> {
    const response = await request.post('/login', {'user': user.value, 'password': password.value})
    console.log(response.data)
    if (response.data.code === 0){
        router.push('/index')
    }
}

const systeminfo =async () => {
    const response = await request.get("/systeminfo", {})
    console.log(response.data)
}

</script>