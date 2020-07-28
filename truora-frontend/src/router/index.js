import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import About from '../components/About.vue'
import Details from '../components/Details.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  { 
    path: '/info/:domainName', 
    name: 'details', 
    component: Details 
  }

]

const router = new VueRouter({
  routes
})

export default router
