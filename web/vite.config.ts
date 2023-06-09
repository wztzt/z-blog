import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  resolve:{
    alias:{
      '@': path.resolve(__dirname,'src'),
    }
  },
  plugins: [vue()],
  server: {
    host: '127.0.0.1',
    port: 3001,
    proxy:{
      '/api':{
        target:'http://127.0.0.1:8081',
        changeOrigin: true,
        rewrite: (path) => path.replace('/^\/api/', '')
      }
    }
  }
})
