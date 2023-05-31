
import Index from '../views/index.vue'

const routes = [
    {
        path: '/index',
        name: 'index',
        tile: '',
        component: Index,
    },
    {
        path: '/head',
        name: 'head',
        title: '',
        component: () => import('../components/Header.vue'),
    }
]

export default routes