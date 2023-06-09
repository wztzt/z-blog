<template>
    <div >
        <button @click="routeA">测试1</button>
        <button @click="routeB">测试2</button>
        <router-view></router-view>
    </div>
</template>
  
  
<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter()
const routeA = ()=>{
    router.push("/index/a")
}

const routeB = () =>{
    router.push("/index/b")
}

</script>
  
<style scoped>
</style>
  
  