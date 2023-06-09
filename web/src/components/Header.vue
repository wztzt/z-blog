<template>
    <div class="center">
        <form action="api/v1/systeminfo" method="get">
            <input v-model="user"> 
            <br>
            <input v-model="password">
            <br>
            <button @click="register">注册</button>
            <button @click="login">登录</button>
            <input type="button" @click="systeminfo" value="系统信息">
        </form>
        

    </div>
</template>


<script setup lang="ts">
import request from '@/utils/request'
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

<style scoped>
div .center{
    background-color: turquoise;
    vertical-align: middle;
    margin: 0;
}
</style>