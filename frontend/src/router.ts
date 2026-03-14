import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './views/HomeView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/:pathMatch(.*)*',
            name: '404',
            component: () => import('./views/misc/404.vue') // lazy load
        }
    ]
})

export default router