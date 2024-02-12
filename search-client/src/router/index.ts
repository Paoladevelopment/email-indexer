import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import EmailsView from '@/views/EmailsView.vue'
import EmailView from '@/views/EmailView.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/emails',
      name: 'emails',
      component: EmailsView
    },
    {
      path: '/emails/:messageId',
      name: 'email',
      component: EmailView
    }
  ]
})

export default router
