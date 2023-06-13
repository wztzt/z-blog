<template>
    <div>
        <div>
            <button> 首页 </button>
            <button :key="index" v-for="(cate,index) in category" @click="clickCatename(cate)"> {{ cate }} </button>


        </div>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive } from 'vue';
import { getCateGoryList,getCateGory } from '@/api/category';

const category:Array<string> = reactive([])

onMounted(async () => {
    const catlist = await getCateGoryList()
    for (var v in catlist.data){
        category.push(catlist.data[v])
    }
})

const clickCatename = async (catename:string) => {
    const catelist = await getCateGory(catename)
    console.log(catelist)
}

</script>


<style scoped></style>