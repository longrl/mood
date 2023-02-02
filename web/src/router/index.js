import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/Home.vue')
    },
    {
      path: '/blog/:id',
      name: 'blog',
      component: () => import('../views/Blog.vue')
    },
    {
      path: '/blog/update/:id',
      name: 'editor',
      component: () => import('../views/Editor.vue')
    },
    {
      path: '/blog/auth',
      name: 'auth',
      component: () => import('../views/Auth.vue')
    }
  ]
})

export default router
