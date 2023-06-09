import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'
import Index from '../views/index.vue'
import Header from '../components/Header.vue'
import HelloWorld from '@/components/HelloWorld.vue'

const reoutes:Array<RouteRecordRaw> = [
    {
        path:"/",
        redirect: "/index"
    },
    {
        path: "/index",
        component: Index,
        children:[
            {
                path: 'a',
                component: HelloWorld
            },
            {
                path:'b',
                component: Header
            }
        ]
    },
    {
        path: "/head",
        component: Header
    },
    {
        path: "/helloword",
        component : () => HelloWorld
    }
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

