import { createRouter, createWebHistory } from "vue-router";
import App from "../App.vue";
import { inject } from "vue";
import LoginDefault from "../Auth/Login-Default.vue";
import MainPage from "../views/Berichtsheft/Main.vue";
import DashboardDefault from "../views/Dashboard/Dashboard-Default.vue";
import AdminDefault from "../views/Admin/Admin-Default.vue";
import InstructorDefault from "../views/Instructor/Instructor-Default.vue";
// import type { LoginServiceInterface } from "@/handler/loginHandler";
import CurriculumDefault from "@/views/Instructor/Curriculum-Default.vue";
import { loginService } from "../handler/loginHandler"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/Berichtsheft",
      name: "home",
      component: App,
      children: [
        {
          path: "/Berichtsheft",
          name: "Berichtsheft",
          component: MainPage,
          meta: { requiresAuth: true },
          beforeEnter: async (to, from, next) => {
            const role = await loginService.getUserRole();
            if (role === "user") next();
            else next("/Login");
          },
        },
        {
          path: "/Dashboard",
          name: "Dashboard",
          component: DashboardDefault,
          meta: { requiresAuth: true },
          beforeEnter: async (to, from, next) => {
            const role = await loginService.getUserRole();
            if (role === "user") next();
            else next("/Login");
          },
        },
        {
          path: "/Instructor",
          name: "Instructor",
          component: InstructorDefault,
          meta: { requiresAuth: true },
          beforeEnter: async (to, from, next) => {
            const role = await loginService.getUserRole();
            if (role === "instructor") next();
            else next("/Login");
          },
        },
        {
          path: "/Curriculum",
          name: "Curriculum",
          component: CurriculumDefault,
          meta: { requiresAuth: true },
          beforeEnter: async (to, from, next) => {
            const role = await loginService.getUserRole();
            if (role === "instructor") next();
            else next("/Login");
          },
        },
        {
          path: "/Admin",
          name: "Admin",
          component: AdminDefault,
          meta: { requiresAuth: true },
          beforeEnter: async (to, from, next) => {
            const role = await loginService.getUserRole();
            if (role === "admin") next();
            else next("/Login");
          },
        },
      ],
    },
    {
      path: "/Login",
      name: "Login",
      component: LoginDefault,
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    try {
      const statusCheck = await loginService.checkStatus();
      console.log(statusCheck);
      if (statusCheck) {
        next();
      } else next({ path: "/Login" });
    } catch (e) {
      next({ path: "/Login" });
    }
  } else {
    next();
  }
});
export default router;
