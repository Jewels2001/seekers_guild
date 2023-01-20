import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/HelloWorld.vue'),
            props: {'msg': "Yay!"},
        },
        {
            path: '/about',
            name: 'about',
            component: () => import('../views/AboutView.vue'),
        },
        {
            path: '/leaderboard',
            name: 'leaderboard',
            component: () => import('../views/LeaderboardView.vue'),
        },
        {
            path: '/rules',
            name: 'rules',
            component: () => import('../views/Rules.vue'),
        },
        {
            path: '/party',
            name: 'party',
            component: () => import('../views/PartyView.vue'),
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'NotFound',
            component: () => import('../views/NotFound.vue')
        }
    ]
})

export default router