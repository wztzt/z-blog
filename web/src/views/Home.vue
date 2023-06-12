<template>
    <div>
        <div>
            <button> 首页 </button>
            <button :key="index" v-for="(cate,index) in category"> {{ cate }} </button>


        </div>
    </div>
</template>

<script lang="ts" setup>
import request from '@/utils/request'
import { onMounted, reactive } from 'vue';

const category:Array<string> = reactive([])

onMounted(async () => {
    const response = await request.get<Array<string>>('/category')
    console.log(response.data)
    for(var v in response.data){
        category.push(response.data[v])
    }
})


</script>


<style scoped></style>