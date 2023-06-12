<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ msg: string }>()
const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>
  <div class="parent">
    <button class = 'child'>1</button>
    <button class = 'child'>2</button>
    <button class = 'child'>3</button>
  </div>
  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p class="center">
    Install
    <a href="https://github.com/vuejs/language-tools" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}

.parent{
  display: flex;
  flex-direction: column;
  color: red;
  border: 1px solid black;
  line-height: 100px;
  margin: 0 auto;
}

.child{
  width: 100px;
  height: 50px;
  vertical-align: middle;
  border: 2px solid blueviolet;
  color: #888;
}
</style>
