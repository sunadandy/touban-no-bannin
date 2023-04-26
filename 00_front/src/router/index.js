import { createRouter, createWebHashHistory } from 'vue-router'
import TopView from '@/views/TopView.vue'
import HomeView from '@/views/HomeView.vue'
import CreateView from '@/views/CreateView.vue'
import EditView from '@/views/EditView.vue'
import ShowView from '@/views/ShowView.vue'

const routes = [
  {
    path: '/',
    name: 'top',
    component: TopView
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/create',
    name: 'create',
    component: CreateView
  },
  {
    path: '/:id/edit',
    name: 'edit',
    component: EditView,
  },
  {
    path: '/:id/show',
    name: 'show',
    component: ShowView,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
