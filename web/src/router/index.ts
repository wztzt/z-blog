import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'
import Index from '../views/index.vue'
import Header from '../components/Header.vue'

const reoutes:Array<RouteRecordRaw> = [
    {
        path: "/index",
        component: Index
    },
    {
        path: "/head",
        component: Header
    }
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

