import { createRouter, createWebHistory } from 'vue-router';
import TodoList from '@/components/TodoList.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: TodoList,
  },
  {
    path: '/todos',
    name: 'TodoList',
    component: TodoList,
  },
  {
    path: '/api',
    redirect: '/',
  },
  {
    path: '/:pathMatch(.*)*', // for undefined routes
    redirect: '/',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
