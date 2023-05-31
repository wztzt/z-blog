import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'
import Index from '../views/index.vue'

const reoutes:Array<RouteRecordRaw> = [
    {
        path: "/index",
        component: Index
    },
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

