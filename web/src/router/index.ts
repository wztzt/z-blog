import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'
import Index from '../views/index.vue'
import Header from '../components/Header.vue'
import HelloWorld from '@/components/HelloWorld.vue'
import Home from '@/views/Home.vue'

const reoutes:Array<RouteRecordRaw> = [
    {
        path:"/",
        redirect: "/index"
    },
    {
        path:'/home',
        component: Home
    },
    {
        path: "/index",
        component: Index,
        children:[
            {
                path: 'a',
                component: HelloWorld,
                props: route => ({ msg : route.query.msg})
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
        component : HelloWorld
    }
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

