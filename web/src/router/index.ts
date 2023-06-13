import {createRouter,createWebHistory,Router,RouteRecordRaw} from 'vue-router'

const Home = ()=>import('@/views/Home.vue')
const Index =  ()=>import('@/views/index.vue')
const HelloWord = ()=>import('@/components/HelloWorld.vue')
const Header = ()=>import('@/components/Header.vue')

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
                component: HelloWord,
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
        component : HelloWord
    }
]

const router:Router = createRouter(
    {
        history: createWebHistory(),
        routes: reoutes
    }

)

export default router

