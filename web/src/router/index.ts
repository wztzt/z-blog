import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'
import Index from '../views/index.vue'
import Header from '../components/Header.vue'

const reoutes:Array<RouteRecordRaw> = [
    {
        path:"/",
        component: Header,
        children:[
            {
                path:"/idx",
                component: Index
            }
        ]
    },
    {
        path: "/index",
        component: Index
    },
    {
        path: "/head",
        component: Header
    },
    {
        path: "/helloword",
        component : () => import("../components/HelloWorld.vue")
    }
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

